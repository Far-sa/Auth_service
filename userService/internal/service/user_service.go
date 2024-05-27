package service

import (
	"context"
	"encoding/json"
	"log"
	"user-service/internal/entity"
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

func (s *UserService) ListenForUserRequests() {
	msgs, err := s.messaging.Consume("user_queue")
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %v", err)
	}

	go func() {
		for d := range msgs {
			var req map[string]string
			if err := json.Unmarshal(d.Body, &req); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			usernameOrEmail, ok := req["usernameOrEmail"]
			if !ok {
				log.Printf("Invalid message format: %v", req)
				continue
			}

			user, err := s.userRepo.FindByUsernameOrEmail(context.Background(), usernameOrEmail)
			if err != nil {
				log.Printf("User not found: %v", usernameOrEmail)
				continue
			}

			// Publish the user details back to an auth service queue or reply mechanism
			response, err := json.Marshal(user)
			if err != nil {
				log.Printf("Failed to marshal user response: %v", err)
				continue
			}

			s.messaging.Publish("user_exchange", "user_response_key", response)
		}
	}()
}

func (s *UserService) GetUserByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entity.User, error) {
	return s.userRepo.FindByUsernameOrEmail(ctx, usernameOrEmail)
}

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

//!!!

// func (s *UserService) StartListening() {
// 	// Listen for registration messages
// 	go s.listenForRegistration()
// 	// Listen for user data requests
// 	go s.listenForUserDataRequests()
// }

// func (s *UserService) listenForRegistration() {
// 	msgs, err := s.rabbitMQ.Consume("user_registration_queue")
// 	if err != nil {
// 		log.Fatalf("Failed to start consuming registration messages: %v", err)
// 	}

// 	for d := range msgs {
// 		var req RegisterUserRequest
// 		if err := json.Unmarshal(d.Body, &req); err != nil {
// 			log.Printf("Failed to unmarshal registration message: %v", err)
// 			continue
// 		}

// 		user := entities.User{
// 			Username:     req.Username,
// 			Email:        req.Email,
// 			PasswordHash: req.PasswordHash,
// 			CreatedAt:    time.Now(),
// 			UpdatedAt:    time.Now(),
// 		}

// 		err := s.userRepository.CreateUser(context.Background(), &user)
// 		var response RegisterUserResponse
// 		if err != nil {
// 			response = RegisterUserResponse{
// 				Success: false,
// 				Error:   err.Error(),
// 			}
// 		} else {
// 			response = RegisterUserResponse{
// 				Success: true,
// 			}
// 		}

// 		responseBody, err := json.Marshal(response)
// 		if err != nil {
// 			log.Printf("Failed to marshal registration response: %v", err)
// 			continue
// 		}

// 		s.rabbitMQ.Publish("user_exchange", "user_registration_response_key", responseBody)
// 	}
// }

// func (s *UserService) listenForUserDataRequests() {
// 	msgs, err := s.rabbitMQ.Consume("user_data_request_queue")
// 	if err != nil {
// 		log.Fatalf("Failed to start consuming user data request messages: %v", err)
// 	}

// 	for d := range msgs {
// 		var req map[string]string
// 		if err := json.Unmarshal(d.Body, &req); err != nil {
// 			log.Printf("Failed to unmarshal user data request message: %v", err)
// 			continue
// 		}

// 		usernameOrEmail, ok := req["usernameOrEmail"]
// 		if !ok {
// 			log.Printf("Invalid user data request message format: %v", req)
// 			continue
// 		}

// 		user, err := s.userRepository.FindByUsernameOrEmail(context.Background(), usernameOrEmail)
// 		var response UserResponse
// 		if err != nil {
// 			response = UserResponse{
// 				ID:           0,
// 				PasswordHash: "",
// 				Error:        "User not found",
// 			}
// 		} else {
// 			response = UserResponse{
// 				ID:           user.ID,
// 				PasswordHash: user.PasswordHash,
// 			}
// 		}

// 		responseBody, err := json.Marshal(response)
// 		if err != nil {
// 			log.Printf("Failed to marshal user data response: %v", err)
// 			continue
// 		}

// 		s.rabbitMQ.Publish("user_exchange", "user_data_response_key", responseBody)
// 	}
// }
