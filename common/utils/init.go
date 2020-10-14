package utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"net"
	"test/common/env"
	"test/common/pb"
)

func InitLogger() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
}

func InitConfigure(confFilePath string) {
	if err := env.Conf.Load(confFilePath); err != nil {
		panic(err)
	}
}

func InitRpc(servePort int, server interface{}) {
	logs.Info("Server starting.")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", servePort))
	if err != nil {
		logs.Error(err)
		return
	}
	grpcServer := grpc.NewServer()
	switch server.(type) {
	case pb.EmailServer:
		pb.RegisterEmailServer(grpcServer, server.(pb.EmailServer))
	case pb.AuthServer:
		pb.RegisterAuthServer(grpcServer, server.(pb.AuthServer))
	default:
		panic(fmt.Sprintf("Invalid rpc server(%+v).", server))
	}
	logs.Info("Server ready to serve.")
	if err := grpcServer.Serve(lis); err != nil {
		logs.Error(err)
		return
	}
}
