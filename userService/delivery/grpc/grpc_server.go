package delivery

import (
	"log"
	"net"
	"user-service/delivery/grpc/handler"
	"user-service/internal/interfaces"
	"user-service/pb"

	// Replace with your package paths
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	userService interfaces.UserService
	userHandler *handler.UserHandler // Handler for authentication service methods
}

func NewGRPCServer(userService interfaces.UserService) *GRPCServer {
	userHandler := handler.NewUserHandler(userService)
	return &GRPCServer{
		userService: userService,
		userHandler: userHandler,
	}
}

func (s *GRPCServer) Serve(lis net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, s.userHandler)

	// Enable server reflection
	reflection.Register(grpcServer)

	log.Printf("Serving gRPC on %s", lis.Addr().String())
	return grpcServer.Serve(lis)
}
