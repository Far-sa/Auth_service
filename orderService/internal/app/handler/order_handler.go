package handler

import (
	"context"
	"order-service/internal/app/service"
	"order-service/internal/domain/models"
	pb "order-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderHandler struct {
	orderService *service.OrderService
	pb.UnimplementedOrderServiceServer
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := models.Order{
		OrderID:     req.OrderId,
		UserID:      req.UserId,
		Status:      req.Status,
		TotalAmount: req.TotalAmount,
		CreatedAt:   req.CreatedAt.AsTime(),
	}
	err := h.orderService.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{OrderId: order.OrderID}, nil
}

func (h *OrderHandler) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := h.orderService.GetOrder(req.OrderId)
	if err != nil {
		return nil, err
	}
	return &pb.GetOrderResponse{
		OrderId:     order.OrderID,
		UserId:      order.UserID,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		CreatedAt:   timestamppb.New(order.CreatedAt),
	}, nil
}

func (h *OrderHandler) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := h.orderService.ListOrders(req.UserId)
	if err != nil {
		return nil, err
	}
	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			OrderId:     order.OrderID,
			UserId:      order.UserID,
			Status:      order.Status,
			TotalAmount: order.TotalAmount,
			CreatedAt:   timestamppb.New(order.CreatedAt),
		})
	}
	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}
