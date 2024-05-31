package mapper

import (
	"user-service/internal/param"
	user "user-service/pb"

	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// func EntityToProtoUser(u *entity.UserProfile) *user.GetUserResponse {
// 	return &user.GetUserResponse{
// 		UserId:    u.ID,
// 		Name:      u.FullName,
// 		Email:     u.Email,
// 		CreatedAt: timestamppb.New(u.CreatedAt.String()),
// 	}
// }

// func ProtoToEntityUser(protoUser *user.GetUserResponse) *entity.UserProfile {
// 	return &entity.UserProfile{
// 		ID:       protoUser.UserId,
// 		FullName: protoUser.Name,
// 		Email:    protoUser.Email,
// 	}
// }

func PbToParamGetUserRequest(req *user.GetUserRequest) param.GetUser {
	return param.GetUser{
		UserID: req.UserId,
	}
}

func ToPbUserProfileResponse(u param.UserProfileResponse) *user.GetUserResponse {
	createdAt, _ := time.Parse(time.RFC3339, u.UserProfile.CreatedAt.String())
	return &user.GetUserResponse{
		UserId:    u.UserProfile.ID,
		Name:      u.UserProfile.FullName,
		Email:     u.UserProfile.Email,
		CreatedAt: timestamppb.New(createdAt),
	}
}
