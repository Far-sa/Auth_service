package main

import (
	"authorization-service/delivery/gprc/handler"
	"authorization-service/infrastructure/database"
	"authorization-service/infrastructure/database/migrator"
	"authorization-service/infrastructure/messaging"
	"authorization-service/infrastructure/messaging/rabbitmq"
	"authorization-service/infrastructure/repository"
	"authorization-service/internal/interfaces"
	"authorization-service/internal/service"
	"authorization-service/pb"

	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGRPCServer(lis net.Listener, authzService interfaces.AuthorizationService) error {
	grpcServer := grpc.NewServer()
	authzHandler := handler.NewAuthzHandler(authzService)
	pb.RegisterAuthorizationServiceServer(grpcServer, authzHandler)
	reflection.Register(grpcServer)

	log.Printf("Serving gRPC on %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
	return nil
}

func runHTTPGateway(ctx context.Context, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterAuthorizationServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	log.Println("Serving gRPC-Gateway on http://localhost:8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	dsn := "postgres://root:password@localhost:5432/authz_db?sslmode=disable"
	db, err := database.NewSQLDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create a new migrator instance.
	migrator, err := migrator.NewMigrator(db.Conn(), "../infrastructure/database/migrations") // Pass db instead of db.DB
	if err != nil {
		log.Fatalf("Failed to create migrator: %v", err)
	}

	// Apply all up migrations.
	if err := migrator.Up(); err != nil {
		log.Fatalf("Failed to migrate up: %v", err)
	}

	// Initialize repository, service, and handler
	userRepo := repository.NewRepository(db)

	amqpUrl := "amqp://guest:guest@rabbitmq:5672/"
	rabbitAdapter, err := rabbitmq.NewRabbitMQAdapter(amqpUrl)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ adapter: %v", err)
	}
	defer rabbitAdapter.Close()

	consumer, err := messaging.NewRabbitMQConsumer(rabbitAdapter, "user_authenticated_queue", "user.authenticated", "auth_exchange", nil)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	}

	authzService := service.NewAuthorizationService(userRepo, consumer)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := authzService.Start(); err != nil {
			log.Printf("Authorization service stopped with error: %v", err)
			cancel()
		}
	}()

	// Start the authorization service
	go func() {
		if err := runGRPCServer(lis, authzService); err != nil {
			log.Printf("gRPC server stopped with error: %v", err)
			cancel()
		}
	}()

	//go runGRPCServer(lis, authzService)

	go func() {
		if err := runHTTPGateway(ctx, lis.Addr().String()); err != nil {
			log.Printf("Failed to run gRPC-Gateway: %v", err)
			cancel()
		}
	}()

	// Handle OS signals for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("Shutdown initiated")
	case sig := <-signalChan:
		log.Printf("Received signal: %v. Shutting down...", sig)
		cancel()
	}

	// Perform any cleanup tasks here if necessary
	log.Println("Server gracefully stopped")
}
