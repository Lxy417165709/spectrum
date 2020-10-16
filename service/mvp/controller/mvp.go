package controller

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
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
	if err := dao.GoodDao.Create(req.GoodName, float64(req.Price), req.Type); err != nil {
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

func (MvpServer) Order(ctx context.Context, req *pb.OrderReq) (*pb.OrderRes, error) {
	logs.Info("Order", ctx, req)

	var res pb.OrderRes
	orderID := getOrderID()

	// 1. 处理原商品
	for _, pbGood := range req.Goods {
		// 1.1 判断原商品是否存在
		sGood, err := dao.GoodDao.GetByName(pbGood.Name)
		if err != nil {
			logs.Error(err)
			return nil, errors.New("Fail to finish GoodDao.GetByName")
		}
		if sGood == nil {
			return nil, errors.New("Good not existed")
		}

		// 1.2 插入订单记录
		if err := dao.OrderRecordDao.Create(orderID, int(sGood.ID), model.FlagOfNotAttachGood); err != nil {
			return nil, errors.New("Fail to finish OrderRecordDao.Create")
		}

		thingID := getThingID()
		// 1.3 判断附属商品是否存在
		for _, attachGoodArray := range pbGood.AttachGoods {
			for _, pbAttachGood := range attachGoodArray.AttachGoods {
				// 1.3.1 判断附属商品是否存在
				good, err := dao.GoodDao.GetByName(pbAttachGood.Name)
				if err != nil {
					logs.Error(err)
					return nil, errors.New("Fail to finish GoodDao.GetByName")
				}
				if good == nil {
					logs.Error("Good not existed")
					return nil, errors.New("Good not existed")
				}
				// 1.3.2 插入附属记录
				if err := dao.AttachRecordDao.Create(orderID, thingID, int(sGood.ID), int(good.ID)); err != nil {
					logs.Error(err)
					return nil, errors.New("Fail to finish AttachRecordDao.Create")
				}

				// 1.3.3 插入订单记录
				if err := dao.OrderRecordDao.Create(orderID, int(good.ID), model.FlagOfAttachGood); err != nil {
					return nil, errors.New("Fail to finish OrderRecordDao.Create")
				}

				// 1.3.4 形成下一个物品号
				nextThingID()
			}
		}

	}

	// 2. 形成下一个订单号
	nextOrderID()

	res.OrderID = int64(orderID)
	return &res, nil
}

func (MvpServer) GetOrderGoods(ctx context.Context, req *pb.GetOrderGoodsReq) (*pb.GetOrderGoodsRes, error) {
	logs.Info("GetOrderGoods", ctx, req)

	var res pb.GetOrderGoodsRes

	// 1. 获取订单的商品
	orderRecords, err := dao.OrderRecordDao.GetByOrderID(int(req.OrderID))
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish OrderRecordDao.GetByOrderID")
	}

	// 2. 形成显示商品
	var pbGoods []*pb.Good
	for _, orderRecord := range orderRecords {
		if orderRecord.IsAttachGood == model.FlagOfNotAttachGood {
			// 2.1 获得商品名
			good, err := dao.GoodDao.Get(orderRecord.GoodID)
			if err != nil {
				logs.Error(err)
				return nil, errors.New("Fail to finish GoodDao.Get")
			}
			if good == nil {
				logs.Error("Good not existed")
				return nil, errors.New("Good not existed")
			}

			// 2.2 形成顶层显示商品
			pbGoods = append(pbGoods, &pb.Good{
				Name: good.Name,
			})
		} else {

			// 2.3 判断商品记录是否正确
			if len(pbGoods) == 0 {
				logs.Error("Order record error")
				return nil, errors.New("Order record error")
			}

			// 2.4 获得附属商品名
			good, err := dao.GoodDao.Get(orderRecord.GoodID)
			if err != nil {
				logs.Error(err)
				return nil, errors.New("Fail to finish GoodDao.Get")
			}
			if good == nil {
				logs.Error("Good not existed")
				return nil, errors.New("Good not existed")
			}

			// 2.5 形成附属显示商品
			pbGoods[len(pbGoods)-1].AttachGoods[good.Type].AttachGoods = append(
				pbGoods[len(pbGoods)-1].AttachGoods[good.Type].AttachGoods,
				&pb.AttachGood{
					Name: good.Name,
				},
			)
		}
	}

	// 3. 返回
	res.Goods = pbGoods
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

func (MvpServer) Checkout(ctx context.Context, req *pb.CheckoutReq) (*pb.CheckoutRes, error) {
	logs.Info("Checkout", ctx, req)
	var res pb.CheckoutRes

	// 1. 判断订单是否存在
	orderRecords, err := dao.OrderRecordDao.GetByOrderID(int(req.OrderID))
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish OrderRecordDao.GetByOrderID")
	}
	if orderRecords == nil {
		logs.Error("Order Record not Exist")
		return nil, errors.New("Order Record not Exist")
	}

	// 2. 获取未结账货物记录
	notCheckoutOrderRecords, err := dao.OrderRecordDao.GetNotCheckoutGoods(int(req.OrderID))
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Fail to finish OrderRecordDao.GetNotCheckoutGoods")
	}

	// 3. 计算价格 (先不考虑打折、满减等情况)
	var wholePrice float64
	for _, notCheckoutOrderRecord := range notCheckoutOrderRecords {
		// 3.1 判断商品是否存在
		good, err := dao.GoodDao.Get(notCheckoutOrderRecord.GoodID)
		if err != nil {
			logs.Error(err)
			return nil, errors.New("Fail to finish GoodDao.GetByName")
		}
		if good == nil {
			return nil, errors.New("Good not exist")
		}

		// 3.2 结账该商品
		wholePrice += good.Price
	}

	// 4. 返回结账金额
	res.Price = float32(wholePrice)
	return &res, nil
}

func (MvpServer) AddOptionClass(ctx context.Context, req *pb.AddOptionClassReq) (*pb.AddOptionClassRes, error) {
	logs.Info("AddOptionClass", ctx, req)
	var res pb.AddOptionClassRes

	// 1. 判断选项类是否存在
	optionClass, err := dao.OptionClassDao.Get(req.OptionClassName)
	if err != nil {
		logger.Error("Fail to get option class",
			zap.Any("optionClassName", req.OptionClassName),
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 如果选项类不存在，则创建
	if optionClass == nil {
		if err := dao.OptionClassDao.Create(req.OptionClassName); err != nil {
			logger.Error("Fail to create option class",
				zap.Any("optionClassName", req.OptionClassName),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
	}

	// 3. 创建选项类
	optionClass, _ = dao.OptionClassDao.Get(req.OptionClassName)

	// 4. 创建选项
	for _, optionName := range req.Options {
		if err := dao.OptionDao.Create(int(optionClass.ID), optionName); err != nil {
			logger.Error("Fail to create option",
				zap.Any("optionClassID", optionClass.ID),
				zap.Any("optionName", optionName),
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}
	}
	return &res, nil
}

func (MvpServer) GetAllOptionClasses(ctx context.Context, req *pb.GetAllOptionClassesReq) (*pb.GetAllOptionClassesRes, error) {
	logs.Info(" GetAllOptionClasses", ctx, req)
	var res pb.GetAllOptionClassesRes

	// 1. 获取
	optionClasses, err := dao.OptionClassDao.GetAll()
	if err != nil {
		logger.Error("Fail to get all option class",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 转换
	for _, optionClass := range optionClasses {
		res.OptionClassNames = append(res.OptionClassNames, optionClass.Name)
	}

	return &res, nil
}
