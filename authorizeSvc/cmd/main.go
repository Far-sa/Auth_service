package main

import (
	"authorization-service/cmd/gateway"
	grpcHandler "authorization-service/delivery/gprc"
	httpHandler "authorization-service/delivery/http"
	"authorization-service/infrastructure/database"
	"authorization-service/infrastructure/database/migrator"
	"authorization-service/infrastructure/messaging/rabbitmq"
	"authorization-service/infrastructure/repository"
	"authorization-service/internal/service"
	authz "authorization-service/pb"
	"context"
	"net"
	"path"
	standard_runtime "runtime"

	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {

	//TODO add config to load them from environment variable

	dsn := "postgres://postgres:password@localhost:5432/authz_db?sslmode=disable"
	db, err := database.NewSQLDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	_, filename, _, _ := standard_runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../infrastructure/database/migrations")
	// Create a new migrator instance.
	migrator, err := migrator.NewMigrator(db.Conn(), dir)
	if err != nil {
		log.Fatalf("Failed to create migrator: %v", err)
	}

	// Apply all up migrations.
	if err := migrator.Up(); err != nil {
		log.Fatalf("Failed to migrate up: %v", err)
	}

	// Initialize repository, service, and handler
	userRepo := repository.NewRepository(db)

	amqpUrl := "amqp://guest:guest@localhost:5672/"
	rabbitAdapter, err := rabbitmq.NewRabbitMQAdapter(amqpUrl)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ adapter: %v", err)
	}
	defer rabbitAdapter.Close()

	consumer, err := rabbitmq.NewRabbitMQAdapter(amqpUrl)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	}

	authzService := service.NewAuthzService(userRepo, consumer)
	authzService.ListenForUserEvents()

	// Start gRPC server in a separate goroutine
	go func() {
		authzHandler := grpcHandler.NewGRPCHandler(authzService)

		grpcServer := grpc.NewServer()
		authz.RegisterAuthorizationServiceServer(grpcServer, authzHandler)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Start gRPC-Gateway in a separate goroutine
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		if err := gateway.RunHTTPGateway(ctx, ":50051"); err != nil {
			log.Fatalf("Failed to run gRPC-Gateway: %v", err)
		}
	}()

	// Start HTTP server in the main goroutine
	authzHandler := httpHandler.NewHTTPAuthzHandler(authzService)

	e := echo.New()
	e.POST("/authz/assign-role", authzHandler.AssignRole)
	e.POST("/authz/update-role", authzHandler.UpdateRole)
	// e.GET("/getUser", authzHandler.CheckPermission)

	log.Println("HTTP server is running on port 8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Handle OS signals for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

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
