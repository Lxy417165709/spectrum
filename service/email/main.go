package main

import (
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"net"
	"test/common/env"
	"test/common/pb"
	"test/service/email/controller"
)

const confFilePath = "configure/alpha.json"

func initLogger() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
}

func initConfigure() {
	if err := env.Conf.Load(confFilePath); err != nil {
		panic(err)
	}
}

func main() {
	initLogger()
	initConfigure()
	logs.Info("Server starting.")
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		logs.Error(err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEmailServer(grpcServer, &controller.EmailServer{})
	logs.Info("Server ready to serve.")
	if err := grpcServer.Serve(lis); err != nil {
		logs.Error(err)
		return
	}
}
