package service

import (
	"authorization-service/interfaces"
	"encoding/json"
	"log"
)

type AuthorizationService struct {
	roleRepository interfaces.RoleRepository
	publisher      interfaces.EventPublisher
}

func NewAuthorizationService(roleRepository interfaces.RoleRepository, publisher interfaces.EventPublisher) *AuthorizationService {
	return &AuthorizationService{roleRepository: roleRepository, publisher: publisher}
}

func (s *AuthorizationService) HandleUserAuthenticatedEvent(message string) error {
	// Extract user information from the message
	userID := extractUserIDFromMessage(message)
	log.Printf("Handling UserAuthenticated event for user: %s", userID)

	// Update user roles/permissions in the repository
	return s.roleRepository.UpdateUserRoles(userID, "new_role")
}

func (s *AuthorizationService) AssignRole(username, role string) error {
	err := s.roleRepository.AssignRole(username, role)
	if err != nil {
		return err
	}

	event := map[string]string{"username": username, "role": role}
	eventData, _ := json.Marshal(event)
	err = s.publisher.Publish("role_assigned", eventData)
	if err != nil {
		log.Printf("Failed to publish event: %v", err)
	}

	return nil
}

func (s *AuthorizationService) CheckPermission(username, permission string) (bool, error) {
	return s.roleRepository.CheckPermission(username, permission)
}

func extractUserIDFromMessage(message string) string {
	// Simulating extraction logic
	return message
}
