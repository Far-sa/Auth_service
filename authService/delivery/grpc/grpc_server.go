package delivery

import (
	"authentication-service/delivery/grpc/handler"
	"authentication-service/interfaces"
	"authentication-service/pb"
	"net"

	// Replace with your package paths
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	authService interfaces.AuthenticationService // Interface for dependency injection
	authHandler *handler.AuthHandler             // Handler for authentication service methods
}

func NewGRPCServer(authService interfaces.AuthenticationService, authHandler *handler.AuthHandler) *grpcServer {
	return &grpcServer{authService: authService, authHandler: authHandler}
}

func (s *grpcServer) Serve(lis net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, s.authHandler) // Register handler for AuthService

	// Enable server reflection
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}

// Implement handler functions for other gRPC service methods defined in your `auth.proto` file
