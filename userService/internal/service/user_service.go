package service

import (
	"user-service/internal/entity"
	"user-service/internal/interfaces"
)

type UserService struct {
	userRepo  interfaces.UserRepository
	messaging interfaces.Messaging
}

func NewUserService(userRepo interfaces.UserRepository, messaging interfaces.Messaging) *UserService {
	return &UserService{userRepo: userRepo, messaging: messaging}
}

func (s *UserService) CreateUser(user entity.User) error {
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	userCreatedMessage := []byte("User created with ID: " + user.UserID)
	return s.messaging.Publish(userCreatedMessage, "user_created")
}

func (s *UserService) GetUser(userID string) (entity.User, error) {
	return s.userRepo.GetUser(userID)
}
