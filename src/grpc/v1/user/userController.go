package user

import (
	db "bjm/db/benjamit"
	"bjm/proto/v1/user"
	"bjm/utils"
	"context"
)

type UserServer struct {
	user.UnimplementedUserServer
}

func (s *UserServer) Login(ctx context.Context, reqModel *user.LoginRequestModel) (*user.LoginResponseModel, error) {
	context, contextErr := db.Connect()
	defer db.ConnectClose(context)
	if contextErr != nil {
		return utils.GrpcResponseErrorJson(&user.LoginResponseModel{}, contextErr.Error(), 13)
	}

	resModel := &user.LoginResponseModel{}
	service := &UserService{context}
	serviceRes := service.Login(reqModel, resModel)

	return utils.GrpcResponseJson(serviceRes, int(serviceRes.StatusCode))
}
