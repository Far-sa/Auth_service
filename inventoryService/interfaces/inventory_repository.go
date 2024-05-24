package interfaces

import "inventory-service/internal/domain/models"

type InventoryRepository interface {
	UpdateStock(productID string, quantity int32) error
	CheckStock(productID string) (models.Inventory, error)
}
