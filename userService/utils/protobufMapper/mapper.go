package mapper

import (
	"user-service/internal/entity"
	"user-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EntityToProtoUser(u *entity.UserProfile) *pb.GetUserResponse {
	return &pb.GetUserResponse{
		UserId:    u.ID,
		Name:      u.FullName,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt.String()),
	}
}

func ProtoToEntityUser(protoUser *pb.GetUserResponse) *entity.UserProfile {
	return &entity.UserProfile{
		ID:       protoUser.UserId,
		FullName: protoUser.Name,
		Email:    protoUser.Email,
	}
}
