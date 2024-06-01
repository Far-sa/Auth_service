package main

import (
	"authentication-service/infrastructure/database/migrator"
	"context"
	"log"
	"net"
	"path"
	standard_runtime "runtime"

	"user-service/cmd/gateway"
	grpcHandler "user-service/delivery/grpc"
	"user-service/infrastructure/database"
	"user-service/infrastructure/messaging"
	"user-service/infrastructure/repository"
	"user-service/internal/service"
	user "user-service/pb"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	dsn := "postgres://postgres:password@localhost:5432/user_db?sslmode=disable"
	db, err := database.NewSQLDB(dsn)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	//! runtime.Caller(0)` returns the file name and line number of the caller's caller.
	//! `path.Dir(filename)` returns the directory of the `main.go` file. `path.Join` constructs the path to the migrations directory.
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
	userEvent, _ := messaging.NewRabbitMQ(amqpUrl)
	userSvc := service.NewUserService(userRepo, userEvent)

	userHandler := grpcHandler.New(userSvc)

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userHandler)

	// Enable server reflection
	reflection.Register(grpcServer)

	// Serve
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	ctx := context.Background()
	if err := gateway.RunHTTPGateway(ctx, lis.Addr().String()); err != nil {
		log.Fatalf("Failed to run gRPC-Gateway: %v", err)
	}

}
