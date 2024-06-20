package service

import (
	"authorization-service/internal/entity"
	"authorization-service/internal/interfaces"
	"authorization-service/internal/param"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const DefaultRole = "User"

type AuthzService struct {
	roleRepo interfaces.RoleRepository
	consumer interfaces.RoleEvents
	// publisher      interfaces.EventPublisher
}

func NewAuthzService(roleRepo interfaces.RoleRepository, consumer interfaces.RoleEvents) *AuthzService {
	return &AuthzService{roleRepo: roleRepo, consumer: consumer}
}

func (s *AuthzService) AssignRole(ctx context.Context, req param.RoleAssignmentRequest) error {
	// Attempt to retrieve the existing role for the user
	role, err := s.roleRepo.GetRoleByUserID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, interfaces.ErrRoleNotFound) {
			// If no role is found, assign the default role
			err = s.roleRepo.AssignRole(ctx, req.UserID, DefaultRole)
			if err != nil {
				// Handle error when assigning a new role fails
				log.Printf("Failed to assign default role to user %s: %v", req.UserID, err)
				return err
			}
			// Log successful default role assignment
			log.Printf("Assigned default role to user %s", req.UserID)
		} else {
			// Log and return any other error that occurred while getting the role
			log.Printf("Failed to get role for user %s: %v", req.UserID, err)
			return err
		}
	} else {
		// If the user already has a role, there's nothing to do
		log.Printf("User %s already has a role assigned: %s", req.UserID, role.Name)
	}

	// The role assignment was successful, or the user already had a role
	return nil
}

func (s *AuthzService) UpdateUserRole(ctx context.Context, req param.RoleUpdateRequest) error {

	// Check if the new role is different from the current one
	currentRole, err := s.roleRepo.GetRoleByUserID(ctx, req.UserID)
	if err != nil {
		// Handle error, possibly not found or db error
		return fmt.Errorf("error retrieving current role for user %s: %w", req.UserID, err)
	}

	if currentRole.Name == req.NewRole {
		// No update needed, return early
		return nil
	}

	newRole := entity.Role{Name: req.NewRole}
	err = s.roleRepo.UpdateRole(ctx, currentRole.ID, newRole)
	if err != nil {
		return fmt.Errorf("failed to update role for user %s: %w", req.UserID, err)
	}

	log.Printf("Updated role for user %s to %s", req.UserID, newRole.Name)
	return nil
}

func (s *AuthzService) ListenForUserEvents() error {
	// Declare the exchange
	if err := s.consumer.DeclareExchange("user_events_exchange", "topic"); err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Create the queue
	q, err := s.consumer.CreateQueue("user_created_queue", true, false)
	if err != nil {
		return fmt.Errorf("failed to create queue: %w", err)
	}

	// Create the binding
	if err := s.consumer.CreateBinding(q.Name, "usr.created.*", "user_events_exchange"); err != nil {
		return fmt.Errorf("failed to create binding: %w", err)
	}

	// Consume messages from the queue
	msgs, err := s.consumer.Consume(q.Name, "authorize_svc", false)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %w", err)
	}

	//* go s.processMessages(msgs)
	// Process messages in a separate goroutine
	go func() {
		for d := range msgs {
			var userID string
			if err := json.Unmarshal(d.Body, &userID); err != nil {
				log.Printf("Error parsing message: %s", err)
				continue
			}

			// Assign role to the user
			if err := s.AssignRole(context.Background(), param.RoleAssignmentRequest{UserID: userID}); err != nil {
				log.Printf("Error assigning role to user: %s", err)
				continue
			}
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")

	// Use a channel to wait for a signal to stop listening for messages
	stopChan := make(chan struct{})
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		close(stopChan)
	}()

	<-stopChan // Wait here until told to stop
	log.Printf("Shutting down user event listener")

	return nil
}
