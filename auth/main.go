package main

import (
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"net"
	"test/auth/controller"
	"test/common/pb"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		logs.Error(err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterServerServer(grpcServer, &controller.Server{})
	if err := grpcServer.Serve(lis); err != nil {
		logs.Error(err)
	}
}
