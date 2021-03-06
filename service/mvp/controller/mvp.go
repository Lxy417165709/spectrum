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

// todo: GetOrder 可以设置为一个查询接口，可以以 ID 为条件 查询，以 是否已结账 为条件查询

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
			Name:             class.Name,
			PictureStorePath: class.PictureStorePath,
			//Goods: getClassGoods(class.Name),
		})
	}

	return &res, nil
}

func (MvpServer) GetAllGoods(ctx context.Context, req *pb.GetAllGoodsReq) (*pb.GetAllGoodsRes, error) {
	logger.Info("GetAllGoodClasses", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllGoodsRes

	res.Goods = getClassGoods(req.ClassName)

	return &res, nil
}

func (MvpServer) AddGoodClass(ctx context.Context, req *pb.AddGoodClassReq) (*pb.AddGoodClassRes, error) {
	logger.Info("AddGoodClass", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddGoodClassRes
	// todo: 判断类名是否为空、是否存在

	// 1. 创建商品类
	if err := dao.ElementClassDao.Create(&model.ElementClass{
		Name:             req.GoodClass.Name,
		PictureStorePath: req.GoodClass.PictureStorePath,
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

	var orderID int64
	if req.OrderID == 0 {
		// 这里表示用户首次点单
		dbOrder := &model.Order{}
		if err := dao.OrderDao.Create(dbOrder); err != nil {
			// todo: log
			return nil, err
		}
		orderID = int64(dbOrder.ID)
	} else {
		// 表示之前有点单了，在点单的基础上再点商品
		orderID = req.OrderID
	}

	for _, good := range req.Goods {
		dbGood := &model.Good{
			Name:              good.MainElement.Name,
			DeskID:            req.DeskID,
			Expense:           good.ExpenseInfo.Expense,
			CheckOutTimestamp: good.ExpenseInfo.CheckOutTimestamp,
			NonFavorExpense:   good.ExpenseInfo.NonFavorExpense,
			OrderID:           orderID,
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

	dbOrder := &model.Order{}
	if err := dao.OrderDao.Create(dbOrder); err != nil {
		// todo: log
		return nil, err
	}

	dbDesk := &model.Desk{
		StartTimestamp:    time.Now().Unix(),
		EndTimestamp:      0,
		SpaceName:         req.Desk.Space.Name,
		SpaceNum:          req.Desk.Space.Num,
		Expense:           req.Desk.ExpenseInfo.Expense,
		CheckOutTimestamp: req.Desk.ExpenseInfo.CheckOutTimestamp,
		NonFavorExpense:   req.Desk.ExpenseInfo.NonFavorExpense,
		OrderID:           int64(dbOrder.ID),
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

func (MvpServer) GetOrder(ctx context.Context, req *pb.GetOrderReq) (*pb.GetOrderRes, error) {
	logger.Info("GetOrder", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.GetOrderRes

	res.Order = getPbOrder(req.OrderID)
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

func (s MvpServer) ChangeDesk(ctx context.Context, req *pb.ChangeDeskReq) (*pb.ChangeDeskRes, error) {
	logger.Info("ChangeDesk", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.ChangeDeskRes

	if err := dao.DeskDao.Update(map[string]interface{}{
		"id":         req.SrcDeskID,
		"space_name": req.DstSpaceName,
		"space_num":  req.DstSpaceNum,
	}); err != nil {
		// todo: log
		return nil, err
	}

	return &res, nil
}

func (s MvpServer) CancelGood(ctx context.Context, req *pb.CancelGoodReq) (*pb.CancelGoodRes, error) {
	logger.Info("CancelGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.CancelGoodRes

	if err := dao.GoodDao.BatchDelete(req.GoodIDs); err != nil {
		// todo: log
		return nil, err
	}
	if err := dao.MainElementAttachElementRecordDao.BatchDelete(req.GoodIDs); err != nil {
		// todo: log
		return nil, err
	}
	if err := dao.MainElementSizeRecordDao.BatchDelete(req.GoodIDs); err != nil {
		// todo: log
		return nil, err
	}
	if err := dao.ChargeableObjectDao.BatchDeleteFavorRecord(model.ChargeableObjectNameOfGood, req.GoodIDs); err != nil {
		// todo: log
		return nil, err
	}

	return &res, nil
}

func (s MvpServer) AddFavorForGood(ctx context.Context, req *pb.AddFavorForGoodReq) (*pb.AddFavorForGoodRes, error) {
	logger.Info("AddFavorForGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddFavorForGoodRes
	if err := dao.ChargeableObjectDao.CreateFavorRecord(model.ChargeableObjectNameOfGood, req.GoodID, req.Favors); err != nil {
		// todo: log
		return nil, err
	}

	return &res, nil
}

func (s MvpServer) DeleteFavorForGood(ctx context.Context, req *pb.DeleteFavorForGoodReq) (*pb.DeleteFavorForGoodRes, error) {
	logger.Info("DeleteFavorForGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.DeleteFavorForGoodRes
	if err := dao.ChargeableObjectDao.DeleteFavorRecord(model.ChargeableObjectNameOfGood, req.GoodID, req.Favor); err != nil {
		// todo: log
		return nil, err
	}
	return &res, nil
}

func (s MvpServer) GetAllDeskClasses(ctx context.Context, req *pb.GetAllDeskClassesReq) (*pb.GetAllDeskClassesRes, error) {
	logger.Info("GetAllDeskClasses", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllDeskClassesRes
	spaces, err := dao.SpaceDao.GetAll()
	if err != nil {
		// todo: log
		return nil, err
	}

	nameToSpaces := make(map[string][]*model.Space)
	for _, space := range spaces {
		nameToSpaces[space.Name] = append(nameToSpaces[space.Name], space)
	}

	deskClasses := make([]*pb.DeskClass, 0)
	for name, spaces := range nameToSpaces {
		desks := make([]*pb.Desk, 0)
		for _, space := range spaces {
			desk, err := dao.DeskDao.GetNonCheckOutDesk(space.Name, space.Num)
			if err != nil {
				// todo:log
				return nil, err
			}
			if desk == nil {
				desks = append(desks, &pb.Desk{
					Id:             0,
					Space:          space.ToPb(),
					StartTimestamp: 0,
					EndTimestamp:   0,
					Favors:         nil,
					ExpenseInfo:    nil,
				})
			} else {
				desks = append(desks, getPbDesk(desk)) // todo: 这里会和 getPbDesk 有部分冗余
			}
		}
		deskClasses = append(deskClasses, &pb.DeskClass{
			Name:  name,
			Desks: desks,
		})
	}

	res.DeskClasses = deskClasses

	return &res, nil
}
