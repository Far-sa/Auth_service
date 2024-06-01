package entity

import "time"

// UserProfile represents a user profile in the user service
type UserProfile struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	FullName    string    `json:"full_name,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Birthdate   time.Time `json:"birthdate,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
