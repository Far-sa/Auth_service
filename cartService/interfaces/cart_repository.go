package interfaces

import "cart-service/internal/domain/models"

type CartRepository interface {
	AddToCart(userID, productID string, quantity int32) error
	RemoveFromCart(userID, productID string) error
	GetCart(userID string) ([]models.CartItem, error)
	ClearCart(userID string) error
}
