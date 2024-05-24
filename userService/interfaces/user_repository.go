package interfaces

import "user-service/internal/domain/models"

type UserRepository interface {
	CreateUser(user models.User) error
	GetUser(userID string) (models.User, error)
}
