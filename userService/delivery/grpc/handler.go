package grpcHandler

import (
	"context"
	"user-service/internal/interfaces"
	user "user-service/pb"
	mapper "user-service/utils/protobufMapper"
	// Replace with your package paths
)

type grpcHandler struct {
	userService interfaces.UserService
	user.UnimplementedUserServiceServer
}

func New(userService interfaces.UserService) *grpcHandler {
	return &grpcHandler{userService: userService}
}

func (h *grpcHandler) CreateUser(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	//* map proto to param
	pb := mapper.PbToParamRegisterRequest(req)

	user, err := h.userService.Register(ctx, pb)
	if err != nil {
		return nil, err
	}

	//* map param to proto
	resp := mapper.ToPbRegisterResponse(user)
	return resp, nil
}

func (h *grpcHandler) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {

	//* map proto to param
	pb := mapper.PbToParamGetUserRequest(req)

	user, err := h.userService.GetUser(ctx, pb.UserID)
	if err != nil {
		return nil, err
	}

	//* map param to proto
	resp := mapper.ToPbUserProfileResponse(user)
	return resp, nil
}

func (h *grpcHandler) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	// Implement user update logic
	// Publish event to RabbitMQ (example)

	//* publishEvent("user_updates", []byte("User updated"))
	return &user.UpdateUserResponse{Success: true}, nil
}

func (h *grpcHandler) GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error) {

	paramReq := mapper.PbToParamGetUserByEmail(req)

	user, err := h.userService.GetUserByEmail(ctx, paramReq.Email)
	if err != nil {
		return nil, err
	}

	//TODO update param for user response
	resp := mapper.ToPbGetUserByEmail(user)

	return resp, nil
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
