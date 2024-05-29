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

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

// AccessToken represents the access token generated upon successful login
type AccessToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// RefreshToken represents the refresh token generated upon successful login
type RefreshToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// TokenPair represents a pair of access and refresh tokens
type TokenPair struct {
	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
}

// type Token struct {
// 	ID        int       `json:"id" gorm:"primaryKey"`
// 	UserID    int       `json:"user_id" gorm:"not null"`
// 	Token     string    `json:"token" gorm:"not null"`
// 	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
// }

// // AuthToken represents the authentication token generated upon successful login
// type AuthToken struct {
// 	Token     string    `json:"token"`
// 	ExpiresAt time.Time `json:"expires_at"`
// }
