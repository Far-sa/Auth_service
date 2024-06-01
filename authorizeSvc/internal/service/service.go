package service

// import (
// 	"authorization-service/infrastructure/messaging"
// 	"authorization-service/internal/interfaces"
// 	"context"
// 	"encoding/json"
// 	"log"
// )

// type AuthorizationService struct {
// 	roleRepository interfaces.RoleRepository
// 	consumer       *messaging.RabbitMQConsumer
// 	// publisher      interfaces.EventPublisher
// }

// func NewAuthorizationService(roleRepository interfaces.RoleRepository, consumer *messaging.RabbitMQConsumer) *AuthorizationService {
// 	return &AuthorizationService{roleRepository: roleRepository, consumer: consumer}
// }

// type Event struct {
// 	Type EventType // The type of the event (e.g., USER_CREATED, USER_UPDATED)
// 	User *User     // The user that the event pertains to
// }

// type EventType int

// // ! events
// func (s *AuthorizationService) Start() error {
// 	if err := s.consumer.StartConsuming(); err != nil {
// 		log.Fatalf("Failed to start consumer: %v", err)
// 	}

// 	// Process messages in a separate goroutine
// 	go func() {
// 		for msg := range msgs {
// 			event := &user.Event{}
// 			if err := json.Unmarshal(msg.Body, event); err != nil {
// 				log.Printf("Failed to unmarshal message: %v", err)
// 				continue
// 			}

// 			switch event.Type {
// 			case user.Event_USER_CREATED:
// 				// Handle user created event
// 				s.handleUserCreated(event)
// 			case user.Event_USER_ROLE_UPDATED:
// 				// Handle user role updated event
// 				s.handleUserRoleUpdated(event)
// 			}
// 		}
// 	}()
// 	return nil
// }

// func (s *AuthorizationService) handleUserCreated(event *user.Event) {
// 	// Extract user data from event
// 	newUser := event.User

// 	// Assign default role to new user
// 	defaultRole := "user" // replace with your actual default role
// 	err := s.assignRoleToUser(newUser.Id, defaultRole)
// 	if err != nil {
// 		log.Printf("Failed to assign role to new user: %v", err)
// 		return
// 	}

// 	// Update internal state with new user's role
// 	// ...
// }

// func (s *AuthorizationService) handleUserRoleUpdated(event *user.Event) {
// 	// Update internal state with updated user's role
// 	// ...

// 	// Extract user data from event
// 	updatedUser := event.User

// 	// Update internal state with updated user's role
// 	err := s.updateUserRoleInState(updatedUser.Id, updatedUser.Role)
// 	if err != nil {
// 		log.Printf("Failed to update user's role: %v", err)
// 		return
// 	}
// }

// // ! events
// func (s *AuthorizationService) HandleUserAuthenticatedEvent(ctx context.Context, message string) error {
// 	// Extract user information from the message
// 	userID := extractUserIDFromMessage(message)
// 	log.Printf("Handling UserAuthenticated event for user: %s", userID)

// 	// Update user roles/permissions in the repository
// 	return s.roleRepository.UpdateUserRoles(ctx, userID, "new_role")
// }

// func (s *AuthorizationService) assignRoleToUser(ctx context.Context, username, role string) error {
// 	err := s.roleRepository.AssignRole(ctx, username, role)
// 	if err != nil {
// 		return err
// 	}

// 	// event := map[string]string{"username": username, "role": role}
// 	// eventData, _ := json.Marshal(event)
// 	// err = s.publisher.Publish("role_assigned", eventData)
// 	// if err != nil {
// 	// 	log.Printf("Failed to publish event: %v", err)
// 	// }

// 	return nil
// }
// func (s *AuthorizationService) updateUserRoleInState(userId, role string) error {
// 	// Implement the logic to update a user's role in the internal state
// 	// This might involve updating a database or an in-memory data structure, for example
// 	// ...
// 	return nil
// }

// func (s *AuthorizationService) CheckPermission(ctx context.Context, username, permission string) (bool, error) {
// 	return s.roleRepository.CheckPermission(ctx, username, permission)
// }

// func extractUserIDFromMessage(message string) string {
// 	// Simulating extraction logic
// 	return message
// }
