package controller

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"strings"
	"time"
)

type MvpServer struct {
	pb.UnimplementedMvpServer
}

func (MvpServer) AddSellGood(ctx context.Context, req *pb.AddSellGoodReq) (*pb.AddSellGoodRes, error) {
	logger.Info("AddSellGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddSellGoodRes
	if err := createElement(req.Good.MainElement, req.ClassName); err != nil {
		logger.Error("Fail to finish createElement",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	if err := writeGoodSizeToDB(req.Good); err != nil {
		logger.Error("Fail to finish writeGoodSizeToDB",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) AddElement(ctx context.Context, req *pb.AddElementReq) (*pb.AddElementRes, error) {
	logger.Info("AddElement", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddElementRes

	if err := createElement(req.Element, req.ClassName); err != nil {
		logger.Error("Fail to finish createElement",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) GetAllGoodClasses(ctx context.Context, req *pb.GetAllGoodClassesReq) (*pb.GetAllGoodClassesRes, error) {
	logger.Info("GetAllGoodClasses", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllGoodClassesRes

	// 1. 获得主元素的所有类
	classes, err := dao.ElementClassDao.GetAllClasses()
	if err != nil {
		logger.Error("Fail to finish ElementClassDao.GetMainElementClass",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 形成 GoodClasses
	for _, class := range classes {
		res.GoodClasses = append(res.GoodClasses, &pb.GoodClass{
			Name:  class.Name,
			Goods: getClassGoods(class.Name),
		})
	}

	return &res, nil
}

func (MvpServer) AddGoodClass(ctx context.Context, req *pb.AddGoodClassReq) (*pb.AddGoodClassRes, error) {
	logger.Info("AddGoodClass", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddGoodClassRes
	// todo: 判断类名是否为空、是否存在

	// 1. 创建商品类
	if err := dao.ElementClassDao.Create(&model.ElementClass{
		Name: req.GoodClass.Name,
	}); err != nil {
		logger.Error("Fail to finish ElementDao.Create",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) OrderGood(ctx context.Context, req *pb.OrderGoodReq) (*pb.OrderGoodRes, error) {
	logger.Info("OrderGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.OrderGoodRes

	for _, good := range req.Goods {
		// 生成货物编号, 将货物与桌位联结
		dbGood := &model.Good{
			Name:   good.MainElement.Name,
			DeskID: req.DeskID,
		}
		if err := dao.GoodDao.Create(dbGood); err != nil {
			// todo:log
			return nil, err
		}
		good.Id = int64(dbGood.ID)
		if err := writeGoodSizeToDB(good); err != nil {
			logger.Error("Fail to finish createGood",
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
		if err := writeFavorToDB(good); err != nil {
			// todo: log
			return nil, err
		}

	}
	return &res, nil

}

func (MvpServer) OpenDesk(ctx context.Context, req *pb.OpenDeskReq) (*pb.OpenDeskRes, error) {
	logger.Info("OpenDesk", zap.Any("ctx", ctx), zap.Any("req", req))

	// todo: 通过 SpaceName Num 查询 Price PriceRuleType
	var res pb.OpenDeskRes

	desk := &model.Desk{
		SpaceName:      req.SpaceName,
		SpaceNum:       int(req.SpaceNum),
		StartTimestamp: time.Now().Unix(),
	}
	if err := dao.DeskDao.Create(desk); err != nil {
		//todo: log
		return nil, err
	}
	if err := writeFavorToDB(desk); err != nil {
		// todo: log
		return nil, err
	}
	res.DeskID = int64(desk.ID)
	return &res, nil
}

func (MvpServer) GetDesk(ctx context.Context, req *pb.GetDeskReq) (*pb.GetDeskRes, error) {
	logger.Info("GetDesk", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetDeskRes

	res.Desk = getDesk(req.DeskID)
	return &res, nil
}

func (MvpServer) AddSpace(ctx context.Context, req *pb.AddSpaceReq) (*pb.AddSpaceRes, error) {
	logger.Info("AddDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.AddSpaceRes

	// 1. 创建
	if err := dao.SpaceDao.Create(getDbSpace(req.Space)); err != nil {
		logger.Error("Fail to finish SpaceDao.Create",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) CloseDesk(ctx context.Context, req *pb.CloseDeskReq) (*pb.CloseDeskRes, error) {
	logger.Info("CloseDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.CloseDeskRes
	if err := closeDeskIfOpening(req.DeskID, req.EndTimestamp); err != nil {
		logger.Error("Fail to finish closeDesk",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

// todo: 这个接口设计的不太好
func (MvpServer) FormExpense(ctx context.Context, req *pb.FormExpenseReq) (*pb.FormExpenseRes, error) {
	// todo: 折扣信息还未记录

	logger.Info("FormExpense", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.FormExpenseRes
	desk := getDesk(req.DeskID)
	if desk == nil {
		err := fmt.Errorf("desk(id = %d) is non", req.DeskID)
		logger.Error("Desk is non", zap.Error(err))
		return nil, err
	}
	formDeskExpense(desk)
	writeToDB(desk, "expense")
	res.Desk = desk
	return &res, nil
}

func (MvpServer) CheckOut(ctx context.Context, req *pb.CheckOutReq) (*pb.CheckOutRes, error) {
	// todo: 形成订单
	logger.Info("CheckOut", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.CheckOutRes

	writeToDB(req.Desk, "had_check_out")
	return &res, nil
}
