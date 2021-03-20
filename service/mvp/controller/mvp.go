package controller

import (
	"context"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

// todo: GetOrder 可以设置为一个查询接口，可以以 ID 为条件 查询，以 是否已结账 为条件查询

type MvpServer struct {
	pb.UnimplementedMvpServer
}

func (MvpServer) AddGood(ctx context.Context, req *pb.AddGoodReq) (*pb.AddGoodRes, error) {
	logger.Info("AddGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddGoodRes
	good, goodClassName, errResult := CheckAddGoodParameter(req)
	if errResult != nil {
		return nil, errResult
	}
	if errResult := writePbElementToDbAndUpdateID(good.MainElement, getDbElementClassByName(goodClassName).ID); errResult != nil {
		return nil, errResult
	}
	if errResult := writePbGoodSizeInfoToDB(good); errResult != nil {
		return nil, errResult
	}
	return &res, nil
}

func (MvpServer) AddElement(ctx context.Context, req *pb.AddElementReq) (*pb.AddElementRes, error) {
	logger.Info("AddElement", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddElementRes

	if errResult := writePbElementToDbAndUpdateID(req.Element, getDbElementClassByName(req.ClassName).ID); errResult != nil {
		return nil, errResult
	}
	if _, errResult := dao.ElementSelectSizeRecordDao.Create(toDbElementSelectSizeRecord(0, req.Element.Id, model.GetPbElementSelectSizeInfo(req.Element).Id));
		errResult != nil {
		return nil, errResult
	}

	return &res, nil
}

func (MvpServer) DeleteElementSizeInfo(ctx context.Context, req *pb.DeleteElementSizeInfoReq) (*pb.DeleteElementSizeInfoRes, error) {
	logger.Info("DeleteElementSizeInfo", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.DeleteElementSizeInfoRes

	return &res, nil
}

//func (MvpServer) AddSpace(ctx context.Context, req *pb.AddSpaceReq) (*pb.AddSpaceRes, error) {
//	logger.Info("AddDesk", zap.Any("ctx", ctx), zap.Any("req", req))
//	var res pb.AddSpaceRes
//
//	// 1. 创建
//	if _, err := dao.SpaceDao.Create(toDbSpace(req.Space)); err != nil {
//		logger.Error("Fail to finish SpaceDao.Create",
//			zap.Any("req", req),
//			zap.Error(err))
//		return nil, err
//	}
//	return &res, nil
//}

func (MvpServer) GetAllGoodClasses(ctx context.Context, req *pb.GetAllGoodClassesReq) (*pb.GetAllGoodClassesRes, error) {
	logger.Info("GetAllGoodClasses", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllGoodClassesRes

	// 1. 获得主元素的所有类
	classes, err := dao.ElementClassDao.GetClasses(pb.ElementType_Main)
	if err != nil {
		logger.Error("Fail to finish ElementClassDao.GetMainElementClass",
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

	res.Goods = getClassGoods(getDbElementClassByName(req.ClassName).ID)

	return &res, nil
}

func (MvpServer) GetAllGoodOptions(ctx context.Context, req *pb.GetAllGoodOptionsReq) (*pb.GetAllGoodOptionsRes, error) {
	logger.Info("GetAllGoodOptions", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllGoodOptionsRes

	var classId int64
	if req.ClassName != "" {
		elementClass := getDbElementClassByName(req.ClassName)
		if elementClass == nil {
			return nil, ers.New("类名不存在。")
		}
		classId = elementClass.ID
	}

	// 1. 从数据库中获取所有附属元素
	dbElements, err := dao.ElementDao.GetAllAttachElements(classId)
	if err != nil {
		logger.Error("Fail to finish ElementDao.GetAll",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 获取 pbElement
	pbElements := make([]*pb.Element, 0)
	for _, dbElement := range dbElements {
		pbElements = append(pbElements, getPbElement(0, dbElement.ID))
	}

	// 4. 写入
	res.Elements = pbElements

	// 5. 返回
	return &res, nil
}

func (MvpServer) AddGoodClass(ctx context.Context, req *pb.AddGoodClassReq) (*pb.AddGoodClassRes, error) {
	logger.Info("AddGoodClass", zap.Any("ctx", ctx), zap.Any("req", req))

	// 1. 参数校验、获取
	goodClass, errResult := CheckAddGoodClassParameter(req)
	if errResult != nil {
		return nil, errResult
	}

	// 2. 创建商品类、获取生成商品类ID
	dbGoodClass := &model.ElementClass{
		ID:               goodClass.Id,
		Name:             goodClass.Name,
		PictureStorePath: goodClass.PictureStorePath,
	}
	id, errResult := dao.ElementClassDao.Create(dbGoodClass)
	if errResult != nil {
		return nil, errResult
	}

	// 3. 写响应
	var res pb.AddGoodClassRes
	dbGoodClass.ID = id
	res.GoodClass = dbGoodClass.ToPb()

	// 4. 返回响应
	return &res, nil
}

func (MvpServer) AddDeskClass(ctx context.Context, req *pb.AddDeskClassReq) (*pb.AddDeskClassRes, error) {
	logger.Info("AddGoodClass", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddDeskClassRes
	// todo: 判断类名是否为空、是否存在

	// 1. 创建商品类
	if _, err := dao.SpaceClassDao.Create(&model.SpaceClass{
		ID:               req.DeskClass.Id,
		Name:             req.DeskClass.Name,
		PictureStorePath: req.DeskClass.PictureStorePath,
	}); err != nil {
		logger.Error("Fail to finish SpaceClassDao.Create",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

//func (MvpServer) OrderGood(ctx context.Context, req *pb.OrderGoodReq) (*pb.OrderGoodRes, error) {
//	logger.Info("OrderGood", zap.Any("ctx", ctx), zap.Any("req", req))
//	// todo: 这里可以点单记录
//	var res pb.OrderGoodRes
//
//	// 正常情况能来到这里的，orderID 不为0
//
//	for _, good := range req.Goods {
//		if good.ExpenseInfo == nil {
//			good.ExpenseInfo = &pb.ExpenseInfo{}
//		}
//		dbGood := &model.Good{
//			Name:            good.MainElement.Name,
//			Expense:         good.ExpenseInfo.Expense,
//			CheckOutAt:      time.Unix(good.ExpenseInfo.CheckOutAt, 0),
//			NonFavorExpense: good.ExpenseInfo.NonFavorExpense,
//			OrderID:         req.OrderID,
//		}
//		if err := dao.GoodDao.Create(dbGood); err != nil {
//			// todo: log
//			return nil, err
//		}
//		if err := dao.ChargeableObjectDao.CreateFavorRecord(dbGood.GetName(), int64(dbGood.ID), good.Favors); err != nil {
//			// todo: log
//			return nil, err
//		}
//		good.Id = int64(dbGood.ID)
//		if err := writePbGoodSizeInfoToDB(good, "todo"); err != nil {
//			logger.Error("Fail to finish createGood",
//				zap.Any("req", req),
//				zap.Error(err))
//			return nil, err
//		}
//	}
//	return &res, nil
//}

func (MvpServer) AddDesk(ctx context.Context, req *pb.AddDeskReq) (*pb.AddDeskRes, error) {
	logger.Info("AddDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	// todo: 这里可以点单记录
	var res pb.AddDeskRes
	if _, err := dao.SpaceDao.Create(toDbSpace(req.Desk.Space, getDbSpaceClassByName(req.ClassName).ID)); err != nil {
		logger.Error("Fail to finish SpaceDao.Create",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

//func (MvpServer) OrderDesk(ctx context.Context, req *pb.OrderDeskReq) (*pb.OrderDeskRes, error) {
//	logger.Info("OrderDesk", zap.Any("ctx", ctx), zap.Any("req", req))
//	// todo: 这里可以点单记录
//	var res pb.OrderDeskRes
//
//	dbOrder := &model.Order{}
//	if err := dao.OrderDao.Create(dbOrder); err != nil {
//		// todo: log
//		return nil, err
//	}
//	if req.Desk.ExpenseInfo == nil {
//		req.Desk.ExpenseInfo = &pb.ExpenseInfo{}
//	}
//	if req.Desk.Space == nil {
//		req.Desk.Space = &pb.Space{}
//	}
//	dbDesk := &model.Desk{
//		StartAt:  time.Now().Unix(),
//		EndAt:    0,
//		SpaceName:       req.Desk.Space.Name,
//		SpaceClassName:  req.Desk.Space.ClassName,
//		Expense:         req.Desk.ExpenseInfo.Expense,
//		CheckOutAt:      time.Unix(req.Desk.ExpenseInfo.CheckOutAt, 0),
//		NonFavorExpense: req.Desk.ExpenseInfo.NonFavorExpense,
//		OrderID:         int64(dbOrder.ID),
//	}
//	if err := dao.DeskDao.Create(dbDesk); err != nil {
//		// todo: log
//		return nil, err
//	}
//	if err := dao.ChargeableObjectDao.CreateFavorRecord(dbDesk.GetName(), int64(dbDesk.ID), req.Desk.Favors); err != nil {
//		// todo: log
//		return nil, err
//	}
//	res.DeskID = int64(dbDesk.ID)
//	res.OrderID = int64(dbOrder.ID)
//	return &res, nil
//}

//func (MvpServer) GetOrder(ctx context.Context, req *pb.GetOrderReq) (*pb.GetOrderRes, error) {
//	logger.Info("GetOrder", zap.Any("ctx", ctx), zap.Any("req", req))
//	var res pb.GetOrderRes
//
//	res.Order = getPbOrder(req.OrderID)
//	return &res, nil
//}

func (MvpServer) CloseDesk(ctx context.Context, req *pb.CloseDeskReq) (*pb.CloseDeskRes, error) {
	logger.Info("CloseDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.CloseDeskRes
	if err := closeDeskIfOpening(req.DeskID, req.EndAt); err != nil {
		logger.Error("Fail to finish closeDesk",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

//func (MvpServer) CheckOut(ctx context.Context, req *pb.CheckOutReq) (*pb.CheckOutRes, error) {
//	logger.Info("CheckOut", zap.Any("ctx", ctx), zap.Any("req", req))
//	var res pb.CheckOutRes
//
//	for _, id := range req.GoodIDs {
//		good, err := dao.GoodDao.Get(id)
//		if err != nil {
//			// todo: log
//			return nil, err
//		}
//		if good == nil {
//			// todo: warning
//			continue
//		}
//		if err := checkOutIfNot(good); err != nil {
//			// todo: log
//			return nil, err
//		}
//	}
//
//	for _, id := range req.DeskIDs {
//		desk, err := dao.DeskDao.Get(id)
//		if err != nil {
//			// todo: log
//			return nil, err
//		}
//		if desk == nil {
//			// todo: warning
//			continue
//		}
//		if err := checkOutIfNot(desk); err != nil {
//			// todo: log
//			return nil, err
//		}
//	}
//
//	return &res, nil
//}

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
	if err := dao.ElementSelectSizeRecordDao.BatchDelete(req.GoodIDs); err != nil {
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

func getDbSpaceClassByID(classID int64) *model.SpaceClass {
	spaceClass, errResult := dao.SpaceClassDao.Get(classID)
	if errResult != nil {
		// todo: log
		return nil
	}
	return spaceClass
}
func getDbElementClassByID(classID int64) *model.ElementClass {
	elementClass, errResult := dao.ElementClassDao.Get(classID)
	if errResult != nil {
		// todo: log
		return nil
	}
	return elementClass
}
func getDbSpaceClassByName(className string) *model.SpaceClass {
	spaceClass, errResult := dao.SpaceClassDao.GetByName(className)
	if errResult != nil {
		// todo: log
		return nil
	}
	return spaceClass
}
func getDbElementClassByName(className string) *model.ElementClass {
	elementClass, errResult := dao.ElementClassDao.GetByName(className)
	if errResult != nil {
		// todo: log
		return nil
	}
	return elementClass
}

func (s MvpServer) GetAllDesks(ctx context.Context, req *pb.GetAllDesksReq) (*pb.GetAllDesksRes, error) {
	logger.Info("GetAllDesks", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllDesksRes

	spaceClass := getDbSpaceClassByName(req.ClassName)
	if spaceClass == nil {
		return &res, nil
	}

	spaces, err := dao.SpaceDao.GetByClassName(spaceClass.ID)
	if err != nil {
		// todo: log
		return nil, err
	}
	desks := make([]*pb.Desk, 0)
	for _, space := range spaces {
		desk, err := dao.DeskDao.GetNonCheckOutDesk(space.ID)
		if err != nil {
			// todo:log
			return nil, err
		}

		if desk == nil {
			desks = append(desks, &pb.Desk{
				Id:          0,
				Space:       space.ToPb(getDbSpaceClassByID(space.ClassID).Name),
				StartAt:     0,
				EndAt:       0,
				Favors:      nil,
				ExpenseInfo: nil,
				OrderID:     0,
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

	deskClasses, err := dao.SpaceClassDao.GetAllClasses()
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
