package interfaces

import "payment-service/internal/domain/models"

type PaymentRepository interface {
	CreatePayment(payment models.Payment) error
	GetPayment(paymentID string) (models.Payment, error)
}
