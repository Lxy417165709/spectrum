package controller

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"test/common/pb"
	"test/service/mvp/dao"
)

type MvpServer struct {
	pb.UnimplementedMvpServer
}

func (MvpServer) AddGood(ctx context.Context, req *pb.AddGoodReq) (*pb.AddGoodRes, error) {
	logs.Info("AddGood", ctx, req)

	var res pb.AddGoodRes

	// 1. 判断商品是否存在
	user, err := dao.GoodDao.GetByName(req.GoodName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodDao.GetByName")
	}
	if user != nil {
		return nil, errors.New("Good has existed")
	}

	// 2. 创建商品
	if err := dao.GoodDao.Create(req.GoodName, float64(req.Price)); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}
