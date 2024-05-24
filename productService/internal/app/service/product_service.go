package service

import (
	"product-service/interfaces"
	"product-service/internal/domain"
)

type ProductService struct {
	repo interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) CreateProduct(name, description string, price float64) (*domain.Product, error) {
	product := &domain.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
	err := s.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProduct(id int64) (*domain.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *ProductService) UpdateProduct(id int64, name, description string, price float64) (*domain.Product, error) {
	product, err := s.repo.GetProduct(id)
	if err != nil {
		return nil, err
	}
	product.Name = name
	product.Description = description
	product.Price = price
	err = s.repo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(id int64) error {
	return s.repo.DeleteProduct(id)
}

func (s *ProductService) ListProducts() ([]*domain.Product, error) {
	return s.repo.ListProducts()
}
