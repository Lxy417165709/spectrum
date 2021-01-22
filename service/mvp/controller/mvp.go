package controller

import (
	"context"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"time"
)

// todo: 考虑一下， desk 里面是否要包括 good 呢？还是将二者分开，之后再通过其他手段联系呢？
type MvpServer struct {
	pb.UnimplementedMvpServer
}

func (MvpServer) AddGood(ctx context.Context, req *pb.AddGoodReq) (*pb.AddGoodRes, error) {
	logger.Info("AddGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddGoodRes
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
	// todo: 这里可以点单记录
	var res pb.OrderGoodRes

	for _, good := range req.Goods {
		dbGood := &model.Good{
			Name:              good.MainElement.Name,
			DeskID:            req.DeskID,
			Expense:           good.ExpenseInfo.Expense,
			CheckOutTimestamp: good.ExpenseInfo.CheckOutTimestamp,
			NonFavorExpense:   good.ExpenseInfo.NonFavorExpense,
		}
		if err := dao.ChargeableObjectDao.Create(dbGood); err != nil {
			// todo: log
			return nil, err
		}
		if err := dao.ChargeableObjectDao.CreateFavorRecord(dbGood.GetName(), int64(dbGood.ID), good.Favors); err != nil {
			// todo: log
			return nil, err
		}
		if err := writeGoodSizeToDB(good); err != nil {
			logger.Error("Fail to finish createGood",
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
	}
	return &res, nil
}

func (MvpServer) OrderDesk(ctx context.Context, req *pb.OrderDeskReq) (*pb.OrderDeskRes, error) {
	logger.Info("OrderDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	// todo: 这里可以点单记录
	var res pb.OrderDeskRes

	dbDesk := &model.Desk{
		StartTimestamp:    time.Now().Unix(),
		EndTimestamp:      0,
		SpaceName:         req.Desk.Space.Name,
		SpaceNum:          req.Desk.Space.Num,
		Expense:           req.Desk.ExpenseInfo.Expense,
		CheckOutTimestamp: req.Desk.ExpenseInfo.CheckOutTimestamp,
		NonFavorExpense:   req.Desk.ExpenseInfo.NonFavorExpense,
	}
	if err := dao.ChargeableObjectDao.Create(dbDesk); err != nil {
		// todo: log
		return nil, err
	}
	if err := dao.ChargeableObjectDao.CreateFavorRecord(dbDesk.GetName(), int64(dbDesk.ID), req.Desk.Favors); err != nil {
		// todo: log
		return nil, err
	}
	res.DeskID = int64(dbDesk.ID)
	return &res, nil
}

func (MvpServer) GetDesk(ctx context.Context, req *pb.GetDeskReq) (*pb.GetDeskRes, error) {
	logger.Info("GetDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.GetDeskRes
	desk, err := dao.DeskDao.Get(req.DeskID)
	if err != nil {
		// todo: log
		return nil, err
	}
	res.Desk = getPbDesk(desk)
	res.Goods = getDeskPbGoods(int64(desk.ID))
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

func (MvpServer) CheckOut(ctx context.Context, req *pb.CheckOutReq) (*pb.CheckOutRes, error) {
	logger.Info("CheckOut", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.CheckOutRes

	for _, id := range req.GoodIDs {
		good, err := dao.GoodDao.Get(id)
		if err != nil {
			// todo: log
			return nil, err
		}
		if good == nil {
			// todo: warning
			continue
		}
		if err := checkOutIfNot(good); err != nil {
			// todo: log
			return nil, err
		}
	}

	for _, id := range req.DeskIDs {
		desk, err := dao.DeskDao.Get(id)
		if err != nil {
			// todo: log
			return nil, err
		}
		if desk == nil {
			// todo: warning
			continue
		}
		if err := checkOutIfNot(desk); err != nil {
			// todo: log
			return nil, err
		}
	}

	return &res, nil
}
