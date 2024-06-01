package service

import (
	"authorization-service/internal/entity"
	"authorization-service/internal/interfaces"
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

func (s *AuthzService) AssignRole(ctx context.Context, userID string) error {
	//TODO  get role from db
	role, err := s.roleRepo.GetRoleByUserID(userID)
	if err != nil {
		if errors.Is(err, interfaces.ErrRoleNotFound) {
			log.Printf("No role found for user %s, assigning default role", userID)
			role = entity.Role{Name: DefaultRole}
		} else {
			log.Printf("Failed to assign role: %v", err)
			return err
		}
	}

	// log.Printf("Assigned role %s to user %s with email %s", role.Name, user.ID, user.Email)
	fmt.Println("role is :", role)
	return nil

	//*  Logic to assign role to user
	// role := "user"
	// err := s.roleRepo.AssignRole(ctx, userID, role)
	// if err != nil {
	// 	return err
	// }

	// log.Printf("Assigned role to user %s", userID)
	// return nil
}

//TODO use param

func (s *AuthzService) UpdateUserRole(userID string, newRole entity.Role) error {
	err := s.roleRepo.UpdateRole(userID, newRole)
	if err != nil {
		log.Printf("Failed to update role for user %s: %v", userID, err)
		return err
	}
	log.Printf("Updated role for user %s to %s", userID, newRole.Name)
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
			if err := s.AssignRole(context.Background(), userID); err != nil {
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
