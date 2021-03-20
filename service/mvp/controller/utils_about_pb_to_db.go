package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"time"
)

func checkOutIfNot(chargeableObj model.Chargeable) error {

	expenseInfo := getExpenseInfo(chargeableObj)
	if expenseInfo.CheckOutAt != 0 {
		// todo: 警告
		// 这里表示商品已结账过了
		return nil
	}
	expenseInfo.CheckOutAt = time.Now().Unix()
	if err := dao.ChargeableObjectDao.UpdateExpenseInfo(chargeableObj, expenseInfo); err != nil {
		logger.Error("Fail to finish chargeableDao.Update", zap.Error(err))
		return err
	}

	// 添加结账记录
	if err := dao.CheckOutRecordDao.Create(&model.CheckOutRecord{
		ChargeableObjectName: chargeableObj.GetName(),
		ChargeableObjectID:   chargeableObj.GetID(),
		CheckOutAt:           expenseInfo.CheckOutAt,
	}); err != nil {
		logger.Error("Fail to finish CheckOutRecordDao.Create", zap.Error(err))
		return err
	}

	return nil
}

func writePbGoodSizeToDB(good *pb.Good) error {
	// 1. 创建主元素、主元素尺寸的对应关系
	if _, errResult := dao.ElementSizeRecordDao.Create(&model.ElementSizeRecord{
		GoodID:      good.Id,
		ElementName: good.MainElement.Name,
		SelectSize:  good.MainElement.SizeInfos[good.MainElement.SelectedIndex].Size,
	}); errResult != nil {
		return errResult
	}

	// 2. 创建主元素、附属元素、附属元素尺寸的对应关系
	for _, attachElement := range good.AttachElements {
		if _, errResult := dao.MainElementAttachElementRecordDao.Create(&model.MainElementAttachElementRecord{
			GoodID:            good.Id,
			MainElementName:   good.MainElement.Name,
			AttachElementName: attachElement.Name,
			SelectSize:        model.GetSelectSizeInfo(attachElement).Size,
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

func storePbElementToDB(pbElement *pb.Element, className string) error {
	dbElements := getDbElements(pbElement, className)
	for _, dbElement := range dbElements {
		if _, errResult := dao.ElementDao.Create(dbElement); errResult != nil {
			return errResult
		}
	}
	return nil
}

func getDbElements(pbElement *pb.Element, className string) []*model.Element {
	if pbElement == nil {
		return nil
	}
	var result []*model.Element
	for _, sizeInfo := range pbElement.SizeInfos {
		result = append(result, &model.Element{
			ClassName:        className,
			Name:             pbElement.Name,
			Type:             pbElement.Type,
			Price:            model.GetDbPrice(sizeInfo.Price),
			Size:             sizeInfo.Size,
			PictureStorePath: sizeInfo.PictureStorePath,
		})
	}
	return result
}

func getDbSpace(pbSpace *pb.Space) *model.Space {

	return &model.Space{
		Name:             pbSpace.Name,
		ClassName:        pbSpace.ClassName,
		Price:            model.GetDbPrice(pbSpace.Price),
		BillingType:      pbSpace.BillingType,
		PictureStorePath: pbSpace.PictureStorePath,
	}
}
