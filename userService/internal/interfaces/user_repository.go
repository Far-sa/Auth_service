package interfaces

import "user-service/internal/entity"

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUser(userID string) (entity.User, error)
}
