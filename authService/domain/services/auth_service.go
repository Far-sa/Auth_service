package services

import (
	"authentication-service/domain/entities"
	"authentication-service/domain/param"
	"authentication-service/interfaces"
	"authentication-service/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AuthenticationService interface defines methods for user authentication

// AuthService implements the AuthenticationService interface
type AuthService struct {
	userRepository   interfaces.UserRepository
	messagePublisher interfaces.MessagePublisher
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepository interfaces.UserRepository, messagePublisher interfaces.MessagePublisher) *AuthService {
	return &AuthService{
		userRepository:   userRepository,
		messagePublisher: messagePublisher,
	}
}

// ! Event Publication: The authentication service publishes a UserRegisteredEvent to RabbitMQ.
//!     event := models.UserRegisteredEvent / user_registered

func (s *AuthService) Register(ctx context.Context, registerRequest param.RegisterRequest) error {
	// Hash the password
	passwordHash, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		return err
	}

	// Create a user request to send to the User Service
	userRequest := param.RegisterUserRequest{
		Username:     registerRequest.Username,
		Email:        registerRequest.Email,
		PasswordHash: passwordHash,
	}
	body, err := json.Marshal(userRequest)
	if err != nil {
		return err
	}

	// Publish the user data to the User Service
	err = s.messagePublisher.Publish("auth_exchange", "auth_to_user_key", body)
	if err != nil {
		return err
	}

	// Listen for a response from the User Service
	msgs, err := s.messagePublisher.Consume("auth_response_queue")
	if err != nil {
		return err
	}

	for d := range msgs {
		var response param.RegisterUserResponse
		if err := json.Unmarshal(d.Body, &response); err != nil {
			log.Printf("Failed to unmarshal user response: %v", err)
			continue
		}

		if response.Success {
			return nil
		} else {
			return errors.New(response.Error)
		}
	}

	return errors.New("no response from user service")
}

func (s *AuthService) Login(ctx context.Context, loginRequest param.LoginRequest) (param.LoginResponse, error) {

	//* Fetch user data from UserService
	user, err := s.fetchUserData(ctx, loginRequest.UsernameOrEmail)
	if err != nil {
		return nil, err
	}

	//* Validate the password (compare hashed password with provided password)
	if !isValidPassword(loginRequest.Password, string(user.PasswordHash)) {
		return param.LoginResponse{}, errors.New("Invalid credentials")
	}

	//* Find the user by username or email based on the login request
	// user, err := s.userRepository.FindByUsernameOrEmail(ctx, loginRequest.UsernameOrEmail)
	// if err != nil {
	// 	return param.LoginResponse{}, err
	// }

	//* Generate tokens using the utils package
	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return param.LoginResponse{}, err
	}

	// Optionally, generate a refresh token as well
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return param.LoginResponse{}, err
	}

	// Convert tokens to entities.Token type
	tokens := s.convertTokens(user.ID, accessToken, refreshToken)
	fmt.Println(tokens)
	return user, nil
}

func (s *AuthService) convertTokens(userID int, tokenStrings ...string) []entities.Token {
	var tokens []entities.Token
	for _, tokenString := range tokenStrings {
		token := entities.Token{
			UserID:    userID,
			Token:     tokenString,
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().Add(24 * time.Hour), // Example expiration time
		}
		tokens = append(tokens, token)
	}
	return tokens
}

// isValidPassword checks if the provided password matches the hashed password.
func isValidPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateCorrelationID() string {
	// Implement a method to generate a unique correlation ID
	return "some-unique-correlation-id"
}

func (s *AuthService) fetchUserData(ctx context.Context, usernameOrEmail string) (param.UserResponse, error) {
	request := map[string]string{"usernameOrEmail": usernameOrEmail}
	body, err := json.Marshal(request)
	if err != nil {
		return param.UserResponse{}, err
	}

	err = s.messagePublisher.Publish("auth_exchange", "auth_to_user_key", body)
	if err != nil {
		return param.UserResponse{}, err
	}

	// Listen for response from User Service
	msgs, err := s.messagePublisher.Consume("auth_response_queue")
	if err != nil {
		return param.UserResponse{}, err
	}

	for d := range msgs {
		var user param.UserResponse
		if err := json.Unmarshal(d.Body, &user); err != nil {
			log.Printf("Failed to unmarshal user response: %v", err)
			continue
		}
		return &user, nil
	}

	return param.UserResponse{}, errors.New("no response from user service")
}

//! Login Method:
//? Fetch user data from the UserService via RabbitMQ.
//? Validate the password.
//? Generate tokens.
//? Return the user data with tokens.
