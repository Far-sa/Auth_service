package interfaces

import (
	"user-service/internal/entity"
	"user-service/internal/param"
)

type UserService interface {
	// CreateUser(user entity.UserDetail) error
	GetUser(userID string) (param.UserInfo, error)
}

type UserRepository interface {
	// CreateUser(user entity.UserDetail) error
	GetUser(userID string) (entity.UserDetail, error)
}
type Messaging interface {
	Publish(message []byte, queueName string) error
	Subscribe(queueName string, handler func(message []byte)) error
}
