package service

import (
	"authorization-service/interfaces"
	"encoding/json"
	"log"
)

type AuthzService struct {
	roleRepo  interfaces.RoleRepository
	publisher interfaces.EventPublisher
}

func NewAuthzService(roleRepo interfaces.RoleRepository, publisher interfaces.EventPublisher) *AuthzService {
	return &AuthzService{roleRepo, publisher}
}

func (s *AuthzService) AssignRole(username, role string) error {
	err := s.roleRepo.AssignRole(username, role)
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

func (s *AuthzService) CheckPermission(username, permission string) (bool, error) {
	return s.roleRepo.CheckPermission(username, permission)
}
