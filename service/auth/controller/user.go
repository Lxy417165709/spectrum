package controller

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"test/common/pb"
	"test/common/utils"
	"test/service/auth/dao"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func (AuthServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	logs.Info("Login", ctx, req)
	var res pb.LoginRes

	// 1. 判断用户是否存在
	user, err := dao.UserDao.GetByEmail(req.Email)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish userDao.GetByEmail")
	}
	if user == nil {
		return nil, errors.New("User not exist")
	}

	// 2. 判断密码是否正确
	hashSaltyPassword, err := utils.GetHashString(utils.GetSaltyPassword(req.Password, user.Salt))
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to get hash string")
	}
	if user.HashSaltyPassword != hashSaltyPassword {
		return nil, errors.New("Password error")
	}
	return &res, nil
}

func (AuthServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	logs.Info("Register", ctx, req)
	var res pb.RegisterRes

	// 1. 判断注册邮箱是否存在
	user, err := dao.UserDao.GetByEmail(req.Email)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish userDao.GetByEmail")
	}
	if user != nil {
		return nil, errors.New("Email has exist")
	}

	// 2. 创建用户
	hashSaltyPassword, salt, err := utils.GetHashSaltyPassword(req.Password)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to get hash salty password")
	}

	if err := dao.UserDao.Create(req.Email, hashSaltyPassword, salt); err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish userDao.Create")
	}
	return &res, nil
}

