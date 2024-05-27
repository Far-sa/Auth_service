package interfaces

import "user-service/internal/entity"

type UserService interface {
	CreateUser(user entity.User) error
	GetUser(userID string) (entity.User, error)
}

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUser(userID string) (entity.User, error)
}
type Messaging interface {
	Publish(message []byte, queueName string) error
	Subscribe(queueName string, handler func(message []byte)) error
}
