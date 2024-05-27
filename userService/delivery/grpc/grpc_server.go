package delivery

import (
	"net"
	"user-service/delivery/grpc/handler"
	"user-service/internal/interfaces"
	"user-service/pb"

	// Replace with your package paths
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	userService interfaces.UserService
	userHandler *handler.UserHandler // Handler for authentication service methods
}

func NewGRPCServer(userService interfaces.UserService, userHandler *handler.UserHandler) *grpcServer {
	return &grpcServer{userService: userService, userHandler: userHandler}
}

func (s *grpcServer) Serve(lis net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, s.userHandler) // Register handler for AuthService

	// Enable server reflection
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}
