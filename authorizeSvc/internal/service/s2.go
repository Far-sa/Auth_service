package service

import (
	"authorization-service/infrastructure/messaging/rabbitmq"
	"authorization-service/internal/interfaces"
	"context"
	"encoding/json"
	"log"
)

type AuthzService struct {
	roleRepo interfaces.RoleRepository
	consumer rabbitmq.RabbitMQAdapter
	// publisher      interfaces.EventPublisher
}

func NewAuthzService(roleRepo interfaces.RoleRepository, consumer rabbitmq.RabbitMQAdapter) *AuthzService {
	return &AuthzService{roleRepo: roleRepo, consumer: consumer}
}

func (s *AuthzService) AssignRole(ctx context.Context, userID string) error {
	// Logic to assign role to user
	role := "user"
	err := s.roleRepo.AssignRole(ctx, userID, role)
	if err != nil {
		return err
	}

	log.Printf("Assigned role to user %s", userID)
	return nil
}

func (s *AuthzService) ListenForUserEvents() error {

	s.consumer.DeclareExchange("user_events_exchange", "topic")
	q, _ := s.consumer.CreateQueue("user_created_queue")
	s.consumer.CreateBinding(q.Name, "", "user_events_exchange")
	msgs, _ := s.consumer.Consume(q.Name)

	go func() {
		for d := range msgs {
			var user dto.UserDTO
			err := json.Unmarshal(d.Body, &user)
			if err != nil {
				log.Printf("Error parsing message: %s", err)
				continue
			}

			// Assign role to the user
			s.AssignRole(ctx, user.ID)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	select {}
}
