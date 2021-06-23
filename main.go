package main

import (
	"net"

	"google.golang.org/grpc"

	. "github.com/ts-new/model"
	. "github.com/ts-new/user_service"
	. "github.com/ts-new/utils/log"
)

const (
	port = ":8081"
)

func main() {
	Log.Init()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		Log.FATAL("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterUserServiceServer(s, &UserService{})

	Log.INFO("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		Log.FATAL("failed to serve: %v", err)
	}
}
