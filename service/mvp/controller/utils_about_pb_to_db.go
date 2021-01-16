package controller

import (
	"fmt"
	"go.uber.org/zap"
	"reflect"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"strings"
	"time"
)

func writeFavorToDB(favorableStruct interface{}) error {

	var chargeableObject model.Chargeable
	var favors []*pb.Favor
	switch favorableStruct.(type) {
	case *pb.Good:
		good := favorableStruct.(*pb.Good)
		chargeableObject = &model.Good{}
		chargeableObject.SetID(good.Id)
		favors = good.Favors
	case *pb.Desk:
		desk := favorableStruct.(*pb.Desk)
		chargeableObject = &model.Desk{}
		chargeableObject.SetID(desk.Id)
		favors = desk.Favors
	default:
		err := fmt.Errorf("unfix type")
		logger.Error("Unfix type", zap.String("type", reflect.TypeOf(favorableStruct).String()), zap.Error(err))
		return err
	}

	for _, favor := range favors {
		if err := dao.FavorRecordDao.Create(&model.FavorRecord{
			ChargeableObjectName: chargeableObject.GetName(),
			ChargeableObjectID:   chargeableObject.GetID(),
			FavorType:            favor.FavorType,
			FavorParameters:      strings.Join(favor.Parameters, "|"),
		}); err != nil {
			// todo: log
			return err
		}
	}

	return nil
}

// todo: 不能重复结账
func checkOut(chargeableObj model.Chargeable, ids []int64) error {
	for _, id := range ids {
		checkOutTimestamp := time.Now().Unix()
		if err := dao.CheckOutRecordDao.Create(&model.CheckOutRecord{
			ChargeableObjectName: chargeableObj.GetName(),
			ChargeableObjectID:   id,
			CheckOutTimestamp:    checkOutTimestamp,
		}); err != nil {
			logger.Error("Fail to finish CheckOutRecordDao.Create", zap.Error(err))
			return err
		}
		chargeableObj.SetID(id)
		expenseInfo, chargeableDao := getExpenseInfoAndChargeableDao(chargeableObj)
		// todo: chargeableDao 接口函数有待完善
		if err := chargeableDao.Update(map[string]interface{}{
			"id":                  id,
			"check_out_timestamp": checkOutTimestamp,
			"expense":             expenseInfo.Expense,
			"non_favor_expense":   expenseInfo.NonFavorExpense,
		}); err != nil {
			logger.Error("Fail to finish chargeableDao.Update", zap.Error(err))
			return err
		}
	}
	return nil
}

func writeGoodSizeToDB(good *pb.Good) error {
	if len(good.MainElement.SizeInfos) == 0 {
		return nil
	}

	// 1. 创建主元素、主元素尺寸的对应关系
	if err := dao.MainElementSizeRecordDao.Create(&model.MainElementSizeRecord{
		GoodID:          good.Id,
		MainElementName: good.MainElement.Name,
		SelectSize:      good.MainElement.SizeInfos[0].Size,
	}); err != nil {
		logger.Error("Fail to finish MainElementSizeRecordDao.Create", zap.Error(err))
		return err
	}

	// 2. 创建主元素、附属元素、附属元素尺寸的对应关系
	for _, attachElement := range good.AttachElements {
		if err := dao.MainElementAttachElementRecordDao.Create(&model.MainElementAttachElementRecord{
			GoodID:            good.Id,
			MainElementName:   good.MainElement.Name,
			AttachElementName: attachElement.Name,
			SelectSize:        model.GetSelectSizeInfo(attachElement.SizeInfos).Size,
		}); err != nil {
			logger.Error("Fail to finish MainElementAttachElementDao.Create", zap.Error(err))
			return err
		}
	}
	return nil
}

func closeDeskIfOpening(deskID int64, endTimestamp int64) error {
	desk, err := dao.DeskDao.Get(deskID)
	if err != nil {
		logger.Error("Fail to finish DeskDao.Get",
			zap.Error(err))
		return err
	}
	if !desk.IsOpening() {
		logger.Warn("Desk had been closed", zap.Int64("deskID", deskID))
		return nil
	}
	if err := dao.DeskDao.Update(map[string]interface{}{
		"id":            deskID,
		"end_timestamp": endTimestamp,
	}); err != nil {
		logger.Error("Fail to finish DeskDao.Update",
			zap.Error(err))
		return err
	}
	return nil
}

func createElement(pbElement *pb.Element, className string) error {
	dbElements := getDbElements(pbElement, className)
	for _, dbElement := range dbElements {
		if err := dao.ElementDao.Create(dbElement); err != nil {
			logger.Error("Fail to finish ElementDao.Create", zap.Error(err))
			return err
		}
	}
	return nil
}

func getDbElements(pbElement *pb.Element, className string) []*model.Element {
	var result []*model.Element

	for _, sizeInfo := range pbElement.SizeInfos {
		result = append(result, &model.Element{
			Name:             pbElement.Name,
			Type:             pbElement.Type,
			ClassName:        className,
			Size:             sizeInfo.Size,
			Price:            sizeInfo.Price,
			PictureStorePath: sizeInfo.PictureStorePath,
		})
	}
	return result
}

func getDbSpace(pbSpace *pb.Space) *model.Space {
	return &model.Space{
		Name:          pbSpace.Name,
		Num:           pbSpace.Num,
		Price:         pbSpace.Price,
		PriceRuleType: pbSpace.PriceRuleType,
	}
}