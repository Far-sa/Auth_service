package service

import (
	"inventory-service/interfaces"
	"inventory-service/internal/domain/models"
)

type InventoryService struct {
	inventoryRepo interfaces.InventoryRepository
	messaging     interfaces.Messaging
}

func NewInventoryService(inventoryRepo interfaces.InventoryRepository, messaging interfaces.Messaging) *InventoryService {
	return &InventoryService{inventoryRepo: inventoryRepo, messaging: messaging}
}

func (s *InventoryService) UpdateStock(productID string, quantity int32) error {
	err := s.inventoryRepo.UpdateStock(productID, quantity)
	if err != nil {
		return err
	}
	stockUpdatedMessage := []byte("Stock updated for product ID: " + productID)
	return s.messaging.Publish(stockUpdatedMessage, "stock_updated")
}

func (s *InventoryService) CheckStock(productID string) (models.Inventory, error) {
	return s.inventoryRepo.CheckStock(productID)
}
