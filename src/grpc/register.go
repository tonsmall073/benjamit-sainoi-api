package grpc

import (
	proto "bjm/proto/v1"
	"bjm/src/grpc/v1/user"

	gRpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Register(server *gRpc.Server) {
	reflection.Register(server)
	proto.RegisterUserServer(server, &user.UserServer{})
}
