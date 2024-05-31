package mapper

import (
	"authentication-service/domain/param"
	auth "authentication-service/pb"
	user "user-service/pb"
)

func PbToParamLoginRequest(protoReq *auth.LoginRequest) param.LoginRequest {
	return param.LoginRequest{
		Email:    protoReq.Email,
		Password: protoReq.Password,
	}
}

func ParamToPbLoginResponse(paramResp param.LoginResponse) *auth.LoginResponse {
	return &auth.LoginResponse{
		UserId:       paramResp.UserID,
		AccessToken:  paramResp.Tokens.RefreshToken,
		RefreshToken: paramResp.Tokens.RefreshToken,
	}
}

func ToProtoGetUserRequest(paramRes param.LoginRequest) *user.GetUserByEmailRequest {
	return &user.GetUserByEmailRequest{
		Email: paramRes.Email,
	}
}

func ToParamGetUserResponse(userResp *user.GetUserByEmailResponse) param.GetUserResponse {
	return param.GetUserResponse{
		Id:       userResp.UserId,
		Password: userResp.Password,
	}
}
