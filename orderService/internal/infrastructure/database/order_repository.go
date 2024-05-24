package database

import (
	"database/sql"
	"order-service/internal/domain/models"
)

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) CreateOrder(order models.Order) error {
	query := `
    INSERT INTO orders (order_id, user_id, status, total_amount, created_at)
    VALUES ($1, $2, $3, $4, $5);
    `
	_, err := r.db.Exec(query, order.OrderID, order.UserID, order.Status, order.TotalAmount, order.CreatedAt)
	return err
}

func (r *PostgresOrderRepository) GetOrder(orderID string) (models.Order, error) {
	query := `
    SELECT order_id, user_id, status, total_amount, created_at FROM orders WHERE order_id = $1;
    `
	var order models.Order
	err := r.db.QueryRow(query, orderID).Scan(&order.OrderID, &order.UserID, &order.Status, &order.TotalAmount, &order.CreatedAt)
	return order, err
}

func (r *PostgresOrderRepository) ListOrders(userID string) ([]models.Order, error) {
	query := `
    SELECT order_id, user_id, status, total_amount, created_at FROM orders WHERE user_id = $1;
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderID, &order.UserID, &order.Status, &order.TotalAmount, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
