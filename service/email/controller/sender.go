package controller

import (
	"context"
	"github.com/astaxie/beego/logs"
	"test/common/env"
	"test/common/pb"
	"test/service/email/model"
)

type EmailServer struct {
	pb.UnimplementedEmailServer
}

func (EmailServer) Send(ctx context.Context, req *pb.SendEmailReq) (*pb.SendEmailRes, error) {
	logs.Info("Send", ctx, req)

	var res pb.SendEmailRes
	if err := send(&model.EmailContext{
		EmailAddrOfSender:    env.Conf.EmailClient.EmailAddr,
		EmailAddrOfReceivers: []string{req.To},
		Subject:              req.Subject,
		Body:                 req.Content,
		Type:                 model.HtmlType,
	}); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}
