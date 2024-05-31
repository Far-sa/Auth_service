package services

import (
	"authentication-service/domain/entities"
	"authentication-service/domain/param"
	"authentication-service/interfaces"
	"authentication-service/utils"
	"context"
	"encoding/json"
	"errors"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticationService interface defines methods for user authentication

// AuthService implements the AuthenticationService interface
type AuthService struct {
	authRepository   interfaces.AuthRepository
	messagePublisher interfaces.AuthEvents
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(authRepository interfaces.AuthRepository, messagePublisher interfaces.AuthEvents) *AuthService {
	return &AuthService{
		authRepository:   authRepository,
		messagePublisher: messagePublisher,
	}
}

// ! Event Publication: The authentication service publishes a UserRegisteredEvent to RabbitMQ.
//!     event := models.UserRegisteredEvent / user_registered

// TODO update method to return message
func (s *AuthService) Register(ctx context.Context, registerRequest param.RegisterRequest) error {
	// Hash the password
	passwordHash, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		return err
	}

	// Create a user request to send to the User Service
	userRequest := &entities.User{
		ID:           "",
		Username:     registerRequest.Username,
		Email:        registerRequest.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}

	body, err := json.Marshal(userRequest)
	if err != nil {
		return err
	}

	// Save user to the database (pseudo-code, replace with actual DB code)
	sErr := s.authRepository.SaveUser(ctx, userRequest)
	if sErr != nil {
		return sErr
	}

	// Publish the user data to the User Service
	err = s.messagePublisher.Publish(ctx, "auth_exchange", "auth_to_user_key", amqp.Publishing{Body: body})
	if err != nil {
		return err
	}

	return errors.New("no response from user service")
}

func (s *AuthService) Login(ctx context.Context, loginRequest param.LoginRequest) (param.LoginResponse, error) {

	//! if receive data from grpc server in user service
	// Retrieve user information from UserService
	// userReq := &userpb.GetUserRequest{
	//     Email: req.Email,
	// }
	// userResp, err := s.userClient.GetUser(ctx, userReq)
	// if err != nil {
	//     return nil, status.Errorf(codes.Internal, "could not get user: %v", err)
	// }

	//* get user data from database and compare passwords
	user, err := s.authRepository.FindByUserEmail(ctx, loginRequest.Email)
	if err != nil {
		return param.LoginResponse{}, err
	}

	//* Validate the password (compare hashed password with provided password)
	if !isValidPassword(loginRequest.Password, string(user.PasswordHash)) {
		return param.LoginResponse{}, errors.New("invalid credentials")
	}

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

	tokens := &entities.TokenPair{
		AccessToken:  entities.AccessToken{Token: accessToken, ExpiresAt: time.Now().Add(24 * time.Hour)},
		RefreshToken: entities.RefreshToken{Token: refreshToken, ExpiresAt: time.Now().Add(7 * 24 * time.Hour)},
	}

	s.authRepository.SaveToken(ctx, tokens)

	return param.LoginResponse{UserID: user.ID, TokenPair: param.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}}, nil
}

// func (s *AuthService) convertTokens(userID string, tokenStrings ...string) []entities.TokenPair {
// 	var tokens []entities.TokenPair
// 	for _, tokenString := range tokenStrings {
// 		token := entities.TokenPair{
// 			AccessToken:  entities.AccessToken{Token: tokenString, ExpiresAt: time.Now().Add(24 * time.Hour)},
// 			RefreshToken: entities.RefreshToken{Token: tokenString, ExpiresAt: time.Now().Add(7 * 24 * time.Hour)},
// 		}
// 		tokens = append(tokens, token)
// 	}
// 	return tokens
// }

// isValidPassword checks if the provided password matches the hashed password.
func isValidPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateCorrelationID() string {
	// Implement a method to generate a unique correlation ID
	return "some-unique-correlation-id"
}

// func (s *AuthService) fetchUserData(ctx context.Context, usernameOrEmail string) (param.UserResponse, error) {
// 	request := map[string]string{"usernameOrEmail": usernameOrEmail}
// 	body, err := json.Marshal(request)
// 	if err != nil {
// 		return param.UserResponse{}, err
// 	}

// 	err = s.messagePublisher.Publish("auth_exchange", "auth_to_user_key", body)
// 	if err != nil {
// 		return param.UserResponse{}, err
// 	}

// 	// Listen for response from User Service
// 	msgs, err := s.messagePublisher.Consume("auth_response_queue")
// 	if err != nil {
// 		return param.UserResponse{}, err
// 	}

// 	for d := range msgs {
// 		var user param.UserResponse
// 		if err := json.Unmarshal(d.Body, &user); err != nil {
// 			log.Printf("Failed to unmarshal user response: %v", err)
// 			continue
// 		}
// 		return &user, nil
// 	}

// 	return param.UserResponse{}, errors.New("no response from user service")
// }
