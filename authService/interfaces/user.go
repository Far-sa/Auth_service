package interfaces

import (
	"authentication-service/domain/entities"
	"authentication-service/domain/param"
	"context"
)

type AuthenticationService interface {
	Login(ctx context.Context, loginRequest param.LoginRequest) (*entities.User, error)
	Register(ctx context.Context, req param.RegisterRequest) error
	// ... other authentication methods (e.g., refresh token)
}

type UserRepository interface {
	Save(ctx context.Context, user *entities.User) error
	FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entities.User, error)
	// FindByID(ctx context.Context, userID string) (*entities.User, error) // Optional
}

type MessagePublisher interface {
	PublishUserAuthenticated(userID string) error
}