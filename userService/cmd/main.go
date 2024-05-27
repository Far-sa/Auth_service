package main

import (
	"authentication-service/infrastructure/database/migrator"
	"context"
	"log"
	"net"
	"net/http"
	"path"
	standard_runtime "runtime"

	"user-service/delivery/grpc/handler"
	"user-service/infrastructure/database"
	"user-service/infrastructure/messaging"
	"user-service/infrastructure/repository"
	"user-service/internal/interfaces"
	"user-service/internal/service"
	"user-service/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGRPCServer(lis net.Listener, userService interfaces.UserService) {
	grpcServer := grpc.NewServer()
	userHandler := handler.NewUserHandler(userService)
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	reflection.Register(grpcServer)

	log.Printf("Serving gRPC on %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func runHTTPGateway(ctx context.Context, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		return err
	}

	log.Println("Serving gRPC-Gateway on http://localhost:8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	lis, err := net.Listen("tcp", ":50052")

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

	go runGRPCServer(lis, userSvc)
	ctx := context.Background()
	if err := runHTTPGateway(ctx, lis.Addr().String()); err != nil {
		log.Fatalf("Failed to run gRPC-Gateway: %v", err)
	}
}
