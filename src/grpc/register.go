package grpc

import (
	userPd "bjm/proto/v1/user"
	"bjm/src/grpc/v1/user"

	gRpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Register(server *gRpc.Server) {
	reflection.Register(server)
	userPd.RegisterUserServer(server, &user.UserServer{})
}
