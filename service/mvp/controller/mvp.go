package controller

import (
	"context"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"spectrum/service/mvp/utils"
	"time"
)

type MvpServer struct {
	pb.UnimplementedMvpServer
}

// 管理员添加商品接口
func (MvpServer) AddGood(ctx context.Context, req *pb.AddGoodReq) (*pb.AddGoodRes, error) {
	logger.Info("AddGood", zap.Any("ctx", ctx), zap.Any("req", req))

	// 1. 参数校验
	good, goodClassName, errResult := checkAddGoodParameter(req)
	if errResult != nil {
		return nil, errResult
	}

	// 2. 将商品的信息记录到数据库
	if errResult := writePbElementMetaObjectToDbAndUpdateID(good.MainElement, getDbElementClassByName(goodClassName).ID); errResult != nil {
		return nil, errResult
	}
	if errResult := writePbGoodSelectedSizeInfoIndexRecordAndMainAttachElementRecordToDB(good); errResult != nil {
		return nil, errResult
	}

	// 3. 创建并返回响应
	var res pb.AddGoodRes
	res.Good = good
	return &res, nil
}

// 用户点单商品接口
func (MvpServer) OrderGood(ctx context.Context, req *pb.OrderGoodReq) (*pb.OrderGoodRes, error) {
	logger.Info("OrderGood", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.OrderGoodRes

	// 1. 参数校验
	orderID, goods, errResult := checkOrderGoodParameter(req)
	if errResult != nil {
		return nil, errResult
	}

	// 2. 点单商品集
	for _, good := range goods {
		if errResult := writePbGoodToDB(good, orderID); errResult != nil {
			return nil, errResult
		}
	}

	// 3. 返回
	return &res, nil
}

// 管理员添加元素接口，如商品附属选项、商品附属配料。
func (MvpServer) AddElement(ctx context.Context, req *pb.AddElementReq) (*pb.AddElementRes, error) {
	logger.Info("AddElement", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddElementRes

	if errResult := writePbElementMetaObjectToDbAndUpdateID(req.Element, getDbElementClassByName(req.ClassName).ID); errResult != nil {
		return nil, errResult
	}
	if _, errResult := dao.ElementSelectSizeRecordDao.Create(toDbElementSelectSizeRecord(0, req.Element.Id,
		model.GetPbElementSelectSizeInfo(req.Element).Id));
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

	// 1. 获取类下的所有商品
	goods, errResult := getClassGoods(getDbElementClassByName(req.ClassName).ID)
	if errResult != nil {
		return nil, errResult
	}

	// 2. 写入响应
	var res pb.GetAllGoodsRes
	res.Goods = goods

	// 3. 返回
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
	goodClass, errResult := checkAddGoodClassParameter(req)
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
	if dbGoodClass.ID == 0 {
		dbGoodClass.ID = id
	}
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

func (MvpServer) AddDesk(ctx context.Context, req *pb.AddDeskReq) (*pb.AddDeskRes, error) {
	logger.Info("AddDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.AddDeskRes

	id, err := dao.SpaceDao.Create(toDbSpace(req.Desk.Space, getDbSpaceClassByName(req.ClassName).ID))
	if req.Desk.Space.Id == 0 {
		req.Desk.Space.Id = id
	}
	if err != nil {
		logger.Error("Fail to finish SpaceDao.Create",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	res.Desk = req.Desk
	return &res, nil
}

func (MvpServer) OrderDesk(ctx context.Context, req *pb.OrderDeskReq) (*pb.OrderDeskRes, error) {
	logger.Info("OrderDesk", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.OrderDeskRes

	orderID, errResult := dao.OrderDao.Create(&model.Order{
		CheckOutAt: utils.NilTime,
	})
	if errResult != nil {
		return nil, errResult
	}
	if req.Desk.ExpenseInfo == nil {
		req.Desk.ExpenseInfo = &pb.ExpenseInfo{}
	}
	if req.Desk.Space == nil {
		req.Desk.Space = &pb.Space{}
	}

	dbDesk := &model.Desk{
		ID:              req.Desk.Id,
		StartAt:         time.Now(),
		EndAt:           utils.NilTime,
		SessionCount:    0,
		SpaceID:         req.Desk.Space.Id,
		Expense:         req.Desk.ExpenseInfo.Expense,
		CheckOutAt:      utils.ToTime(req.Desk.ExpenseInfo.CheckOutAt),
		NonFavorExpense: req.Desk.ExpenseInfo.NonFavorExpense,
		OrderID:         orderID,
	}
	deskID, errResult := dao.DeskDao.Create(dbDesk)
	if errResult != nil {
		return nil, errResult
	}
	if errResult := dao.FavorRecordDao.CreateFavorRecord(dbDesk.GetChargeableObjectName(), deskID, req.Desk.Favors); errResult != nil {
		return nil, errResult
	}
	res.Order = getPbOrder(orderID)
	return &res, nil
}

func (MvpServer) GetOrder(ctx context.Context, req *pb.GetOrderReq) (*pb.GetOrderRes, error) {
	logger.Info("GetOrder", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.GetOrderRes

	dbOrders, errResult := dao.OrderDao.Query(&model.QueryOrderParameter{
		CheckOutState: req.CheckOutState,
		CreatedTimeInterval: model.TimeInterval{
			Start: time.Unix(req.StartAt, 0),
			End:   time.Unix(req.EndAt, 0),
		},
		OrderID: req.OrderID,
	})
	if errResult != nil {
		return nil, errResult
	}
	for _, order := range dbOrders {
		res.Orders = append(res.Orders, getPbOrder(order.ID))
	}
	return &res, nil
}

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
	if err := dao.FavorRecordDao.BatchDeleteFavorRecord(utils.ChargeableObjectNameOfGood, req.GoodIDs); err != nil {
		// todo: log
		return nil, err
	}

	return &res, nil
}

func (s MvpServer) AddFavorForGood(ctx context.Context, req *pb.AddFavorForGoodReq) (*pb.AddFavorForGoodRes, error) {
	logger.Info("AddFavorForGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddFavorForGoodRes
	if err := dao.FavorRecordDao.CreateFavorRecord(utils.ChargeableObjectNameOfGood, req.GoodID, req.Favors); err != nil {
		// todo: log
		return nil, err
	}

	return &res, nil
}

func (s MvpServer) DeleteFavorForGood(ctx context.Context, req *pb.DeleteFavorForGoodReq) (*pb.DeleteFavorForGoodRes, error) {
	logger.Info("DeleteFavorForGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.DeleteFavorForGoodRes
	if err := dao.FavorRecordDao.DeleteFavorRecord(utils.ChargeableObjectNameOfGood, req.GoodID, req.Favor); err != nil {
		// todo: log
		return nil, err
	}
	return &res, nil
}

func (s MvpServer) GetAllDesks(ctx context.Context, req *pb.GetAllDesksReq) (*pb.GetAllDesksRes, error) {
	logger.Info("GetAllDesks", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllDesksRes

	spaceClass := getDbSpaceClassByName(req.ClassName)
	if spaceClass == nil {
		return &res, nil
	}

	spaces, errResult := dao.SpaceDao.GetByClassID(spaceClass.ID)
	if errResult != nil {
		return nil, errResult
	}
	desks := make([]*pb.Desk, 0)
	for _, space := range spaces {
		desk, errResult := dao.DeskDao.GetNonCheckOutDesk(space.ID)
		if errResult != nil {
			return nil, errResult
		}
		if desk == nil {
			desks = append(desks, &pb.Desk{
				Id:          0,
				OrderID:     0,
				Space:       space.ToPb(getDbSpaceClassByID(space.ClassID).Name),
				StartAt:     utils.NilTime.Unix(),
				EndAt:       utils.NilTime.Unix(),
				Favors:      nil,
				ExpenseInfo: nil,
			})
		} else {
			desks = append(desks, getPbDesk(desk.ID, space.ID))
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
