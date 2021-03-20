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
	if errResult := createElement(req.Good.MainElement, req.ClassName); errResult != nil {
		return nil, errResult
	}
	if errResult := writeGoodSizeToDB(req.Good); errResult != nil {
		return nil, errResult
	}
	return &res, nil
}

func (MvpServer) DeleteElementSizeInfo(ctx context.Context, req *pb.DeleteElementSizeInfoReq) (*pb.DeleteElementSizeInfoRes, error) {
	logger.Info("DeleteElementSizeInfo", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.DeleteElementSizeInfoRes
	if errResult := dao.ElementDao.Del(req.ElementName, req.SizeInfoSize); errResult != nil {
		return nil, errResult
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
	if _, errResult := dao.ElementSizeRecordDao.Create(&model.ElementSizeRecord{
		GoodID:      0,
		ElementName: req.Element.Name,
		SelectSize:  req.Element.SizeInfos[req.Element.SelectedIndex].Size,
	}); errResult != nil {
		return nil, errResult
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
	classes, err := dao.GoodClassDao.GetAllClasses()
	if err != nil {
		logger.Error("Fail to finish GoodClassDao.GetMainElementClass",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 形成 GoodClasses
	for _, class := range classes {
		res.GoodClasses = append(res.GoodClasses, class.ToPb())
	}

	return &res, nil
}

func (MvpServer) GetAllGoods(ctx context.Context, req *pb.GetAllGoodsReq) (*pb.GetAllGoodsRes, error) {
	logger.Info("GetAllGoods", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllGoodsRes

	res.Goods = getClassGoods(req.ClassName)

	return &res, nil
}

func (MvpServer) GetAllGoodOptions(ctx context.Context, req *pb.GetAllGoodOptionsReq) (*pb.GetAllGoodOptionsRes, error) {
	logger.Info("GetAllElements", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllGoodOptionsRes

	dbElements, err := dao.ElementDao.GetAllAttachElements(req.ClassName)
	if err != nil {
		logger.Error("Fail to finish ElementDao.GetAll",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	nameToElements := make(map[string][]*model.Element)
	for _, dbElement := range dbElements {
		nameToElements[dbElement.Name] = append(nameToElements[dbElement.Name], dbElement)
	}

	pbElements := make([]*pb.Element, 0)
	for name := range nameToElements {
		pbElements = append(pbElements, getElement(0, name, req.ClassName))
	}

	res.Elements = pbElements

	return &res, nil
}

func (MvpServer) AddGoodClass(ctx context.Context, req *pb.AddGoodClassReq) (*pb.AddGoodClassRes, error) {
	logger.Info("AddGoodClass", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddGoodClassRes
	// todo: 判断类名是否为空、是否存在

	// 1. 创建商品类
	dbGoodClass := &model.GoodClass{
		ID:               uint(req.GoodClass.Id),
		Name:             req.GoodClass.Name,
		PictureStorePath: req.GoodClass.PictureStorePath,
	}
	id, errResult := dao.GoodClassDao.Create(dbGoodClass)
	if errResult != nil {
		return nil, errResult
	}

	dbGoodClass.ID = uint(id)
	res.GoodClass = dbGoodClass.ToPb()

	return &res, nil
}

func (MvpServer) AddDeskClass(ctx context.Context, req *pb.AddDeskClassReq) (*pb.AddDeskClassRes, error) {
	logger.Info("AddGoodClass", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddDeskClassRes
	// todo: 判断类名是否为空、是否存在

	// 1. 创建商品类
	if err := dao.DeskClassDao.Create(&model.DeskClass{
		Name:             req.DeskClass.Name,
		PictureStorePath: req.DeskClass.PictureStorePath,
	}); err != nil {
		logger.Error("Fail to finish DeskClassDao.Create",
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

	// 正常情况能来到这里的，orderID 不为0

	for _, good := range req.Goods {
		if good.ExpenseInfo == nil {
			good.ExpenseInfo = &pb.ExpenseInfo{}
		}
		dbGood := &model.Good{
			Name:              good.MainElement.Name,
			Expense:           good.ExpenseInfo.Expense,
			CheckOutTimestamp: good.ExpenseInfo.CheckOutTimestamp,
			NonFavorExpense:   good.ExpenseInfo.NonFavorExpense,
			OrderID:           req.OrderID,
		}
		if err := dao.GoodDao.Create(dbGood); err != nil {
			// todo: log
			return nil, err
		}
		if err := dao.ChargeableObjectDao.CreateFavorRecord(dbGood.GetName(), int64(dbGood.ID), good.Favors); err != nil {
			// todo: log
			return nil, err
		}
		good.Id = int64(dbGood.ID)
		if err := writeGoodSizeToDB(good); err != nil {
			logger.Error("Fail to finish createGood",
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
	}
	return &res, nil
}

func (MvpServer) AddDesk(ctx context.Context, req *pb.AddDeskReq) (*pb.AddDeskRes, error) {
	logger.Info("AddDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	// todo: 这里可以点单记录
	var res pb.AddDeskRes

	req.Desk.Space.ClassName = req.ClassName

	if err := dao.SpaceDao.Create(getDbSpace(req.Desk.Space)); err != nil {
		logger.Error("Fail to finish SpaceDao.Create",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
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
	if req.Desk.ExpenseInfo == nil {
		req.Desk.ExpenseInfo = &pb.ExpenseInfo{}
	}
	if req.Desk.Space == nil {
		req.Desk.Space = &pb.Space{}
	}
	dbDesk := &model.Desk{
		StartTimestamp:    time.Now().Unix(),
		EndTimestamp:      0,
		SpaceName:         req.Desk.Space.Name,
		SpaceClassName:    req.Desk.Space.ClassName,
		Expense:           req.Desk.ExpenseInfo.Expense,
		CheckOutTimestamp: req.Desk.ExpenseInfo.CheckOutTimestamp,
		NonFavorExpense:   req.Desk.ExpenseInfo.NonFavorExpense,
		OrderID:           int64(dbOrder.ID),
	}
	if err := dao.DeskDao.Create(dbDesk); err != nil {
		// todo: log
		return nil, err
	}
	if err := dao.ChargeableObjectDao.CreateFavorRecord(dbDesk.GetName(), int64(dbDesk.ID), req.Desk.Favors); err != nil {
		// todo: log
		return nil, err
	}
	res.DeskID = int64(dbDesk.ID)
	res.OrderID = int64(dbOrder.ID)
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
	if err := dao.ElementSizeRecordDao.BatchDelete(req.GoodIDs); err != nil {
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

func (s MvpServer) GetAllDesks(ctx context.Context, req *pb.GetAllDesksReq) (*pb.GetAllDesksRes, error) {
	logger.Info("GetAllDesks", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllDesksRes

	spaces, err := dao.SpaceDao.GetByClassName(req.ClassName)
	if err != nil {
		// todo: log
		return nil, err
	}
	desks := make([]*pb.Desk, 0)
	for _, space := range spaces {
		desk, err := dao.DeskDao.GetNonCheckOutDesk(space.Name, space.ClassName)
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
				OrderID:        0,
			})
		} else {
			desks = append(desks, getPbDesk(desk))
		}
	}

	res.Desks = desks
	return &res, nil
}

func (s MvpServer) GetAllDeskClasses(ctx context.Context, req *pb.GetAllDeskClassesReq) (*pb.GetAllDeskClassesRes, error) {
	logger.Info("GetAllDeskClasses", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllDeskClassesRes
	//spaces, err := dao.SpaceDao.GetAll()
	//if err != nil {
	//	// todo: log
	//	return nil, err
	//}
	//
	//nameToSpaces := make(map[string][]*model.Space)
	//for _, space := range spaces {
	//	nameToSpaces[space.Name] = append(nameToSpaces[space.Name], space)
	//}

	//deskClasses := make([]*pb.DeskClass, 0)
	//for name, spaces := range nameToSpaces {
	//	desks := make([]*pb.Desk, 0)
	//	for _, space := range spaces {
	//		desk, err := dao.DeskDao.GetNonCheckOutDesk(space.Name, space.Num)
	//		if err != nil {
	//			// todo:log
	//			return nil, err
	//		}
	//		if desk == nil {
	//			desks = append(desks, &pb.Desk{
	//				Id:             0,
	//				Space:          space.ToPb(),
	//				StartTimestamp: 0,
	//				EndTimestamp:   0,
	//				Favors:         nil,
	//				ExpenseInfo:    nil,
	//			})
	//		} else {
	//			desks = append(desks, getPbDesk(desk)) // todo: 这里会和 getPbDesk 有部分冗余
	//		}
	//	}
	//	deskClasses = append(deskClasses, &pb.DeskClass{
	//		Name: name,
	//		//Desks: desks,
	//	})
	//}

	deskClasses, err := dao.DeskClassDao.GetAllClasses()
	if err != nil {
		// todo: log
		return nil, err
	}

	var pbDeskClass []*pb.DeskClass

	for _, deskClass := range deskClasses {
		pbDeskClass = append(pbDeskClass, deskClass.ToPb())
	}

	res.DeskClasses = pbDeskClass

	return &res, nil
}
