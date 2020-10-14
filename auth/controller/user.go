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
	user, err := userDao.GetByEmail(req.Email)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish userDao.GetByEmail")
	}
	if user == nil {
		return nil, errors.New("User not exist")
	}
	if user.Password != req.Password{
		return nil,errors.New("Password error")
	}
	return &res, nil
}

func (Server) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	logs.Info("Register", ctx, req)
	var res pb.RegisterRes
	if err := userDao.Create(req.Email, req.Password); err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish userDao.Create")
	}
	return &res, nil
}
