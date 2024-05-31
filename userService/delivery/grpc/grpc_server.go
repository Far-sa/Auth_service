package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"user-service/internal/interfaces"
	user "user-service/pb"

	// Replace with your package paths
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	userService interfaces.UserService
	user.UnimplementedUserServiceServer
}

func New(userService interfaces.UserService) *grpcServer {
	return &grpcServer{userService: userService}
}

func (s *grpcServer) Start() {
	// listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50053))
	if err != nil {
		panic(err)
	}

	userServiceServer := &grpcServer{}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userServiceServer)

	// Enable server reflection
	reflection.Register(grpcServer)

	// Serve
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
}

func (s *grpcServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	user, err := s.userService.GetUser(req.UserId)
	if err != nil {
		return nil, err
	}
	// createdAt, _ := time.Parse(time.RFC3339, user.CreateAt)
	return &user.GetUserResponse{
		Email: user.UserProfile.Email,
		Name:  *user.UserProfile.FullName,
	}, nil
}

func (s *grpcServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	// Implement user update logic
	// Publish event to RabbitMQ (example)

	//* publishEvent("user_updates", []byte("User updated"))
	return &user.UpdateUserResponse{Success: true}, nil
}
