package handler

import (
	"cart-service/internal/app/service"
	pb "cart-service/pb"
	"context"
)

type CartHandler struct {
	pb.UnimplementedCartServiceServer
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	err := h.service.AddToCart(req.UserId, req.ProductId, req.Quantity)
	if err != nil {
		return &pb.AddToCartResponse{Success: false, Message: err.Error()}, err
	}
	return &pb.AddToCartResponse{Success: true, Message: "Product added to cart"}, nil
}

func (h *CartHandler) RemoveFromCart(ctx context.Context, req *pb.RemoveFromCartRequest) (*pb.RemoveFromCartResponse, error) {
	err := h.service.RemoveFromCart(req.UserId, req.ProductId)
	if err != nil {
		return &pb.RemoveFromCartResponse{Success: false, Message: err.Error()}, err
	}
	return &pb.RemoveFromCartResponse{Success: true, Message: "Product removed from cart"}, nil
}

func (h *CartHandler) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	items, err := h.service.GetCart(req.UserId)
	if err != nil {
		return nil, err
	}
	var cartItems []*pb.CartItem
	for _, item := range items {
		cartItems = append(cartItems, &pb.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}
	return &pb.GetCartResponse{Items: cartItems}, nil
}
