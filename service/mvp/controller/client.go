package controller

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"spectrum/common/env"
	"spectrum/common/pb"
)

var mvpClient pb.MvpClient

func initMvpClient() {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", env.Conf.Rpc.MvpServerPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logs.Error(err)
		return
	}
	mvpClient = pb.NewMvpClient(conn)
}
