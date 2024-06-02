package mapper

import (
	"authorization-service/internal/param"
	authz "authorization-service/pb"
)

func PbToParamAssignRoleRequest(req *authz.AssignRoleRequest) param.RoleAssignmentRequest {
	return param.RoleAssignmentRequest{
		UserID: req.UserId,
	}
}

func ToPbAssignRoleResponse(r *param.RoleAssignmentRequest) *authz.AssignRoleResponse {
	return &authz.AssignRoleResponse{
		Message: "Role assigned successfully",
	}
}

func PbToParamUpdateRoleRequest(req *authz.UpdateRoleRequest) param.RoleUpdateRequest {
	return param.RoleUpdateRequest{
		UserID:  req.UserId,
		NewRole: req.Role,
	}
}

func ToPbUpdateRoleResponse(r *param.RoleUpdateRequest) *authz.UpdateRoleResponse {
	return &authz.UpdateRoleResponse{
		Message: "Role updated successfully",
	}
}
