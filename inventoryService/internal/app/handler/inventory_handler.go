package handler

import (
	"context"
	"inventory-service/internal/app/service"
	pb "inventory-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type InventoryHandler struct {
	inventoryService *service.InventoryService
	pb.UnimplementedInventoryServiceServer
}

func NewInventoryHandler(inventoryService *service.InventoryService) *InventoryHandler {
	return &InventoryHandler{inventoryService: inventoryService}
}

func (h *InventoryHandler) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.UpdateStockResponse, error) {
	err := h.inventoryService.UpdateStock(req.ProductId, req.Quantity)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateStockResponse{Success: true}, nil
}

func (h *InventoryHandler) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
	inventory, err := h.inventoryService.CheckStock(req.ProductId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckStockResponse{
		ProductId:   inventory.ProductID,
		Quantity:    inventory.Quantity,
		LastUpdated: timestamppb.New(inventory.LastUpdated),
	}, nil
}
