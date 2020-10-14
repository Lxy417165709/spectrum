package controller

import (
	"context"
	"github.com/astaxie/beego/logs"
	"test/common/pb"
	"testing"
)

func TestEmailServer_Send(t *testing.T) {
	initEmailClient()
	req := &pb.SendEmailReq{
		To:      "417165709@qq.com",
		Subject: "test",
		Content: "<h1>hello</h1>",
	}
	res, err := emailClient.Send(context.Background(), req)
	logs.Info(res, err)
	if err != nil {
		t.Fatal(err)
	}
}
