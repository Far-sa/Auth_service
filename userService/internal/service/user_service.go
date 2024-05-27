package service

import (
	"user-service/internal/interfaces"
	"user-service/internal/param"
)

type UserService struct {
	userRepo  interfaces.UserRepository
	messaging interfaces.Messaging
}

func NewUserService(userRepo interfaces.UserRepository, messaging interfaces.Messaging) *UserService {
	return &UserService{userRepo: userRepo, messaging: messaging}
}

// func (s *UserService) CreateUser(user entity.User) error {
// 	err := s.userRepo.CreateUser(user)
// 	if err != nil {
// 		return err
// 	}
// 	userCreatedMessage := []byte("User created with ID: " + user.UserID)
// 	return s.messaging.Publish(userCreatedMessage, "user_created")
// }

func (s *UserService) GetUser(userID string) (param.UserInfo, error) {
	userDetail, err := s.userRepo.GetUser(userID)
	if err != nil {
		return param.UserInfo{}, err
	}
	userInfo := param.UserInfo{
		ID:          userDetail.UserID,
		PhoneNumber: userDetail.PhoneNumber,
		Name:        userDetail.FirstName,
		// Set other fields as needed
	}
	return userInfo, nil
}
