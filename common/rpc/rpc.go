package rpc

import (
	"github.com/astaxie/beego/logs"
	"google.golang.org/grpc"
	"spectrum/common/pb"
)

var (
	MvpClient pb.MvpClient
)

func init() {
	conn, err := grpc.Dial(":8089", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logs.Error(err)
		return
	}
	MvpClient = pb.NewMvpClient(conn)
}
