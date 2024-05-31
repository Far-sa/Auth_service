package main

import (
	grpcServer "authentication-service/delivery/grpc"
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

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := "postgres://postgres:password@localhost:5432/auth_db?sslmode=disable"
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

	userRepo := repository.NewRepository(db)

	amqpUrl := "amqp://guest:guest@rabbitmq:5672/"
	publisher, _ := messaging.NewRabbitMQ(amqpUrl)
	authService := services.NewAuthService(userRepo, publisher)

	authHandler := grpcServer.NewGRPCHandler(authService)

	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authHandler) // Register handler for AuthService

	// Enable server reflection
	reflection.Register(grpcServer)
	// Serve
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

}
