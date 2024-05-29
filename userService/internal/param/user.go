package param

import (
	"time"
	"user-service/internal/entity"
)

// * Paramas

type UserInfo struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	CreateAt    string `json:"create_at"`
}

type ProfileRequest struct {
	UserID string `json:"user_id"`
}
type ProfileResponse struct {
	UserInfo
}

// UserRegisteredEvent represents the event when a user registers
type UserRegisteredEvent struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// UserProfileResponse represents the response after querying a user profile
type UserProfileResponse struct {
	UserProfile entity.UserProfile `json:"user_profile"`
}
