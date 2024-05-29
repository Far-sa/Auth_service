package entity

import "time"

// UserProfile represents a user profile in the user service
type UserProfile struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FullName  *string    `json:"full_name,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}
