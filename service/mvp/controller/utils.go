package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

// ------------------------------------- 自增相关 -------------------------------------
var curOrderID int // 这里不考虑高并发

func getOrderID() int {
	return curOrderID
}

func nextOrderID() {
	curOrderID++
}

var curThingID int // 这里不考虑高并发

func getThingID() int {
	return curThingID
}

func nextThingID() {
	curThingID++
}

// ------------------------------------- 转换相关 -------------------------------------
func formPbGoodClass(goodClass *model.GoodClass) (*pb.GoodClass, error) {
	if goodClass == nil {
		return nil, nil
	}

	// 1. 形成初始字段
	pbGoodClass := &pb.GoodClass{
		Name: goodClass.Name,
		ClassType: pb.ClassType(goodClass.ClassType),
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
