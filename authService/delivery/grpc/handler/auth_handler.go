package handler

import (
	"authentication-service/domain/param"
	"authentication-service/interfaces"
	"authentication-service/pb"
	"context"
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
	resp, err := h.authService.Login(ctx, param.LoginRequest{
		UsernameOrEmail: req.GetUsernameOrEmail(),
		Password:        req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{AccessToken: resp.TokenPair.AccessToken, UserId: resp.UserID}, nil
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
