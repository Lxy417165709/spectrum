package controller

import (
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"test/common/pb"
)

var authClient pb.AuthClient

func initAuthClient() {
	conn, err := grpc.Dial(":8087", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logs.Error(err)
		return
	}
	authClient = pb.NewAuthClient(conn)
}
