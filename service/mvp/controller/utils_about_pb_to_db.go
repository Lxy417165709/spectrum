package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

//func checkOutIfNot(chargeableObj model.Chargeable) error {
//
//	expenseInfo := getExpenseInfo(chargeableObj)
//	if expenseInfo.CheckOutAt != 0 {
//		// todo: 警告
//		// 这里表示商品已结账过了
//		return nil
//	}
//	expenseInfo.CheckOutAt = time.Now().Unix()
//	if err := dao.ChargeableObjectDao.UpdateExpenseInfo(chargeableObj, expenseInfo); err != nil {
//		logger.Error("Fail to finish chargeableDao.Update", zap.Error(err))
//		return err
//	}
//
//	// 添加结账记录
//	if err := dao.CheckOutRecordDao.Create(&model.CheckOutRecord{
//		ChargeableObjectName: chargeableObj.GetName(),
//		ChargeableObjectID:   chargeableObj.GetID(),
//		CheckOutAt:           expenseInfo.CheckOutAt,
//	}); err != nil {
//		logger.Error("Fail to finish CheckOutRecordDao.Create", zap.Error(err))
//		return err
//	}
//
//	return nil
//}

func toDbElementSelectSizeRecord(goodID, elementID, elementSizeInfoID int64) *model.ElementSelectSizeRecord {
	return &model.ElementSelectSizeRecord{
		GoodID:           goodID,
		ElementID:        elementID,
		SelectSizeInfoID: elementSizeInfoID,
	}
}

func getPbElementSelectSizeInfoID(pbElement *pb.Element) int64 {
	return model.GetPbElementSelectSizeInfo(pbElement).Id
}

func writePbElementSelectSizeRecord(goodID int64, pbElement *pb.Element) error {
	elementSelectedSizeInfoID := getPbElementSelectSizeInfoID(pbElement)
	if _, errResult := dao.ElementSelectSizeRecordDao.Create(toDbElementSelectSizeRecord(goodID, pbElement.Id, elementSelectedSizeInfoID));
		errResult != nil {
		return errResult
	}
	return nil
}

func writePbGoodSizeInfoToDB(good *pb.Good) error {
	// 1. 创建主元素、主元素尺寸的对应关系
	if errResult := writePbElementSelectSizeRecord(good.Id, good.MainElement); errResult != nil {
		return errResult
	}

	// 2. 创建附属元素、附属元素尺寸的对应关系
	for _, attachElement := range good.AttachElements {
		if errResult := writePbElementSelectSizeRecord(good.Id, attachElement); errResult != nil {
			return errResult
		}
	}

	// 3. 创建主元素、附属元素、附属元素尺寸的对应关系
	for _, attachElement := range good.AttachElements {
		if _, errResult := dao.MainElementAttachElementRecordDao.Create(&model.MainElementAttachElementRecord{
			GoodID:          good.Id,
			AttachElementID: attachElement.Id,
			MainElementID:   good.MainElement.Id,
		}); errResult != nil {
			return errResult
		}
	}

	// 3. 返回
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

// 将Pb元素写入数据库，并更新其ID字段
func writePbElementToDbAndUpdateID(pbElement *pb.Element, classID int64) error {
	// 1. 将元素写入数据库，如果先前该元素不存在，则更新元素ID
	elementID, errResult := dao.ElementDao.Create(toDbElement(pbElement, classID))
	if errResult != nil {
		return errResult
	}
	if pbElement.Id == 0 {
		pbElement.Id = elementID
	}

	// 2. 将元素的尺寸信息写入数据库，如果先前尺寸不存在，则更新元素的尺寸ID
	for _, pbSizeInfo := range pbElement.SizeInfos {
		sizeInfoID, errResult := dao.ElementSizeInfoDao.Create(toDbElementSizeInfo(pbSizeInfo, pbElement.Id))
		if errResult != nil {
			return errResult
		}
		if pbSizeInfo.Id == 0 {
			pbSizeInfo.Id = sizeInfoID
		}
	}

	// 3. 返回
	return nil
}

func toDbElementSizeInfo(pbSizeInfo *pb.SizeInfo, elementID int64) *model.ElementSizeInfo {
	return &model.ElementSizeInfo{
		ID:               pbSizeInfo.Id,
		ElementID:        elementID,
		Size:             pbSizeInfo.Size,
		Price:            model.GetDbPrice(pbSizeInfo.Price),
		PictureStorePath: pbSizeInfo.PictureStorePath,
	}
}

func toDbElement(pbElement *pb.Element, classID int64) *model.Element {
	return &model.Element{
		ID:      pbElement.Id,
		Name:    pbElement.Name,
		Type:    pbElement.Type,
		ClassID: classID,
	}
}

func toDbSpace(pbSpace *pb.Space, classID int64) *model.Space {
	return &model.Space{
		ID:               pbSpace.Id,
		Name:             pbSpace.Name,
		ClassID:          classID,
		Price:            model.GetDbPrice(pbSpace.Price),
		BillingType:      pbSpace.BillingType,
		PictureStorePath: pbSpace.PictureStorePath,
	}
}
