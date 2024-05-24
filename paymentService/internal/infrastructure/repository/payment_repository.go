package repository

import (
	"database/sql"
	"payment-service/interfaces"
	"payment-service/internal/domain/models"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) interfaces.PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) CreatePayment(payment models.Payment) error {
	query := `INSERT INTO payments (payment_id, order_id, amount, method, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, payment.PaymentID, payment.OrderID, payment.Amount, payment.Method, payment.CreatedAt)
	return err
}

func (r *PaymentRepository) GetPayment(paymentID string) (models.Payment, error) {
	var payment models.Payment
	query := `SELECT payment_id, order_id, amount, method, created_at FROM payments WHERE payment_id = $1`
	row := r.db.QueryRow(query, paymentID)
	err := row.Scan(&payment.PaymentID, &payment.OrderID, &payment.Amount, &payment.Method, &payment.CreatedAt)
	return payment, err
}
