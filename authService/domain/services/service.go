package services

import (
	"authentication-service/domain/entities"
	"authentication-service/domain/param"
	"authentication-service/interfaces"
	"authentication-service/utils"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// AuthenticationService interface defines methods for user authentication

// AuthService implements the AuthenticationService interface
type AuthService struct {
	userRepository   interfaces.UserRepository
	messagePublisher interfaces.MessagePublisher

	// Optional: event publisher for authentication events
	// eventPublisher event_publisher.EventPublisher
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepository interfaces.UserRepository, messagePublisher interfaces.MessagePublisher) *AuthService {
	return &AuthService{
		userRepository:   userRepository,
		messagePublisher: messagePublisher,
	}
}

func (s *AuthService) Login(ctx context.Context, loginRequest param.LoginRequest) (*entities.User, error) {
	// Find the user by username or email based on the login request
	user, err := s.userRepository.FindByUsernameOrEmail(ctx, loginRequest.UsernameOrEmail)
	if err != nil {
		return nil, err
	}

	// Validate the password (compare hashed password with provided password)
	if !isValidPassword(loginRequest.Password, string(user.PasswordHash)) {
		return nil, errors.New("Invalid credentials")
	}

	//TODO generate tokens- use utils package
	// Generate tokens using the utils package
	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	// Optionally, generate a refresh token as well
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	// Set the tokens to the user entity or a custom auth response if required
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	return user, nil
}

func (s *AuthService) Register(ctx context.Context, req param.RegisterRequest) error {
	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new user entity
	user := &entities.User{
		Username:     req.Username,
		PasswordHash: hashedPassword,
		// Populate other fields as necessary
	}

	// Use the user repository to store the registered user
	err = s.userRepository.Save(ctx, user)
	if err != nil {
		return err
	}

	// Optionally, publish a UserRegisteredEvent after successful registration
	if s.messagePublisher != nil {
		err := s.messagePublisher.PublishUserAuthenticated(user.ID)
		if err != nil {
			return err
		}
	}

	return nil

}

// isValidPassword checks if the provided password matches the hashed password.
func isValidPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
