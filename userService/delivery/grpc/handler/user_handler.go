package handler

import (
	"context"
	"user-service/internal/entity"
	"user-service/internal/interfaces"
	pb "user-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	userService interfaces.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := entity.User{
		UserID:    req.UserId,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: req.CreatedAt.AsTime(),
	}
	err := h.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{UserId: user.UserID}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userService.GetUser(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		UserId:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}, nil
}
