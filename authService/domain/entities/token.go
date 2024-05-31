package entities

import "time"

type Token struct {
	ID                    int       `json:"id" gorm:"primaryKey"`
	UserID                int       `json:"user_id" gorm:"not null"`
	AccessToken           string    `json:"access_token" gorm:"not null"`
	RefreshToken          string    `json:"refresh_token" gorm:"not null"`
	CreatedAt             time.Time `json:"created_at" gorm:"autoCreateTime"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at" gorm:"not null"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at" gorm:"not null"`
}
