package controller

import (
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"spectrum/common/pb"
)

var mvpClient pb.MvpClient

func initMvpClient() {
	conn, err := grpc.Dial(":8089", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logs.Error(err)
		return
	}
	mvpClient = pb.NewMvpClient(conn)
}
