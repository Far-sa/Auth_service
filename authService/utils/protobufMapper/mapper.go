package mapper

import (
	"authentication-service/domain/param"
	"authentication-service/pb"
)

func ToProtoLoginRequest(logReq param.LoginRequest) *pb.LoginRequest {
	return &pb.LoginRequest{
		Email:    logReq.Email,
		Password: logReq.Password,
	}
}

func ToParamLoginResponse(protoResp *pb.LoginResponse) param.LoginResponse {
	return param.LoginResponse{
		UserID:    protoResp.UserId,
		TokenPair: protoResp.AccessToken,
	}
}
