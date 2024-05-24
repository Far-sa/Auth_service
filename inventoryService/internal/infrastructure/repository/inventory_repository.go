package database

import (
	"database/sql"
	"inventory-service/interfaces"
	"inventory-service/internal/domain/models"
)

type InventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) interfaces.InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) UpdateStock(productID string, quantity int32) error {
	query := `UPDATE inventory SET quantity = $1, last_updated = NOW() WHERE product_id = $2`
	_, err := r.db.Exec(query, quantity, productID)
	return err
}

func (r *InventoryRepository) CheckStock(productID string) (models.Inventory, error) {
	var inventory models.Inventory
	query := `SELECT product_id, quantity, last_updated FROM inventory WHERE product_id = $1`
	row := r.db.QueryRow(query, productID)
	err := row.Scan(&inventory.ProductID, &inventory.Quantity, &inventory.LastUpdated)
	return inventory, err
}
