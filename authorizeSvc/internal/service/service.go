package service

import (
	"authorization-service/infrastructure/messaging"
	"authorization-service/internal/interfaces"
	"context"
	"log"
)

type AuthorizationService struct {
	roleRepository interfaces.RoleRepository
	consumer       *messaging.RabbitMQConsumer
	// publisher      interfaces.EventPublisher
}

func NewAuthorizationService(roleRepository interfaces.RoleRepository, consumer *messaging.RabbitMQConsumer) *AuthorizationService {
	return &AuthorizationService{roleRepository: roleRepository, consumer: consumer}
}

func (s *AuthorizationService) Start() error {
	if err := s.consumer.StartConsuming(); err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	return nil
}

func (s *AuthorizationService) HandleUserAuthenticatedEvent(ctx context.Context, message string) error {
	// Extract user information from the message
	userID := extractUserIDFromMessage(message)
	log.Printf("Handling UserAuthenticated event for user: %s", userID)

	// Update user roles/permissions in the repository
	return s.roleRepository.UpdateUserRoles(ctx, userID, "new_role")
}

func (s *AuthorizationService) AssignRole(ctx context.Context, username, role string) error {
	err := s.roleRepository.AssignRole(ctx, username, role)
	if err != nil {
		return err
	}

	// event := map[string]string{"username": username, "role": role}
	// eventData, _ := json.Marshal(event)
	// err = s.publisher.Publish("role_assigned", eventData)
	// if err != nil {
	// 	log.Printf("Failed to publish event: %v", err)
	// }

	return nil
}

func (s *AuthorizationService) CheckPermission(ctx context.Context, username, permission string) (bool, error) {
	return s.roleRepository.CheckPermission(ctx, username, permission)
}

func extractUserIDFromMessage(message string) string {
	// Simulating extraction logic
	return message
}
