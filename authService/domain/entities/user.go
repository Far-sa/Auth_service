package entities

import "time"

//	type User struct {
//		ID           int       `json:"id" gorm:"primaryKey"`
//		Username     string    `json:"username" gorm:"unique;not null"`
//		Email        string    `json:"email" gorm:"unique;not null"`
//		PasswordHash string    `json:"password_hash" gorm:"not null"`
//		CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
//		UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
//		Tokens       []Token   `json:"tokens" gorm:"foreignKey:UserID"`
//	}
type Token struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
}
