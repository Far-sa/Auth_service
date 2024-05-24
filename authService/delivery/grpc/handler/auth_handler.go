package handler

import (
	"authentication-service/domain/param"
	"authentication-service/interfaces"
	"authentication-service/pb"
	"authentication-service/utils"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthHandler handles gRPC requests for authentication
type AuthHandler struct {
	authService interfaces.AuthenticationService
	// Embed the pb.AuthServiceServer interface to implement its methods
	pb.UnimplementedAuthServiceServer
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authService interfaces.AuthenticationService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login implements the gRPC Login method
func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.authService.Login(ctx, param.LoginRequest{
		UsernameOrEmail: req.GetUsernameOrEmail(),
		Password:        req.GetPassword(),
	})
	if err != nil {
		// ... (existing error handling)
	}

	accessToken, err := utils.GenerateAccessToken(user.ID) // Use the token generation utility
	if err != nil {
		return nil, status.Error(codes.Internal, "Error generating access token")
	}

	return &pb.LoginResponse{AccessToken: accessToken}, nil
}

// Register implements the gRPC Register method
func (h *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	panic("")
}

// generateAccessToken (placeholder) generates an access token based on user ID (replace with your logic)
func generateAccessToken(userID string) (string, error) {
	// ... (implementation for access token generation)
	return "", nil
}
