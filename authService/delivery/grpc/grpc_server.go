package delivery

import (
	"authentication-service/domain/param"
	"authentication-service/interfaces"
	"authentication-service/pb"
	"context"
	"fmt"
	"log"
	"net"

	// Replace with your package paths
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	authService interfaces.AuthenticationService // Interface for dependency injection
	pb.UnimplementedAuthServiceServer
}

func NewGRPC(authService interfaces.AuthenticationService) *grpcServer {
	return &grpcServer{authService: authService}
}

func (s *grpcServer) Serve() {

	// listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50052))
	if err != nil {
		panic(err)
	}

	authServiceServer := grpcServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authServiceServer) // Register handler for AuthService

	// Enable server reflection
	reflection.Register(grpcServer)
	// Serve
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
}

// Implement handler functions for other gRPC service methods defined in your `auth.proto` file
func (s grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := s.authService.Login(ctx, param.LoginRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{AccessToken: resp.TokenPair.AccessToken, UserId: resp.UserID}, nil
}

// Register implements the gRPC Register method
func (s grpcServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	panic("")
}
