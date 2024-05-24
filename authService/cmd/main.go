package main

import (
	"authentication-service/internal/di"
	"authentication-service/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	defaultPort = "50051"
	jwtKey      = "supersecretkey"
)

func main() {
	// Initialize Auth Handler with DI
	authHandler, cleanup, err := di.InitializeAuthHandler([]byte(jwtKey))
	if err != nil {
		log.Fatalf("Failed to initialize auth handler: %v", err)
	}
	defer cleanup()

	listener, err := net.Listen("tcp", ":"+defaultPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", defaultPort, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authHandler)

	log.Println("Starting gRPC server on port", defaultPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
