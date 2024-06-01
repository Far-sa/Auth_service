package main

// import (
// 	"authorization-service/cmd/gateway"
// 	grpcHandler "authorization-service/delivery/gprc"
// 	"authorization-service/infrastructure/database"
// 	"authorization-service/infrastructure/database/migrator"
// 	"authorization-service/infrastructure/messaging"
// 	"authorization-service/infrastructure/messaging/rabbitmq"
// 	"authorization-service/infrastructure/repository"
// 	"authorization-service/internal/service"
// 	authz "authorization-service/pb"
// 	"path"
// 	standard_runtime "runtime"

// 	"context"
// 	"log"
// 	"net"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"
// )

// func main() {

// 	//TODO add config to load them from environment variable
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	dsn := "postgres://postgres:password@localhost:5432/authz_db?sslmode=disable"
// 	db, err := database.NewSQLDB(dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	_, filename, _, _ := standard_runtime.Caller(0)
// 	dir := path.Join(path.Dir(filename), "../infrastructure/database/migrations")
// 	// Create a new migrator instance.
// 	migrator, err := migrator.NewMigrator(db.Conn(), dir)
// 	if err != nil {
// 		log.Fatalf("Failed to create migrator: %v", err)
// 	}

// 	// Apply all up migrations.
// 	if err := migrator.Up(); err != nil {
// 		log.Fatalf("Failed to migrate up: %v", err)
// 	}

// 	// Initialize repository, service, and handler
// 	userRepo := repository.NewRepository(db)

// 	amqpUrl := "amqp://guest:guest@rabbitmq:5672/"
// 	rabbitAdapter, err := rabbitmq.NewRabbitMQAdapter(amqpUrl)
// 	if err != nil {
// 		log.Fatalf("Failed to create RabbitMQ adapter: %v", err)
// 	}
// 	defer rabbitAdapter.Close()

// 	consumer, err := messaging.NewRabbitMQConsumer(rabbitAdapter, "user_authenticated_queue", "user.authenticated", "auth_exchange", nil)
// 	if err != nil {
// 		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
// 	}

// 	authzService := service.NewAuthorizationService(userRepo, consumer)

// 	authzHandler := grpcHandler.NewGRPCHandler(authzService)

// 	grpcServer := grpc.NewServer()
// 	authz.RegisterAuthorizationServiceServer(grpcServer, authzHandler)

// 	// Enable server reflection
// 	reflection.Register(grpcServer)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	go func() {
// 		if err := grpcServer.Serve(lis); err != nil {
// 			log.Fatalf("Failed to serve gRPC server: %v", err)
// 		}
// 	}()

// 	if err := gateway.RunHTTPGateway(ctx, lis.Addr().String()); err != nil {
// 		log.Fatalf("Failed to run gRPC-Gateway: %v", err)
// 	}

// 	// Handle OS signals for graceful shutdown
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

// 	select {
// 	case <-ctx.Done():
// 		log.Println("Shutdown initiated")
// 	case sig := <-signalChan:
// 		log.Printf("Received signal: %v. Shutting down...", sig)
// 		cancel()
// 	}

// 	// Perform any cleanup tasks here if necessary
// 	log.Println("Server gracefully stopped")
// }
