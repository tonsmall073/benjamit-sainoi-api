package user

import (
	db "bjm/db/benjamit"
	v1 "bjm/proto/v1"
	"context"
)

type UserServer struct {
	v1.UnimplementedUserServer
}

func (s *UserServer) Login(ctx context.Context, reqModel *v1.LoginRequestModel) (*v1.LoginResponseModel, error) {
	context, contextErr := db.Connect()
	defer db.ConnectClose(context)
	if contextErr != nil {
		return &v1.LoginResponseModel{MessageDesc: contextErr.Error(), StatusCode: 500}, nil
	}

	resModel := &v1.LoginResponseModel{}
	service := &UserService{context}
	serviceRes := service.Login(reqModel, resModel)

	return serviceRes, nil
}
