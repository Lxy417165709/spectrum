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
func formPbGood(good *model.Good) (*pb.Good, error) {
	if good == nil{
		return nil,nil
	}

	// 1. 形成初始字段
	pbGood := &pb.Good{
		Name:             good.Name,
		Type:             int64(good.Type),
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

	optionClassIDs := make([]int,0)
	for _,record := range goodOptionClassRecords{
		optionClassIDs = append(optionClassIDs,record.OptionClassID)
	}
	optionClasses, err := dao.OptionClassDao.GetByIDs(optionClassIDs)
	if err!=nil{
		logger.Error("Fail to get option class",
			zap.Any("optionClassIDs", optionClassIDs),
			zap.Error(err))
		return nil, err
	}

	for _,optionClass :=range optionClasses {
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
	if optionClass == nil{
		return nil,nil
	}

	// 0. 初始化 pb.OptionClass
	pbOptionClass := &pb.OptionClass{
		ClassName: optionClass.Name,
	}

	// 1. 获得选项
	options, err := dao.OptionDao.GetByOptionClassID(int(optionClass.ID))
	if err != nil {
		logger.Error("Fail to get all options",
			zap.Error(err))
		return nil, err
	}

	// 2. 获得选项名
	var optionNames []string
	for _, option := range options {
		optionNames = append(optionNames, option.Name)
	}

	// 3. 更新 pb.OptionClass
	pbOptionClass.OptionNames = optionNames

	return pbOptionClass, nil
}
