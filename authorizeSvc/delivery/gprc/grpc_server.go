package delivery

import (
	"authorization-service/delivery/gprc/handler"
	"authorization-service/internal/interfaces"
	"authorization-service/pb"
	"net"

	// Replace with your package paths
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	authService interfaces.AuthorizationService // Interface for dependency injection
	authHandler *handler.AuthzHandler           // Handler for authentication service methods
}

func NewGRPCServer(authService interfaces.AuthorizationService, authHandler *handler.AuthzHandler) *grpcServer {
	return &grpcServer{authService: authService, authHandler: authHandler}
}

func (s *grpcServer) Serve(lis net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthorizationServiceServer(grpcServer, s.authHandler) // Register handler for AuthService

	// Enable server reflection
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}

// Implement handler functions for other gRPC service methods defined in your `auth.proto` file
