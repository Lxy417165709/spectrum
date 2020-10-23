package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

// ------------------------------------- 自增相关 -------------------------------------
var curOrderID = 1 // 这里不考虑高并发

func getSetOrderID() int {
	curOrderID++
	return curOrderID - 1
}

var curThingID = 1 // 这里不考虑高并发

func getSetThingID() int {
	curThingID++
	return curThingID - 1
}

// ------------------------------------- 转换相关 -------------------------------------
func formPbGoodClass(goodClass *model.GoodClass) (*pb.GoodClass, error) {
	if goodClass == nil {
		return nil, nil
	}

	// 1. 形成初始字段
	pbGoodClass := &pb.GoodClass{
		Name:            goodClass.Name,
		ClassType:       pb.ClassType(goodClass.ClassType),
		SelectGoodNames: []string{},
	}

	// 2. 获得该类的商品
	goods, err := dao.GoodDao.GetByGoodClassID(int(goodClass.ID))
	if err != nil {
		logger.Error("Fail to finish GoodDao.GetByGoodClassID",
			zap.Any("goodID", goodClass.ID),
			zap.Error(err))
		return nil, err
	}

	// 3. 转换
	for _, good := range goods {
		pbGood, err := formPbGood(good)
		if err != nil {
			logger.Error("Fail to form pb.Good",
				zap.Any("tableGood", good),
				zap.Error(err))
			return nil, err
		}
		pbGoodClass.Goods = append(pbGoodClass.Goods, pbGood)
	}

	return pbGoodClass, nil
}

func formPbGood(good *model.Good) (*pb.Good, error) {
	if good == nil {
		return nil, nil
	}

	// 1. 形成初始字段
	pbGood := &pb.Good{
		Name:             good.Name,
		Price:            float32(good.Price),
		PictureStorePath: good.PictureStorePath,
	}

	// 2. 获得选项类记录
	goodOptionClassRecords, err := dao.GoodOptionClassRecordDao.GetByGoodID(int(good.ID))
	if err != nil {
		logger.Error("Fail to get all good option class records",
			zap.Any("goodID", good.ID),
			zap.Error(err))
		return nil, err
	}

	optionClassIDs := make([]int, 0)
	for _, record := range goodOptionClassRecords {
		optionClassIDs = append(optionClassIDs, record.OptionClassID)
	}
	optionClasses, err := dao.OptionClassDao.GetByIDs(optionClassIDs)
	if err != nil {
		logger.Error("Fail to get option class",
			zap.Any("optionClassIDs", optionClassIDs),
			zap.Error(err))
		return nil, err
	}

	for _, optionClass := range optionClasses {
		pbOptionClass, err := formPbOptionClass(optionClass)
		if err != nil {
			logger.Error("Fail to form pb.OptionClass",
				zap.Any("tableOptionClass", optionClass),
				zap.Error(err))
			return nil, err
		}
		pbGood.OptionClasses = append(pbGood.OptionClasses, pbOptionClass)
	}

	// 3. 获得附属产品类记录
	goodAttachClassRecords, err := dao.GoodAttachClassRecordDao.GetByGoodID(int(good.ID))
	if err != nil {
		logger.Error("Fail to get all good attach class records",
			zap.Any("goodID", good.ID),
			zap.Error(err))
		return nil, err
	}
	attachGoodClassIDs := make([]int, 0)
	for _, record := range goodAttachClassRecords {
		attachGoodClassIDs = append(attachGoodClassIDs, record.AttachGoodClassID)
	}
	attachGoodClasses, err := dao.GoodClassDao.GetByIDs(attachGoodClassIDs)
	if err != nil {
		logger.Error("Fail to get attach good class",
			zap.Any("attachGoodClassIDs", attachGoodClassIDs),
			zap.Error(err))
		return nil, err
	}
	for _, attachGoodClass := range attachGoodClasses {
		pbAttachGoodClass, err := formPbGoodClass(attachGoodClass)
		if err != nil {
			logger.Error("Fail to form pb.PbGoodClass",
				zap.Any("tableAttachGoodClass", attachGoodClass),
				zap.Error(err))
			return nil, err
		}
		pbGood.AttachGoodClasses = append(pbGood.AttachGoodClasses, pbAttachGoodClass)
	}

	return pbGood, nil
}

func formPbOptionClass(optionClass *model.OptionClass) (*pb.OptionClass, error) {
	if optionClass == nil {
		return nil, nil
	}

	// 1. 初始化
	pbOptionClass := &pb.OptionClass{
		Name:              optionClass.Name,
		SelectOptionIndex: int32(optionClass.DefaultSelectOptionIndex),
	}

	// 2. 形成选项类
	options, err := dao.OptionDao.GetByOptionClassID(int(optionClass.ID))
	if err != nil {
		logger.Error("Fail to get all options",
			zap.Error(err))
		return nil, err
	}

	// 3. 形成选项
	var optionNames []string
	for _, option := range options {
		pbOptionClass.Options = append(pbOptionClass.Options, &pb.Option{
			Name: option.Name,
		})
		optionNames = append(optionNames, option.Name)
	}
	return pbOptionClass, nil
}

func getOrderPrice(things []*model.Thing) float64 {
	price := 0.0
	for _, thing := range things {
		price += thing.Price
	}
	return price
}

func getThingPrice(good *model.Good, attachGoods []*model.Good) float64 {
	price := good.Price
	for _, attachGood := range attachGoods {
		price += attachGood.Price
	}
	return price
}

