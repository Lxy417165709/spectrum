package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"time"
)

// 如果未结账，则结账
func checkOutIfNot(chargeableObj model.Chargeable) error {
	expenseInfo := getExpenseInfo(chargeableObj)
	if expenseInfo.CheckOutTimestamp != 0 {
		// todo: 警告
		// 这里表示商品已结账过了
		return nil
	}
	expenseInfo.CheckOutTimestamp = time.Now().Unix()
	if err := dao.ChargeableObjectDao.UpdateExpenseInfo(chargeableObj, expenseInfo); err != nil {
		logger.Error("Fail to finish chargeableDao.Update", zap.Error(err))
		return err
	}

	// 添加结账记录
	if err := dao.ChargeableObjectDao.CreateCheckOutRecord(&model.CheckOutRecord{
		ChargeableObjectName: chargeableObj.GetName(),
		ChargeableObjectID:   chargeableObj.GetID(),
		CheckOutTimestamp:    expenseInfo.CheckOutTimestamp,
	}); err != nil {
		logger.Error("Fail to finish CheckOutRecordDao.Create", zap.Error(err))
		return err
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

func writeChargeableObjectInfoToDbAndAttachID(pbChargeableObject pb.Chargeable, attachValues ...interface{}) error {
	dbChargeableObject := getDbChargeableObject(pbChargeableObject, attachValues...)
	if err := dao.ChargeableObjectDao.Create(dbChargeableObject); err != nil {
		// todo:log
		return err
	}
	dbToPbAttachID(pbChargeableObject, dbChargeableObject)
	if err := dao.ChargeableObjectDao.CreateFavorRecord(pbChargeableObject); err != nil {
		// todo: log
		return err
	}
	return nil
}

func getDbChargeableObject(pbChargeableObject pb.Chargeable, attachValues ...interface{}) model.Chargeable {
	// todo: 校验参数合法性
	switch pbChargeableObject.(type) {
	case *pb.Good:
		good := pbChargeableObject.(*pb.Good)
		return &model.Good{
			Name:              good.MainElement.Name,
			DeskID:            attachValues[0].(int64),
			Expense:           good.ExpenseInfo.Expense,
			CheckOutTimestamp: good.ExpenseInfo.CheckOutTimestamp,
			NonFavorExpense:   good.ExpenseInfo.NonFavorExpense,
		}
	case *pb.Desk:
		desk := pbChargeableObject.(*pb.Desk)
		return &model.Desk{
			SpaceName:      desk.Space.Name,
			SpaceNum:       desk.Space.Num,
			StartTimestamp: attachValues[0].(int64),
		}
	default:
		panic("unfix type")
	}
}
