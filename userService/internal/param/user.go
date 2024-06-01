package param

import (
	"time"
	"user-service/internal/entity"
)

// * Paramas

// type UserInfo struct {
// 	ID          string `json:"id"`
// 	PhoneNumber string `json:"phone_number"`
// 	UserName    string `json:"username"`
// 	Email       string `json:"email"`
// 	CreateAt    string `json:"create_at"`
// }

// type ProfileRequest struct {
// 	UserID string `json:"user_id"`
// }
// type ProfileResponse struct {
// 	UserInfo
// }

type GetUser struct {
	UserID string `json:"user_id"`
}

type GetUserByEmail struct {
	Email string `json:"email"`
}

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

type RegisterRequest struct {
	FullName    string `json:"full_name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}

type UserInfo struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}
type ProfileResponse struct {
	Name string `json:"name"`
}
