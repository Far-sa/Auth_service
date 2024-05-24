package handler

import (
	"context"
	"payment-service/internal/app/service"
	"payment-service/internal/domain/models"
	pb "payment-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
	pb.UnimplementedPaymentServiceServer
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	payment := models.Payment{
		PaymentID: req.PaymentId,
		OrderID:   req.OrderId,
		Amount:    req.Amount,
		Method:    req.Method,
		CreatedAt: req.CreatedAt.AsTime(),
	}
	err := h.paymentService.ProcessPayment(payment)
	if err != nil {
		return nil, err
	}
	return &pb.ProcessPaymentResponse{PaymentId: payment.PaymentID}, nil
}

func (h *PaymentHandler) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	payment, err := h.paymentService.GetPayment(req.PaymentId)
	if err != nil {
		return nil, err
	}
	return &pb.GetPaymentResponse{
		PaymentId: payment.PaymentID,
		OrderId:   payment.OrderID,
		Amount:    payment.Amount,
		Method:    payment.Method,
		CreatedAt: timestamppb.New(payment.CreatedAt),
	}, nil
}
