package handler

import (
	"context"
	"user-service/internal/interfaces"
	pb "user-service/pb"
)

type UserHandler struct {
	userService interfaces.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userService.GetUser(req.UserId)
	if err != nil {
		return nil, err
	}
	// createdAt, _ := time.Parse(time.RFC3339, user.CreateAt)
	return &pb.GetUserResponse{Email: user.UserProfile.Email, Name: *user.UserProfile.FullName}, nil
}

// func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
// 	user := entity.User{
// 		UserID:    req.UserId,
// 		Name:      req.Name,
// 		Email:     req.Email,
// 		CreatedAt: req.CreatedAt.AsTime(),
// 	}
// 	err := h.userService.CreateUser(user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.CreateUserResponse{UserId: user.UserID}, nil
// }
