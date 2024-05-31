package grpcHandler

import (
	"context"
	"user-service/internal/interfaces"
	user "user-service/pb"
	// Replace with your package paths
)

type grpcHandler struct {
	userService interfaces.UserService
	user.UnimplementedUserServiceServer
}

func New(userService interfaces.UserService) *grpcHandler {
	return &grpcHandler{userService: userService}
}

func (s *grpcHandler) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {

	//* map proto to param
	
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

func (s *grpcHandler) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	// Implement user update logic
	// Publish event to RabbitMQ (example)

	//* publishEvent("user_updates", []byte("User updated"))
	return &user.UpdateUserResponse{Success: true}, nil
}

func (s *grpcHandler) GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error) {

	panic("not implemented")
}

// func (s *grpcServer) Start() {
// 	// listener
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50053))
// 	if err != nil {
// 		panic(err)
// 	}

// 	userServiceServer := &grpcServer{}

// 	grpcServer := grpc.NewServer()
// 	user.RegisterUserServiceServer(grpcServer, userServiceServer)

// 	// Enable server reflection
// 	reflection.Register(grpcServer)

// 	// Serve
// 	go func() {
// 		if err := grpcServer.Serve(lis); err != nil {
// 			log.Fatalf("Failed to serve gRPC server: %v", err)
// 		}
// 	}()
// }