func getOrderLog(orderID int) (*pb.OrderLog, error) {
	var orderLog pb.OrderLog

	// 0. 获得订单
	order, err := dao.OrderDao.Get(orderID)
	if err != nil {
		logger.Error("Fail to finish OrderDao.Get",
			zap.Any("orderID", orderID),
			zap.Error(err))
		return nil, err
	}
	if order == nil {
		logger.Warn("Order not exist", zap.Any("orderID", orderID),
			zap.Error(err))
		return nil, err
	}
	orderLog.Price = float32(order.Price)
	orderLog.HasCheckedOut = order.HasCheckedOut == model.HasCheckedOut

	// 1. 获得订单-物品记录
	orderThingRecords, err := dao.OrderThingRecordDao.GetByOrderID(int(orderID))
	if err != nil {
		logger.Error("Fail to finish OrderThingRecordDao.GetByOrderID",
			zap.Error(err))
		return nil, err
	}

	// 2. 分物品处理
	for _, orderThingRecord := range orderThingRecords {
		goodLog, err := getGoodLog(orderThingRecord.ThingID)
		if err != nil {
			logger.Error("Fail to finish getGoodLog",
				zap.Int("thingID", orderThingRecord.ThingID),
				zap.Error(err))
			return nil, err
		}
		// ---------------- 贡献 orderLog ----------------
		orderLog.GoodLogs = append(orderLog.GoodLogs, goodLog)
	}
	return &orderLog, nil
}

func getGoodLog(thingID int) (*pb.GoodLog, error) {
	var goodLog pb.GoodLog
	// ---------------- thing ----------------
	thing, err := dao.ThingDao.Get(thingID)
	if err != nil {
		logger.Error("Fail to finish ThingDao.Get",
			zap.Any("thingID", thingID),
			zap.Error(err))
		return nil, err
	}

	// ---------------- thing-option ----------------
	thingOptionRecords, err := dao.ThingOptionRecordDao.GetByThingID(int(thing.ID))
	if err != nil {
		logger.Error("Fail to finish ThingOptionRecordDao.GetByThingID",
			zap.Any("thingID", thingID),
			zap.Error(err))
		return nil, err
	}

	optionClassNameToOptionLogs := make(map[string][]*pb.OptionLog)
	for _, thingOptionRecord := range thingOptionRecords {
		option, err := dao.OptionDao.Get(thingOptionRecord.OptionID)
		if err != nil {
			logger.Error("Fail to finish OptionDao.Get",
				zap.Any("optionID", thingOptionRecord.OptionID),
				zap.Error(err))
			return nil, err
		}

		optionClass, err := dao.OptionClassDao.Get(option.OptionClassID)
		if err != nil {
			logger.Error("Fail to finish OptionClassDao.Get",
				zap.Any("optionClassID", option.OptionClassID),
				zap.Error(err))
			return nil, err
		}
		if optionClass != nil {
			optionClassNameToOptionLogs[optionClass.Name] = append(
				optionClassNameToOptionLogs[optionClass.Name],
				&pb.OptionLog{Name: option.Name},
			)
		}
	}

	// ---------------- thing-attachGood ----------------
	thingAttachGoodRecords, err := dao.ThingAttachGoodRecordDao.GetByThingID(int(thing.ID))
	if err != nil {
		logger.Error("Fail to finish ThingAttachGoodRecordDao.GetByThingID",
			zap.Any("thingID", thingID),
			zap.Error(err))
		return nil, err
	}

	attachGoodClassNameToGoodLogs := make(map[string][]*pb.AttachGoodLog)
	for _, thingAttachGoodRecord := range thingAttachGoodRecords {
		attachGood, err := dao.GoodDao.Get(thingAttachGoodRecord.AttachGoodID)
		if err != nil {
			logger.Error("Fail to finish OptionDao.Get",
				zap.Any("attachGoodID", thingAttachGoodRecord.AttachGoodID),
				zap.Error(err))
			return nil, err
		}

		attachGoodClass, err := dao.GoodClassDao.Get(attachGood.ClassID)
		if err != nil {
			logger.Error("Fail to finish GoodClassDao.Get",
				zap.Any("attachGoodClassID", attachGood.ClassID),
				zap.Error(err))
			return nil, err
		}
		if attachGoodClass != nil {
			attachGoodClassNameToGoodLogs[attachGoodClass.Name] = append(
				attachGoodClassNameToGoodLogs[attachGoodClass.Name],
				&pb.AttachGoodLog{
					Name:  attachGood.Name,
					Price: float32(attachGood.Price),
				},
			)
		}
	}

	// ---------------- goodLog ----------------
	good, err := dao.GoodDao.Get(thing.GoodID)
	if err != nil {
		logger.Error("Fail to finish GoodDao.Get",
			zap.Any("goodID", thing.GoodID),
			zap.Error(err))
		return nil, err
	}
	goodLog.Price = float32(thing.Price)
	goodLog.Name = good.Name

	for optionClassName, optionLogs := range optionClassNameToOptionLogs {
		if len(optionLogs) == 0 {
			goodLog.OptionClassLogs = append(goodLog.OptionClassLogs, &pb.OptionClassLog{
				Name: optionClassName,
			})
			continue
		}
		goodLog.OptionClassLogs = append(goodLog.OptionClassLogs, &pb.OptionClassLog{
			Name:      optionClassName,
			OptionLog: optionLogs[0],
		})
	}
	for attachGoodClassName, attachGoods := range attachGoodClassNameToGoodLogs {
		goodLog.AttachGoodClassLogs = append(goodLog.AttachGoodClassLogs, &pb.AttachGoodClassLog{
			Name:           attachGoodClassName,
			AttachGoodLogs: attachGoods,
		})
	}
	return &goodLog, nil
}
