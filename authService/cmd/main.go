package main

// import (
// 	delivery "authentication-service/delivery/grpc"
// 	"authentication-service/domain/services"
// 	"authentication-service/infrastructure/database"
// 	"authentication-service/infrastructure/database/migrator"
// 	"authentication-service/infrastructure/messaging"
// 	"authentication-service/infrastructure/repository"
// 	"context"
// 	"log"
// 	"net"
// 	"path"
// 	standard_runtime "runtime"
// 	"user-service/delivery/gateway"
// )

// func main() {
// 	lis, err := net.Listen("tcp", ":50052")

// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	dsn := "postgres://postgres:password@localhost:5432/auth_db?sslmode=disable"
// 	db, err := database.NewSQLDB(dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to create database: %v", err)
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
// 	publisher, _ := messaging.NewRabbitMQ(amqpUrl)
// 	authService := services.NewAuthService(userRepo, publisher)

// 	authHandler := delivery.NewGRPC(authService)
// 	authHandler.Serve()

// 	ctx := context.Background()
// 	if err := gateway.RunHTTPGateway(ctx, lis.Addr().String()); err != nil {
// 		log.Fatalf("Failed to run gRPC-Gateway: %v", err)
// 	}
// }
