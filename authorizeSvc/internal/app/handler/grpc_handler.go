package handler

import (
	"authorization-service/internal/app/service"
	"authorization-service/pb"
	"context"
)

type AuthzHandler struct {
	pb.UnimplementedAuthorizationServiceServer
	service *service.AuthzService
}

func NewAuthzHandler(service *service.AuthzService) *AuthzHandler {
	return &AuthzHandler{service: service}
}

func (h *AuthzHandler) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.AssignRoleResponse, error) {
	err := h.service.AssignRole(req.Username, req.Role)
	if err != nil {
		return nil, err
	}
	return &pb.AssignRoleResponse{Message: "Role assigned successfully"}, nil
}

func (h *AuthzHandler) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	hasPermission, err := h.service.CheckPermission(req.Username, req.Permission)
	if err != nil {
		return nil, err
	}
	//return &pb.CheckPermissionResponse{Has_permission: hasPermission}, nil
	return &pb.CheckPermissionResponse{HasPermission: hasPermission}, nil
}
