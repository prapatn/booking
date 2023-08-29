package config

import (
	"booking/cmd/server/request"
	"booking/pb"

	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterRoomsServer(s, &request.RoomsServer{})

	return s
}
