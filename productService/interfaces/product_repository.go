package interfaces

import "product-service/internal/domain"

type ProductRepository interface {
	CreateProduct(product *domain.Product) error
	GetProduct(id int64) (*domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id int64) error
	ListProducts() ([]*domain.Product, error)
}
