package services

// import (
// 	"authentication-service/interfaces"
// 	"authentication-service/internal/domain"
// 	"errors"
// 	"fmt"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"golang.org/x/crypto/bcrypt"
// )

// var (
// 	ErrInvalidCredentials = errors.New("invalid credentials")
// )

// type AuthService struct {
// 	userRepo       interfaces.UserRepository
// 	jwtKey         []byte
// 	eventPublisher interfaces.EventPublisher
// }

// func NewAuthService(userRepo interfaces.UserRepository, jwtKey []byte, eventPublisher interfaces.EventPublisher) *AuthService {
// 	return &AuthService{userRepo: userRepo, jwtKey: jwtKey, eventPublisher: eventPublisher}
// }

// func (s *AuthService) SignUp(username, password string) error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	user := &domain.User{
// 		ID:       generateID(),
// 		Username: username,
// 		Password: string(hashedPassword),
// 	}
// 	if err := s.userRepo.Save(user); err != nil {
// 		return err
// 	}

// 	// Publish event to message broker
// 	event := "UserSignedUp"
// 	body := fmt.Sprintf(`{"username": "%s"}`, user.Username)
// 	err = s.eventPublisher.Publish(event, []byte(body))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *AuthService) SignIn(username, password string) (string, error) {
// 	user, err := s.userRepo.FindByUsername(username)
// 	if err != nil {
// 		return "", ErrInvalidCredentials
// 	}
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		return "", ErrInvalidCredentials
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username": user.Username,
// 		"exp":      time.Now().Add(time.Hour * 72).Unix(),
// 	})
// 	return token.SignedString(s.jwtKey)
// }

// func generateID() string {
// 	return "some-unique-id" // replace with a proper ID generator
// }
