package controller

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

type MvpServer struct {
	pb.UnimplementedMvpServer
}

func (MvpServer) AddGood(ctx context.Context, req *pb.AddGoodReq) (*pb.AddGoodRes, error) {
	logs.Info("AddGood", ctx, req)

	var res pb.AddGoodRes

	// 1. 判断商品是否存在
	good, err := dao.GoodDao.GetByName(req.Good.Name)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodDao.GetByName")
	}
	if good != nil {
		logs.Error(err)
		return nil, ers.New(0, "商品(%s)已存在", good.Name)
	}

	// 2. 获得商品类ID
	goodClass, err := dao.GoodClassDao.GetByName(req.GoodClassName)
	if err != nil {
		logger.Error("Fail to finish GoodClassDao.Get",
			zap.Any("goodClassName", req.GoodClassName),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 3. 创建商品
	if err := dao.GoodDao.Create(
		req.Good.Name,
		float64(req.Good.Price),
		req.Good.PictureStorePath,
		int(goodClass.ID),
	); err != nil {
		logs.Error(err)
		return nil, err
	}

	// 4. 获得商品
	good, err = dao.GoodDao.GetByName(req.Good.Name)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodDao.GetByName")
	}

	// 5. 为商品添加附属选项
	for _, optionClass := range req.Good.OptionClasses {
		// 5.1 获得附属选项类
		tbOptionClass, err := dao.OptionClassDao.GetByName(optionClass.Name)
		if err != nil {
			logger.Error("Fail to get option class",
				zap.Any("optionClassName", optionClass.Name),
				zap.Any("req", req),
				zap.Error(err))
			return nil, ers.MysqlError
		}
		if tbOptionClass == nil {
			logger.Error("Option class not exist",
				zap.Any("optionClassName", optionClass.Name),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}

		// 5.2 执行添加操作
		if err := dao.GoodOptionClassRecordDao.Create(int(good.ID), int(tbOptionClass.ID)); err != nil {
			logger.Error("Fail to finish GoodOptionClassRecordDao.Create",
				zap.Any("goodID", good.ID),
				zap.Any("optionClassID", tbOptionClass.ID),
				zap.Error(err))
			return nil, err
		}
	}

	// 6. 为商品添加附属商品类
	for _, attachGoodClass := range req.Good.AttachGoodClasses {
		// 6.1 获得附属附属商品类
		tbAttachGoodClass, err := dao.GoodClassDao.GetByName(attachGoodClass.Name)
		if err != nil {
			logger.Error("Fail to finish GoodClassDao.Get",
				zap.Any("attachGoodClassName", attachGoodClass.Name),
				zap.Any("req", req),
				zap.Error(err))
			return nil, ers.MysqlError
		}
		if tbAttachGoodClass == nil {
			logger.Error("Attach good class not exist",
				zap.Any("attachGoodClassName", attachGoodClass.Name),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}

		// 6.2 执行添加操作
		if err := dao.GoodAttachClassRecordDao.Create(int(good.ID), int(tbAttachGoodClass.ID)); err != nil {
			logger.Error("Fail to finish GoodOptionClassRecordDao.Create",
				zap.Any("goodID", good.ID),
				zap.Any("attachGoodClassName", attachGoodClass.Name),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
	}
	return &res, nil
}

//func (MvpServer) SellGood(ctx context.Context, req *pb.SellGoodReq) (*pb.SellGoodRes, error) {
//	logs.Info("SellGood", ctx, req)
//
//	var res pb.SellGoodRes
//
//	// 1. 判断商品是否存在
//	good, err := dao.GoodDao.GetByName(req.GoodName)
//	if err != nil {
//		logs.Error(err)
//		return nil, errors.New("Fail to finish GoodDao.GetByName")
//	}
//	if good == nil {
//		return nil, errors.New("Good not existed")
//	}
//
//	// 2. 创建销售记录
//	if err := unuse.SellRecordDao.Create(int(good.ID), float64(req.SellPrice)); err != nil {
//		logs.Error(err)
//		return nil, err
//	}
//	return &res, nil
//}
//
//func (MvpServer) AddBilliardDesk(ctx context.Context, req *pb.AddBilliardDeskReq) (*pb.AddBilliardDeskRes, error) {
//	logs.Info("AddBilliardDesk", ctx, req)
//
//	var res pb.AddBilliardDeskRes
//
//	// 1. 判断桌名是否重复
//	desk, err := unuse.BilliardDeskDao.GetByName(req.BilliardDeskName)
//	if err != nil {
//		logs.Error(err)
//		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
//	}
//	if desk != nil {
//		return nil, errors.New("Desk name dumplicate")
//	}
//
//	// 2. 创建台球桌
//	if err := unuse.BilliardDeskDao.Create(req.BilliardDeskName); err != nil {
//		logs.Error(err)
//		return nil, err
//	}
//	return &res, nil
//}
//
//func (MvpServer) BeginPlayBilliard(ctx context.Context, req *pb.BeginPlayBilliardReq) (*pb.BeginPlayBilliardRes, error) {
//	logs.Info("BeginPlayBilliard", ctx, req)
//
//	var res pb.BeginPlayBilliardRes
//
//	// 1. 判断桌名是否存在
//	desk, err := unuse.BilliardDeskDao.GetByName(req.BilliardDeskName)
//	if err != nil {
//		logs.Error(err)
//		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
//	}
//	if desk == nil {
//		return nil, errors.New("Desk not exist")
//	}
//
//	// 2. 判断记录是否存在
//	record, err := unuse.PlayRecordDao.Get(int(desk.ID), time.Unix(req.BeginPlayTimestamp, 0))
//	if err != nil {
//		logs.Error(err)
//		return nil, errors.New("Fail to finish PlayRecordDao.Get")
//	}
//	if record != nil {
//		return nil, errors.New("Record has exist")
//	}
//
//	// 3. 创建玩台球的记录
//	if err := unuse.PlayRecordDao.Create(
//		int(desk.ID),
//		time.Unix(req.BeginPlayTimestamp, 0),
//	); err != nil {
//		logs.Error(err)
//		return nil, err
//	}
//	return &res, nil
//}
//
//func (MvpServer) StopPlayBilliard(ctx context.Context, req *pb.StopPlayBilliardReq) (*pb.StopPlayBilliardRes, error) {
//	logs.Info("StopPlayBilliard", ctx, req)
//
//	var res pb.StopPlayBilliardRes
//
//	// 1. 判断桌名是否存在
//	desk, err := unuse.BilliardDeskDao.GetByName(req.BilliardDeskName)
//	if err != nil {
//		logs.Error(err)
//		return nil, errors.New("Fail to finish BilliardDeskDao.GetByName")
//	}
//	if desk == nil {
//		return nil, errors.New("Desk not exist")
//	}
//
//	// 2. 更新玩台球的记录
//	if err := unuse.PlayRecordDao.UpdateStopPlayTime(
//		int(desk.ID),
//		time.Unix(req.BeginPlayTimestamp, 0),
//		time.Unix(req.StopPlayTimestamp, 0),
//	); err != nil {
//		logs.Error(err)
//		return nil, err
//	}
//	return &res, nil
//}

func (MvpServer) Order(ctx context.Context, req *pb.OrderReq) (*pb.OrderRes, error) {
	logs.Info("Order", ctx, req)

	var res pb.OrderRes
	orderID := getSetOrderID()
	tbThings := make([]*model.Thing, 0)
	// 1. 处理原商品
	for _, pbGood := range req.Order.Goods {
		// 1.1 判断原商品是否存在
		tbGood, err := dao.GoodDao.GetByName(pbGood.Name)
		if err != nil {
			logs.Error(err)
			return nil, errors.New("Fail to finish GoodDao.GetByName")
		}
		if tbGood == nil {
			return nil, errors.New("Good not existed")
		}

		// 1.2 形成物品号
		thingID := getSetThingID()

		// 1.4 处理附属选项
		for _, class := range pbGood.OptionClasses {
			// 1.4.1 判断附属选项是否存在
			tbOption, err := dao.OptionDao.GetByName(class.Options[class.SelectOptionIndex-1].Name) // todo: 要简化
			if err != nil {
				logs.Error(err)
				return nil, errors.New("Fail to finish OptionDao.GetByName")
			}
			if tbOption == nil {
				logs.Error("Option not existed")
				return nil, errors.New("Option not existed")
			}
			// 1.4.2 插入附属选项记录
			if err := dao.ThingOptionRecordDao.Create(&model.ThingOptionRecord{
				ThingID:  thingID,
				OptionID: int(tbOption.ID),
			}); err != nil {
				logs.Error(err)
				return nil, errors.New("Fail to finish ThingOptionRecordDao.Create")
			}
		}

		// 1.5 处理附属商品
		var tbAttachGoods []*model.Good
		for _, attachGoodClass := range pbGood.AttachGoodClasses {
			for _, selectGoodName := range attachGoodClass.SelectGoodNames {
				// 1.5.1 判断附属商品是否存在
				tbAttachGood, err := dao.GoodDao.GetByName(selectGoodName)
				if err != nil {
					logs.Error(err)
					return nil, errors.New("Fail to finish GoodDao.GetByName")
				}
				if tbAttachGood == nil {
					logs.Error("Good not existed")
					return nil, errors.New("Good not existed")
				}
				// 1.5.2 插入附属商品记录
				if err := dao.ThingAttachGoodRecordDao.Create(&model.ThingAttachGoodRecord{
					ThingID:      thingID,
					AttachGoodID: int(tbAttachGood.ID),
				}); err != nil {
					logs.Error(err)
					return nil, errors.New("Fail to finish AttachRecordDao.Create")
				}

				tbAttachGoods = append(tbAttachGoods, tbAttachGood)
			}
		}

		// 1.3 插入订单记录、插入物品记录
		if err := dao.OrderThingRecordDao.Create(&model.OrderThingRecord{
			OrderID: orderID,
			ThingID: thingID,
		}); err != nil {
			return nil, errors.New("Fail to finish OrderRecordDao.Create")
		}

		tbThings = append(tbThings, &model.Thing{
			Model: gorm.Model{
				ID: uint(thingID),
			},
			GoodID: int(tbGood.ID),
			Price:  getThingPrice(tbGood, tbAttachGoods),
		})
	}

	for _, tbThing := range tbThings {
		if err := dao.ThingDao.Create(tbThing); err != nil {
			logger.Error("Fail to finish ThingDao.Create",
				zap.Any("tbThing", tbThing),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
	}

	tbOrder := &model.Order{
		Model: gorm.Model{
			ID: uint(orderID),
		},
		HasCheckedOut: model.NotCheckedOut,
		Price:         getOrderPrice(tbThings),
	}
	if err := dao.OrderDao.Create(tbOrder); err != nil {
		logger.Error("Fail to finish OrderDao.Create",
			zap.Any("tbOrder", tbOrder),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	res.OrderID = int64(orderID)
	return &res, nil
}

func (MvpServer) AddGoodType(ctx context.Context, req *pb.AddGoodTypeReq) (*pb.AddGoodTypeRes, error) {
	logs.Info("AddGoodType", ctx, req)
	var res pb.AddGoodTypeRes

	// 1. 判断商品类型名是否存在
	goodType, err := dao.GoodTypeDao.GetByName(req.TypeName)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodTypeDao.GetByName")
	}
	if goodType != nil {
		logs.Error("Good Type Has Exist")
		return nil, errors.New("Good Type Has Exist")
	}

	// 2. 创建商品类型
	if err := dao.GoodTypeDao.Create(req.TypeName); err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish GoodTypeDao.Create")
	}

	return &res, nil
}

//func (MvpServer) Checkout(ctx context.Context, req *pb.CheckoutReq) (*pb.CheckoutRes, error) {
//	logs.Info("Checkout", ctx, req)
//	var res pb.CheckoutRes
//
//	// 1. 判断订单是否存在
//	orderRecords, err := dao.OrderRecordDao.GetByOrderID(int(req.OrderID))
//	if err != nil {
//		logs.Error(err)
//		return nil, errors.New("Fail to finish OrderRecordDao.GetByOrderID")
//	}
//	if orderRecords == nil {
//		logs.Error("Order Record not Exist")
//		return nil, errors.New("Order Record not Exist")
//	}
//
//	// 2. 获取未结账货物记录
//	notCheckoutOrderRecords, err := dao.OrderRecordDao.GetNotCheckoutGoods(int(req.OrderID))
//	if err != nil {
//		logs.Error(err)
//		return nil, errors.New("Fail to finish OrderRecordDao.GetNotCheckoutGoods")
//	}
//
//	// 3. 计算价格 (先不考虑打折、满减等情况)
//	var wholePrice float64
//	for _, notCheckoutOrderRecord := range notCheckoutOrderRecords {
//		// 3.1 判断商品是否存在
//		good, err := dao.GoodDao.Get(notCheckoutOrderRecord.GoodID)
//		if err != nil {
//			logs.Error(err)
//			return nil, errors.New("Fail to finish GoodDao.GetByName")
//		}
//		if good == nil {
//			return nil, errors.New("Good not exist")
//		}
//
//		// 3.2 结账该商品
//		wholePrice += good.Price
//	}
//
//	// 4. 返回结账金额
//	res.Price = float32(wholePrice)
//	return &res, nil
//}

func (MvpServer) AddOptionClass(ctx context.Context, req *pb.AddOptionClassReq) (*pb.AddOptionClassRes, error) {
	logs.Info("AddOptionClass", ctx, req)
	var res pb.AddOptionClassRes

	// 0. 判断请求数据是否合法
	if req.OptionClass == nil {
		return nil, ers.New(0, "选项类为空")
	}
	if req.OptionClass.Name == "" {
		return nil, ers.New(ers.CodeOfBlankOptionClassName, "选项类名为空")
	}
	if len(req.OptionClass.Options) == 0 {
		return nil, ers.New(ers.CodeOfEmptyOption, "选项类中没有选项")
	}
	for index, option := range req.OptionClass.Options {
		if option.Name == "" {
			return nil, ers.New(ers.CodeOfBlankOptionName, "第 %v 个选项名不能为空", index+1)
		}
	}

	// 1. 判断选项类是否存在
	optionClass, err := dao.OptionClassDao.GetByName(req.OptionClass.Name)
	if err != nil {
		logger.Error("Fail to get option class",
			zap.Any("optionClassName", req.OptionClass.Name),
			zap.Any("req", req),
			zap.Error(err))
		return nil, ers.MysqlError
	}

	// 2. 如果选项类不存在，则创建
	if optionClass == nil {
		if err := dao.OptionClassDao.Create(req.OptionClass.Name); err != nil {
			logger.Error("Fail to create option class",
				zap.Any("optionClassName", req.OptionClass.Name),
				zap.Any("req", req),
				zap.Error(err))
			return nil, ers.MysqlError
		}
	}
	optionClass, err = dao.OptionClassDao.GetByName(req.OptionClass.Name)
	if err != nil {
		logger.Error("Fail to get option class",
			zap.Any("optionClassName", req.OptionClass.Name),
			zap.Any("req", req),
			zap.Error(err))
		return nil, ers.MysqlError
	}

	// 4. 创建选项
	for _, option := range req.OptionClass.Options {
		if err := dao.OptionDao.Create(int(optionClass.ID), option.Name); err != nil {
			logger.Error("Fail to create option",
				zap.Any("optionClassID", optionClass.ID),
				zap.Any("optionName", option.Name),
				zap.Any("req", req),
				zap.Error(err))
			return nil, ers.MysqlError
		}
	}
	return &res, nil
}

func (MvpServer) GetAllOptionClasses(ctx context.Context, req *pb.GetAllOptionClassesReq) (*pb.GetAllOptionClassesRes, error) {
	logs.Info("GetAllOptionClasses", ctx, req)
	var res pb.GetAllOptionClassesRes

	// 1. 获取所有选项类名
	optionClasses, err := dao.OptionClassDao.GetAll()
	if err != nil {
		logger.Error("Fail to get all option classes",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 转换
	for _, optionClass := range optionClasses {
		pbOptionClass, err := formPbOptionClass(optionClass)
		if err != nil {
			logger.Error("Fail to form pb option class",
				zap.Any("tableOptionClass", optionClass),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
		res.OptionClasses = append(res.OptionClasses, pbOptionClass)
	}
	return &res, nil
}

func (MvpServer) DelOption(ctx context.Context, req *pb.DelOptionReq) (*pb.DelOptionRes, error) {
	logs.Info("DelOption", ctx, req)
	var res pb.DelOptionRes

	// 1. 获得选项类, 判断是否需要删除
	optionClass, err := dao.OptionClassDao.GetByName(req.ClassName)
	if err != nil {
		logger.Error("Fail to get all option classes",
			zap.Any("className", req.ClassName),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	if optionClass == nil {
		return &res, nil
	}

	// 2. 删除选项类
	if err := dao.OptionDao.Del(int(optionClass.ID), req.OptionName); err != nil {
		logger.Error("Fail to del option",
			zap.Any("optionClassID", optionClass.ID),
			zap.Any("optionName", req.OptionName),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) GetAllGoods(ctx context.Context, req *pb.GetAllGoodsReq) (*pb.GetAllGoodsRes, error) {
	logs.Info("GetAllGoods", ctx, req)
	var res pb.GetAllGoodsRes

	// 1. 获得所有商品
	goods, err := dao.GoodDao.GetAll()
	if err != nil {
		logger.Error("Fail to get all goods",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 转化
	for _, good := range goods {
		pbGood, err := formPbGood(good)
		if err != nil {
			logger.Error("Fail to form pb.good",
				zap.Any("tableGood", good),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
		res.Goods = append(res.Goods, pbGood)
	}

	return &res, nil
}

func (MvpServer) DelOptionClass(ctx context.Context, req *pb.DelOptionClassReq) (*pb.DelOptionClassRes, error) {
	logs.Info("DelOptionClass", ctx, req)
	var res pb.DelOptionClassRes

	// 1. 获取选项类名
	optionClassNames := make([]string, 0)
	for _, optionClass := range req.OptionClasses {
		optionClassNames = append(optionClassNames, optionClass.Name)
	}

	// 2. 删除
	if err := dao.OptionClassDao.DeleteByNames(optionClassNames); err != nil {
		logger.Error("Fail to finish OptionClassDao.DeleteByName",
			zap.Any("optionClassNames", optionClassNames),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) GetAllGoodClasses(ctx context.Context, req *pb.GetAllGoodClassesReq) (*pb.GetAllGoodClassesRes, error) {
	logs.Info("GetAllGoodClasses", ctx, req)
	var res pb.GetAllGoodClassesRes

	// 1. 获得所有商品类
	goodClasses, err := dao.GoodClassDao.GetAll()
	if err != nil {
		logger.Error("Fail to get all good classes",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 转化
	for _, goodClass := range goodClasses {
		pbGoodClass, err := formPbGoodClass(goodClass)
		if err != nil {
			logger.Error("Fail to form pb.goodClass",
				zap.Any("tableGoodClass", goodClass),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
		res.GoodClasses = append(res.GoodClasses, pbGoodClass)
	}

	return &res, nil
}

func (MvpServer) AddGoodClass(ctx context.Context, req *pb.AddGoodClassReq) (*pb.AddGoodClassRes, error) {
	logs.Info("AddGoodClass", ctx, req)

	var res pb.AddGoodClassRes

	// 1. 创建商品类
	if err := dao.GoodClassDao.Create(&model.GoodClass{
		Name:      req.GoodClass.Name,
		ClassType: int(req.GoodClass.ClassType),
	}); err != nil {
		logger.Error("Fail to finish GoodClassDao.Create",
			zap.String("goodClassName", req.GoodClass.Name),
			zap.Any("goodClassType", req.GoodClass.ClassType),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 获得商品类
	goodClass, err := dao.GoodClassDao.GetByName(req.GoodClass.Name)
	if err != nil {
		logger.Error("Fail to finish GoodClassDao.Get",
			zap.Any("goodClassName", req.GoodClass.Name),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	if goodClass == nil {
		logger.Warn("Fail to create good class",
			zap.Any("goodClassName", req.GoodClass.Name),
			zap.Any("req", req),
			zap.Error(err))
		return nil, nil
	}

	// 3. 更新商品
	for _, good := range req.GoodClass.Goods {
		if err := dao.GoodDao.UpdateGoodClassID(good.Name, int(goodClass.ID)); err != nil {
			logger.Error("Fail to finish GoodDao.UpdateGoodClassID",
				zap.String("goodClassName", req.GoodClass.Name),
				zap.Uint("goodClassID", goodClass.ID),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
	}
	return &res, nil
}

func (MvpServer) DelGoodClass(ctx context.Context, req *pb.DelGoodClassReq) (*pb.DelGoodClassRes, error) {
	logs.Info("DelGoodClass", ctx, req)
	var res pb.DelGoodClassRes

	// 1. 获取商品类名
	goodClassNames := make([]string, 0)
	for _, goodClass := range req.GoodClasses {
		goodClassNames = append(goodClassNames, goodClass.Name)
	}

	// 2. 删除
	if err := dao.GoodClassDao.DeleteByNames(goodClassNames); err != nil {
		logger.Error("Fail to finish GoodClassDao.DeleteByNames",
			zap.Any("optionClassNames", goodClassNames),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) GetOrderLog(ctx context.Context, req *pb.GetOrderLogReq) (*pb.GetOrderLogRes, error) {
	logger.Info("GetOrderLog", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetOrderLogRes

	orderLog, err := getOrderLog(int(req.OrderID))
	if err != nil {
		logger.Error("Fail to finish getOrderLog",
			zap.Int64("orderID", req.OrderID),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	res.OrderLog = orderLog
	return &res, nil
}

