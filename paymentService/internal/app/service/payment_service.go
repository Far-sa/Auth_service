package service

import (
	"payment-service/interfaces"
	"payment-service/internal/domain/models"
)

type PaymentService struct {
	paymentRepo interfaces.PaymentRepository
	messaging   interfaces.Messaging
}

func NewPaymentService(paymentRepo interfaces.PaymentRepository, messaging interfaces.Messaging) *PaymentService {
	return &PaymentService{paymentRepo: paymentRepo, messaging: messaging}
}

func (s *PaymentService) ProcessPayment(payment models.Payment) error {
	err := s.paymentRepo.CreatePayment(payment)
	if err != nil {
		return err
	}
	paymentProcessedMessage := []byte("Payment processed with ID: " + payment.PaymentID)
	return s.messaging.Publish(paymentProcessedMessage, "payment_processed")
}

func (s *PaymentService) GetPayment(paymentID string) (models.Payment, error) {
	return s.paymentRepo.GetPayment(paymentID)
}
