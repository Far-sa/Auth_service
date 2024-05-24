package handler

import (
	"context"
	"product-service/internal/app/service"
	"product-service/internal/domain"
	pb "product-service/pb"
	"time"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := h.service.CreateProduct(req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{Product: convertToProtoProduct(product)}, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := h.service.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductResponse{Product: convertToProtoProduct(product)}, nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	product, err := h.service.UpdateProduct(req.Id, req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProductResponse{Product: convertToProtoProduct(product)}, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := h.service.DeleteProduct(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProductResponse{Success: true}, nil
}

func (h *ProductHandler) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := h.service.ListProducts()
	if err != nil {
		return nil, err
	}
	var protoProducts []*pb.Product
	for _, product := range products {
		protoProducts = append(protoProducts, convertToProtoProduct(product))
	}
	return &pb.ListProductsResponse{Products: protoProducts}, nil
}

func convertToProtoProduct(product *domain.Product) *pb.Product {
	return &pb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}
}
