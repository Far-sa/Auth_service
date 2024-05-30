package param

import (
	"time"
)

type LoginRequest struct {
	Email    string `json:"username_or_email"`
	Password string `json:"password"`
}

// LoginResponse represents the response after a successful login
type LoginResponse struct {
	UserID    string    `json:"user_id"`
	TokenPair TokenPair `json:"token_pair"`
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

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"` // Unix timestamp
}
