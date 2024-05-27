package main

import (
	"authentication-service/delivery/grpc/handler"
	"authentication-service/domain/services"
	"authentication-service/infrastructure/database"
	"authentication-service/infrastructure/database/migrator"
	"authentication-service/infrastructure/messaging"
	"authentication-service/infrastructure/repository"
	"authentication-service/interfaces"
	"authentication-service/pb"
	"context"
	"log"
	"net"
	"net/http"
	"path"
	standard_runtime "runtime"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGRPCServer(lis net.Listener, authService interfaces.AuthenticationService) {
	grpcServer := grpc.NewServer()
	authHandler := handler.NewAuthHandler(authService)
	pb.RegisterAuthServiceServer(grpcServer, authHandler)
	reflection.Register(grpcServer)

	log.Printf("Serving gRPC on %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func runHTTPGateway(ctx context.Context, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
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
	// Initialize repository, service, and handler
	userRepo := repository.NewRepository(db)

	amqpUrl := "amqp://guest:guest@rabbitmq:5672/"
	publisher, _ := messaging.NewRabbitMQPublisher(amqpUrl)
	authService := services.NewAuthService(userRepo, publisher)

	// grpc := delivery.NewGRPCServer(authService)

	// // Use the Serve function from the gRPC server implementation
	// go func() {
	// 	if err := grpc.Serve(lis); err != nil {
	// 		log.Fatalf("Failed to serve gRPC server: %v", err)
	// 	}
	// }()

	go runGRPCServer(lis, authService)
	ctx := context.Background()
	if err := runHTTPGateway(ctx, lis.Addr().String()); err != nil {
		log.Fatalf("Failed to run gRPC-Gateway: %v", err)
	}
}
