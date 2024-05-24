package interfaces

import "order-service/internal/domain/models"

type OrderRepository interface {
	CreateOrder(order models.Order) error
	GetOrder(orderID string) (models.Order, error)
	ListOrders(userID string) ([]models.Order, error)
}
