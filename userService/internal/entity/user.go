package entity

import "time"

type UserDetail struct {
	UserID         int       `json:"user_id" gorm:"primaryKey"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	PhoneNumber    string    `json:"phone_number"`
	Address        string    `json:"address"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
