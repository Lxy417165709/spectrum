package main

import (
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"net"
	"test/auth/controller"
	"test/common/pb"
)

func main() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)

	logs.Info("Server beginning.")
	lis, err := net.Listen("tcp", ":8087")
	if err != nil {
		logs.Error(err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterServerServer(grpcServer, &controller.Server{})
	logs.Info("Server starting.")
	if err := grpcServer.Serve(lis); err != nil {
		logs.Error(err)
		return
	}
}
