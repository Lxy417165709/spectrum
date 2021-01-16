package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

func getClassGoods(className string) []*pb.Good {
	var goods []*pb.Good
	for _, mainElementName := range getElementNames(className) {
		goods = append(goods, getGood(&model.Good{
			Name: mainElementName,
		}))
	}
	return goods
}

// 返回的 desk:
// 已结账时: 返回结账的金额信息
// 未结账时: 返回最新的金额信息
func getDesk(deskID int64) *pb.Desk {
	desk, err := dao.DeskDao.Get(deskID)
	if err != nil {
		// todo: log
		return nil
	}
	space, err := dao.SpaceDao.Get(desk.SpaceName, desk.SpaceNum)
	if err != nil {
		// todo: log
		return nil
	}
	favor := getFavors(desk)
	return &pb.Desk{
		Id:             deskID,
		Space:          space.ToPb(),
		StartTimestamp: desk.StartTimestamp,
		EndTimestamp:   desk.EndTimestamp,
		Goods:          getDeskGoods(deskID),
		Favors:         favor,
		ExpenseInfo:    desk.GetExpenseInfo(space.Price, favor),
	}
}

func getDeskGoods(deskID int64) []*pb.Good {
	dbGoods, err := dao.GoodDao.GetByDeskID(deskID)
	if err != nil {
		// todo: log
		return nil
	}
	var goods []*pb.Good
	for _, dbGood := range dbGoods {
		goods = append(goods, getGood(dbGood))
	}
	return goods
}

// 返回的 good:
// 已结账时: 返回结账的金额信息
// 未结账时: 返回最新的金额信息
func getGood(good *model.Good) *pb.Good {
	mainElement := getMainElement(int64(good.ID), good.Name)
	attachElements := getAttachElements(int64(good.ID), good.Name)
	favors := getFavors(good)
	return &pb.Good{
		Id:             int64(good.ID),
		MainElement:    mainElement,
		AttachElements: attachElements,
		Favors:         favors,
		ExpenseInfo:    good.GetExpenseInfo(mainElement, attachElements, favors),
	}
}

func getMainElement(goodID int64, mainElementName string) *pb.Element {
	mainElements, err := dao.ElementDao.GetByName(mainElementName)
	if err != nil {
		// todo: log
		return nil
	}
	sizeRecord, err := dao.MainElementSizeRecordDao.GetByGoodIdAndMainElementName(goodID, mainElementName)
	if err != nil {
		// todo: log
		return nil
	}
	if sizeRecord == nil {
		// todo: log
		return nil
	}
	return &pb.Element{
		Name:      mainElementName,
		Type:      pb.ElementType_Main,
		SizeInfos: getSizeInfos(sizeRecord.SelectSize, mainElements),
	}
}

func getAttachElements(goodID int64, mainElementName string) []*pb.Element {
	var attachElements []*pb.Element
	attachRecords, err := dao.MainElementAttachElementRecordDao.GetByGoodIdAndMainElementName(goodID, mainElementName)
	if err != nil {
		logger.Error("Fail to finish MainElementAttachElementRecordDao.GetByGoodIdAndMainElementName", zap.Error(err))
		return nil
	}
	for _, attachRecord := range attachRecords {
		elements, err := dao.ElementDao.GetByName(attachRecord.AttachElementName)
		if err != nil {
			logger.Error("Fail to finish ElementDao.GetByName", zap.Error(err))
			return nil
		}
		attachElements = append(attachElements, &pb.Element{
			Name:      attachRecord.AttachElementName,
			SizeInfos: getSizeInfos(attachRecord.SelectSize, elements),
		})
	}
	return attachElements
}

func getFavors(chargeableObj model.Chargeable) []*pb.Favor {
	records, err := dao.FavorRecordDao.Get(chargeableObj.GetName(), chargeableObj.GetID())
	if err != nil {
		// todo: log
		return nil
	}
	result := make([]*pb.Favor, 0)
	for _, record := range records {
		result = append(result, record.ToPb())
	}
	return result
}

func getSizeInfos(selectSize string, sameNameElements []*model.Element) []*pb.SizeInfo {
	var sizeInfos []*pb.SizeInfo
	for _, element := range sameNameElements {
		sizeInfos = append(sizeInfos, &pb.SizeInfo{
			Size:             element.Size,
			Price:            element.Price,
			PictureStorePath: element.PictureStorePath,
			IsSelected:       selectSize == element.Size,
		})
	}
	return sizeInfos
}
