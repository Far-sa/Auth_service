package service

import (
	"user-service/interfaces"
	"user-service/internal/domain/models"
)

type UserService struct {
	userRepo  interfaces.UserRepository
	messaging interfaces.Messaging
}

func NewUserService(userRepo interfaces.UserRepository, messaging interfaces.Messaging) *UserService {
	return &UserService{userRepo: userRepo, messaging: messaging}
}

func (s *UserService) CreateUser(user models.User) error {
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	userCreatedMessage := []byte("User created with ID: " + user.UserID)
	return s.messaging.Publish(userCreatedMessage, "user_created")
}

func (s *UserService) GetUser(userID string) (models.User, error) {
	return s.userRepo.GetUser(userID)
}
