package mapper

import (
	"user-service/internal/entity"
	"user-service/internal/param"
	user "user-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EntityToProtoUser(u *entity.UserProfile) *user.GetUserResponse {
	return &user.GetUserResponse{
		UserId:    u.ID,
		Name:      u.FullName,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt.String()),
	}
}

func ProtoToEntityUser(protoUser *user.GetUserResponse) *entity.UserProfile {
	return &entity.UserProfile{
		ID:       protoUser.UserId,
		FullName: protoUser.Name,
		Email:    protoUser.Email,
	}
}
func PbToParamGetUser(req *user.GetUserRequest) param.