package main

import (
	grpcHandler "authentication-service/delivery/grpc"
	httpHandler "authentication-service/delivery/http"
	"authentication-service/domain/services"
	"authentication-service/infrastructure/database"
	"authentication-service/infrastructure/database/migrator"
	"authentication-service/infrastructure/messaging"
	"authentication-service/infrastructure/repository"
	auth "authentication-service/pb"
	"log"
	"net"
	"path"
	standard_runtime "runtime"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {

	dsn := "postgres://postgres:password@postgres-auth:5432/auth_db?sslmode=disable"
	db, err := database.NewSQLDB(dsn)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
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

	amqpUrl := "amqp://guest:guest@rabbitmq:5672/"
	publisher, _ := messaging.NewRabbitMQ(amqpUrl)
	authService := services.NewAuthService(userRepo, publisher)

	// Start gRPC server in a separate goroutine
	go func() {
		authHandler := grpcHandler.NewGRPCHandler(authService)

		grpcServer := grpc.NewServer()
		auth.RegisterAuthServiceServer(grpcServer, authHandler)

		lis, err := net.Listen("tcp", ":")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Start gRPC-Gateway in a separate goroutine
	// go func() {
	// 	ctx := context.Background()
	// 	ctx, cancel := context.WithCancel(ctx)
	// 	defer cancel()

	// 	if err := gateway.RunHTTPGateway(ctx, ":50052"); err != nil {
	// 		log.Fatalf("Failed to run gRPC-Gateway: %v", err)
	// 	}
	// }()

	// Start HTTP server in the main goroutine
	authHandler := httpHandler.NewHTTPAuthHandler(authService)

	e := echo.New()
	e.POST("/auth/login", authHandler.Login)

	log.Println("HTTP server is running on port 8081")
	if err := e.Start(":8081"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
