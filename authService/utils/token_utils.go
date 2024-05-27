package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	//jwt "github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	// Ideally, this should be loaded from an environment variable or config file for security reasons.
	SecretKeyEnvVar = "JWT_SECRET_KEY"
)

type MyClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateAccessToken(userID string) (string, error) {
	// Retrieve the secret key from an environment variable
	secretKey := os.Getenv(SecretKeyEnvVar)
	if secretKey == "" {
		return "", errors.New("secret key not found")
	}

	// Set token claims, such as expiration time
	expirationTime := time.Now().Add(24 * time.Hour)
	myClaims := MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken(userID int) (string, error) {
	// Retrieve the secret key from an environment variable
	secretKey := os.Getenv(SecretKeyEnvVar)
	if secretKey == "" {
		return "", errors.New("secret key not found")
	}

	// Set token claims, with a longer expiration time for refresh tokens
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 7 days validity
	myClaims := MyClaims{
		UserID: strconv.Itoa(userID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	return token.SignedString([]byte(secretKey))
}

func ValidateAccessToken(tokenString string) (*MyClaims, error) {
	// Retrieve the secret key from an environment variable
	secretKey := os.Getenv(SecretKeyEnvVar)
	if secretKey == "" {
		return nil, errors.New("secret key not found")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token claims are valid and of the correct type
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
