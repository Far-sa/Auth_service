package main

import (
	"authorization-service/internal/di"
	"authorization-service/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	defaultPort = "50052"
)

func main() {
	// Initialize Authz Handler with DI
	authzHandler, cleanup, err := di.InitializeAuthzHandler()
	if err != nil {
		log.Fatalf("Failed to initialize authz handler: %v", err)
	}
	defer cleanup()

	listener, err := net.Listen("tcp", ":"+defaultPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", defaultPort, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthorizationServiceServer(grpcServer, authzHandler)

	log.Println("Starting gRPC server on port", defaultPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
