package param

import (
	"authentication-service/domain/entities"
	"time"
)

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

// LoginResponse represents the response after a successful login
type LoginResponse struct {
	TokenPair entities.TokenPair `json:"token_pair"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse represents the response after a successful registration
type RegisterResponse struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// !
type UserResponse struct {
	ID           int    `json:"id"`
	PasswordHash string `json:"password_hash"`
	Error        string `json:"error,omitempty"`
}

type RegisterUserRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type RegisterUserResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
