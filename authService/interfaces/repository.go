package interfaces

import "authentication-service/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}
