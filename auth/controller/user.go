package controller

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"test/common/pb"
)

type Server struct {
	pb.UnimplementedServerServer
}

func (Server) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	logs.Info("Login", ctx, req)
	var res pb.LoginRes
	user,err := userDao.GetByEmail(req.Email)
	if err != nil{
		logs.Error(err)
		return nil,errors.New("fail to finish GetByEmail")
	}
	if user == nil{
		return nil,errors.New("email or password error")
	}
	return &res, nil
}
