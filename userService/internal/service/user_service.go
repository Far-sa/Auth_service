package service

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/entity"
	"user-service/internal/interfaces"
	"user-service/internal/param"
	"user-service/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserService struct {
	userRepo       interfaces.UserRepository
	eventPublisher interfaces.UserEvents
}

func NewUserService(userRepo interfaces.UserRepository, eventPublisher interfaces.UserEvents) *UserService {
	return &UserService{userRepo: userRepo, eventPublisher: eventPublisher}
}

//! Event Consumption in User Service: The user service listens to the UserRegisteredEvent and creates a new entry
//! in its user_profiles table. consume from "user_registered",   var event models.UserRegisteredEvent

func (s *UserService) Register(ctx context.Context, req param.RegisterRequest) (param.RegisterResponse, error) {
	// Hash the password
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return param.RegisterResponse{}, err
	}

	// Create a user request to send to the User Service
	userRequest := &entity.UserProfile{
		ID:          "",
		Username:    req.UserName,
		Email:       req.Email,
		Password:    passwordHash,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   time.Now(),
	}

	// Save user to the database (pseudo-code, replace with actual DB code)
	createdUser, cErr := s.userRepo.CreateUser(ctx, userRequest)
	if cErr != nil {
		return param.RegisterResponse{}, cErr
	}

	//TODO publish just user id
	s.publishUserIDEvent(ctx, createdUser.ID)

	//* marshal user data
	// data, err := json.Marshal(createdUser)
	// if err != nil {
	// 	return param.RegisterResponse{}, err
	// }

	//* Declare exchange and publish user as event

	// xErr := s.eventPublisher.DeclareExchange("user_events_exchange", "topic")
	// if xErr != nil {
	// 	return param.RegisterResponse{}, xErr
	// }

	// pErr := s.eventPublisher.Publish(ctx, "user_events_exchange", "user.created", amqp.Publishing{
	// 	ContentType: "application/json",
	// 	Body:        data,
	// })
	// if pErr != nil {
	// 	return param.RegisterResponse{}, pErr
	// }

	return param.RegisterResponse{User: param.UserInfo{ID: createdUser.ID, PhoneNumber: createdUser.PhoneNumber,
		Email: createdUser.Email, FullName: createdUser.Username}}, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, Email string) (param.UserInfo, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, Email)
	if err != nil {
		return param.UserInfo{}, nil
	}

	return param.UserInfo{ID: user.ID, Email: user.Email, FullName: user.FullName}, nil
}

func (s *UserService) GetUser(ctx context.Context, userID string) (param.UserProfileResponse, error) {
	userDetail, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return param.UserProfileResponse{}, err
	}
	userInfo := param.UserProfileResponse{
		UserProfile: entity.UserProfile{ID: userDetail.ID, Email: userDetail.Email, Username: userDetail.Username,
			CreatedAt: userDetail.CreatedAt},
	}
	return userInfo, nil
}

// !!!
// * publishUserIDEvent publishes an event with the user ID
func (s *UserService) publishUserIDEvent(ctx context.Context, userID string) error {
	const exchangeName = "user_events_exchange"
	const routingKey = "user.created"

	// Declare the exchange if not already declared
	err := s.eventPublisher.DeclareExchange(exchangeName, "topic")
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Publish the event
	err = s.eventPublisher.Publish(ctx, exchangeName, routingKey, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(fmt.Sprintf(`{"id":"%s"}`, userID)),
	})
	if err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	return nil
}

//!
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
