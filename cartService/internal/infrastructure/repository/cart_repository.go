package database

import (
	"cart-service/internal/domain/models"
	"database/sql"
)

type PostgresCartRepository struct {
	db *sql.DB
}

func NewPostgresCartRepository(db *sql.DB) *PostgresCartRepository {
	return &PostgresCartRepository{db: db}
}

func (r *PostgresCartRepository) AddToCart(userID, productID string, quantity int32) error {
	query := `
    INSERT INTO cart (user_id, product_id, quantity)
    VALUES ($1, $2, $3)
    ON CONFLICT (user_id, product_id)
    DO UPDATE SET quantity = cart.quantity + EXCLUDED.quantity;
    `
	_, err := r.db.Exec(query, userID, productID, quantity)
	return err
}

func (r *PostgresCartRepository) RemoveFromCart(userID, productID string) error {
	query := `
    DELETE FROM cart WHERE user_id = $1 AND product_id = $2;
    `
	_, err := r.db.Exec(query, userID, productID)
	return err
}

func (r *PostgresCartRepository) GetCart(userID string) ([]models.CartItem, error) {
	query := `
    SELECT product_id, quantity FROM cart WHERE user_id = $1;
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItem
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *PostgresCartRepository) ClearCart(userID string) error {
	query := `
    DELETE FROM cart WHERE user_id = $1;
    `
	_, err := r.db.Exec(query, userID)
	return err
}
