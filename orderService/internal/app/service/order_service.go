package service

import (
	"order-service/interfaces"
	"order-service/internal/domain/models"
)

type OrderService struct {
	orderRepo interfaces.OrderRepository
	messaging interfaces.Messaging
}

func NewOrderService(orderRepo interfaces.OrderRepository, messaging interfaces.Messaging) *OrderService {
	return &OrderService{orderRepo: orderRepo, messaging: messaging}
}

func (s *OrderService) CreateOrder(order models.Order) error {
	err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return err
	}
	orderCreatedMessage := []byte("Order created with ID: " + order.OrderID)
	err = s.messaging.Publish(orderCreatedMessage, "order_created")
	return err
}

func (s *OrderService) GetOrder(orderID string) (models.Order, error) {
	return s.orderRepo.GetOrder(orderID)
}

func (s *OrderService) ListOrders(userID string) ([]models.Order, error) {
	return s.orderRepo.ListOrders(userID)
}
