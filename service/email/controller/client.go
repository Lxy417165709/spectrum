package controller

import (
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"test/common/pb"
)

var emailClient pb.EmailClient

func initEmailClient() {
	conn, err := grpc.Dial(":8088", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logs.Error(err)
		return
	}
	emailClient = pb.NewEmailClient(conn)
}
