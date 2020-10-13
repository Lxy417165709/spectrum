package controller

import (
	"context"
	"github.com/astaxie/beego/logs"
	"test/common/pb"
	"testing"
)

func TestServer_Login(t *testing.T) {
	req := &pb.LoginReq{
		Email:    "417165709@qq.com",
		Password: "123456",
	}
	res, err := serverClient.Login(context.Background(), req)
	logs.Info(res, err)
	if err != nil {
		t.Fatal(err)
	}
}
