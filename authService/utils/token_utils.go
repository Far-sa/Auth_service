package utils

import (
	"github.com/dgrijalva/jwt-go" // (or your preferred token library)
)

type MyClaims struct {
	UserID string `json:"user_id"`
	// ... (other relevant claims)
}

func GenerateAccessToken(userID string) (string, error) {
	myClaims := MyClaims{UserID: userID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	return token.SignedString([]byte("your_secret_key")) // Replace with a secure secret key
}

func ValidateAccessToken(tokenString string) (*MyClaims, error) {
	// Implement validation logic using your token library
	// ...
}
