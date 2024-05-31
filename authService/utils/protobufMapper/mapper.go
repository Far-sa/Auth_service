package mapper

import (
	"authentication-service/domain/param"
	"authentication-service/pb"
)

func PbToParamLoginRequest(protoReq *pb.LoginRequest) param.LoginRequest {
	return param.LoginRequest{
		Email:    protoReq.Email,
		Password: protoReq.Password,
	}
}

func ParamToPbLoginResponse(paramResp param.LoginResponse) *pb.LoginResponse {
	return &pb.LoginResponse{
		UserId:       paramResp.UserID,
		AccessToken:  paramResp.Tokens.RefreshToken,
		RefreshToken: paramResp.Tokens.RefreshToken,
	}
}
