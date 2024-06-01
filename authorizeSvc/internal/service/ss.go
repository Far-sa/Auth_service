package service

// import "log"

// func (s *AuthorizationService) handleUserCreated(event *Event) {
// 	// Create a new role for the user
// 	role := &Role{
// 		UserId: event.User.Id,
// 		// Set the role's properties based on the event
// 	}

// 	// Save the role to the repository
// 	if err := s.roleRepository.Save(role); err != nil {
// 		log.Printf("Failed to save role: %v", err)
// 	}
// }

// func (s *AuthorizationService) handleUserRoleUpdated(event *Event) {
// 	// Get the existing role for the user
// 	role, err := s.roleRepository.GetByUserId(event.User.Id)
// 	if err != nil {
// 		log.Printf("Failed to get role: %v", err)
// 		return
// 	}

// 	// Update the role's properties based on the event
// 	// ...

// 	// Save the updated role to the repository
// 	if err := s.roleRepository.Save(role); err != nil {
// 		log.Printf("Failed to save role: %v", err)
// 	}
// }

// import (
//     "encoding/json"
//     "log"

//     "github.com/streadway/amqp"
// )

// type AuthorizationService struct {
//     // ... other fields ...
//     conn *amqp.Connection
//     ch   *amqp.Channel
// }

// func (s *AuthorizationService) Start() error {
//     var err error

//     // Connect to RabbitMQ
//     s.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
//     if err != nil {
//         return err
//     }

//     // Create a channel
//     s.ch, err = s.conn.Channel()
//     if err != nil {
//         return err
//     }

//     // Declare a queue
//     q, err := s.ch.QueueDeclare(
//         "user.events", // name
//         true,          // durable
//         false,         // delete when unused
//         false,         // exclusive
//         false,         // no-wait
//         nil,           // arguments
//     )
//     if err != nil {
//         return err
//     }

//     // Consume messages from the queue
//     msgs, err := s.ch.Consume(
//         q.Name, // queue
//         "",     // consumer
//         true,   // auto-ack
//         false,  // exclusive
//         false,  // no-local
//         false,  // no-wait
//         nil,    // args
//     )
//     if err != nil {
//         return err
//     }

//     // Process messages in a separate goroutine
//     go s.processMessages(msgs)

//     return nil
// }

// func (s *AuthorizationService) processMessages(msgs <-chan amqp.Delivery) {
//     for d := range msgs {
//         var event user.Event
//         if err := json.Unmarshal(d.Body, &event); err != nil {
//             log.Printf("Failed to unmarshal event: %v", err)
//             continue
//         }

//         switch event.Type {
//         case user.USER_CREATED:
//             s.handleUserCreated(&event)
//         case user.USER_UPDATED:
//             s.handleUserRoleUpdated(&event)
//         default:
//             log.Printf("Unknown event type: %v", event.Type)
//         }
//     }
// }

// ... handleUserCreated and handleUserRoleUpdated methods ...
