package handler

import (
	"authorization-service/internal/interfaces"
	"authorization-service/pb"
	"context"
)

type AuthzHandler struct {
	pb.UnimplementedAuthorizationServiceServer
	service interfaces.AuthorizationService
}

func NewAuthzHandler(service interfaces.AuthorizationService) *AuthzHandler {
	return &AuthzHandler{service: service}
}

func (h *AuthzHandler) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.AssignRoleResponse, error) {
	err := h.service.AssignRole(ctx, req.Username, req.Role)
	if err != nil {
		return nil, err
	}
	return &pb.AssignRoleResponse{Message: "Role assigned successfully"}, nil
}

func (h *AuthzHandler) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	hasPermission, err := h.service.CheckPermission(ctx, req.Username, req.Permission)
	if err != nil {
		return nil, err
	}
	//return &pb.CheckPermissionResponse{Has_permission: hasPermission}, nil
	return &pb.CheckPermissionResponse{HasPermission: hasPermission}, nil
}
