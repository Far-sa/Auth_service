package service

import (
	"authorization-service/pb"
)

type AuthorizationServer struct {
	pb.UnimplementedAuthorizationServiceServer
	roleRepository *RoleRepository
}

func NewAuthorizationServer(roleRepo *RoleRepository) *AuthorizationServer {
	return &AuthorizationServer{roleRepository: roleRepo}
}

// Implement gRPC methods
