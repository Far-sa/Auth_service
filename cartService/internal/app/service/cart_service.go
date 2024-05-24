package service

import (
	"cart-service/interfaces"
	"cart-service/internal/domain/models"
)

type CartService struct {
	repository interfaces.CartRepository
	messaging  interfaces.Messaging
}

func NewCartService(repository interfaces.CartRepository, messaging interfaces.Messaging) *CartService {
	return &CartService{
		repository: repository,
		messaging:  messaging,
	}
}

func (s *CartService) AddToCart(userID, productID string, quantity int32) error {
	return s.repository.AddToCart(userID, productID, quantity)
}

func (s *CartService) RemoveFromCart(userID, productID string) error {
	return s.repository.RemoveFromCart(userID, productID)
}

func (s *CartService) GetCart(userID string) ([]models.CartItem, error) {
	return s.repository.GetCart(userID)
}

func (s *CartService) ClearCart(userID string) error {
	return s.repository.ClearCart(userID)
}
