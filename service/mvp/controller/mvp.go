package controller

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"test/common/pb"
	"test/service/mvp/dao"
	"time"
)

type MvpServer struct {
	pb.UnimplementedMvpServer
}

func (MvpServer) AddGood(ctx context.Context, req *pb.AddGoodReq) (*pb.AddGoodRes, error) {
	logs.Info("AddGood", ctx, req)

	var res pb.AddGoodRes

	// 1. 判断商品是否存在
	good, err := dao.GoodDao.GetByName(req.GoodName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodDao.GetByName")
	}
	if good != nil {
		return nil, errors.New("Good has existed")
	}

	// 2. 创建商品
	if err := dao.GoodDao.Create(req.GoodName, float64(req.Price)); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) SellGood(ctx context.Context, req *pb.SellGoodReq) (*pb.SellGoodRes, error) {
	logs.Info("SellGood", ctx, req)

	var res pb.SellGoodRes

	// 1. 判断商品是否存在
	good, err := dao.GoodDao.GetByName(req.GoodName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodDao.GetByName")
	}
	if good == nil {
		return nil, errors.New("Good not existed")
	}

	// 2. 创建销售记录
	if err := dao.SellRecordDao.Create(int(good.ID), float64(req.SellPrice)); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) AddBilliardDesk(ctx context.Context, req *pb.AddBilliardDeskReq) (*pb.AddBilliardDeskRes, error) {
	logs.Info("AddBilliardDesk", ctx, req)

	var res pb.AddBilliardDeskRes

	// 1. 判断桌名是否重复
	desk, err := dao.BilliardDeskDao.GetByName(req.BilliardDeskName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
	}
	if desk != nil {
		return nil, errors.New("Desk name dumplicate")
	}

	// 2. 创建台球桌
	if err := dao.BilliardDeskDao.Create(req.BilliardDeskName); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) BeginPlayBilliard(ctx context.Context, req *pb.BeginPlayBilliardReq) (*pb.BeginPlayBilliardRes, error) {
	logs.Info("BeginPlayBilliard", ctx, req)

	var res pb.BeginPlayBilliardRes

	// 1. 判断桌名是否存在
	desk, err := dao.BilliardDeskDao.GetByName(req.BilliardDeskName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
	}
	if desk == nil {
		return nil, errors.New("Desk not exist")
	}

	// 2. 判断记录是否存在
	record, err := dao.PlayRecordDao.Get(int(desk.ID), time.Unix(req.BeginPlayTimestamp, 0))
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish PlayRecordDao.Get")
	}
	if record != nil {
		return nil, errors.New("Record has exist")
	}

	// 3. 创建玩台球的记录
	if err := dao.PlayRecordDao.Create(
		int(desk.ID),
		time.Unix(req.BeginPlayTimestamp, 0),
	); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

func (MvpServer) StopPlayBilliard(ctx context.Context, req *pb.StopPlayBilliardReq) (*pb.StopPlayBilliardRes, error) {
	logs.Info("StopPlayBilliard", ctx, req)

	var res pb.StopPlayBilliardRes

	// 1. 判断桌名是否存在
	desk, err := dao.BilliardDeskDao.GetByName(req.BilliardDeskName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
	}
	if desk == nil {
		return nil, errors.New("Desk not exist")
	}

	// 2. 更新玩台球的记录
	if err := dao.PlayRecordDao.UpdateStopPlayTime(
		int(desk.ID),
		time.Unix(req.BeginPlayTimestamp, 0),
		time.Unix(req.StopPlayTimestamp, 0),
	); err != nil {
		logs.Error(err)
		return nil, err
	}
	return &res, nil
}

