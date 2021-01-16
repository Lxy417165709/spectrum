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
		goods = append(goods, getGood(0, mainElementName))
	}
	return goods
}

func getDesk(deskID int64) *pb.Desk {
	desk, err := dao.DeskDao.Get(deskID)
	if err != nil {
		// todo: log
		return nil
	}

	space, err := dao.SpaceDao.Get(desk.SpaceName, int64(desk.SpaceNum))
	if err != nil {
		// todo: log
		return nil
	}

	return &pb.Desk{
		Goods:          getDeskGoods(deskID),
		StartTimestamp: desk.StartTimestamp,
		EndTimestamp:   desk.EndTimestamp,
		Space:          space.ToPb(),
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
		goods = append(goods, getGood(int64(dbGood.ID), dbGood.Name))
	}
	return goods
}

func getGood(goodID int64, mainElementName string) *pb.Good {
	return &pb.Good{
		Id:             goodID,
		MainElement:    getMainElement(goodID, mainElementName),
		AttachElements: getAttachElements(goodID, mainElementName),
	}
}

func getMainElement(goodID int64, mainElementName string) *pb.Element {
	mainElements, err := dao.ElementDao.GetByName(mainElementName)
	if err != nil {
		// todo: log
		return nil
	}
	var sizeRecord *model.MainElementSizeRecord
	if goodID == 0 {
		sizeRecord, err = dao.MainElementSizeRecordDao.GetByMainElementName(mainElementName)
	} else {
		sizeRecord, err = dao.MainElementSizeRecordDao.GetByGoodID(goodID)
	}
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
		SizeInfos: getSizeInfos(sizeRecord.SelectSize, mainElements),
	}
}

func getAttachElements(goodID int64, mainElementName string) []*pb.Element {
	var attachElements []*pb.Element
	var attachRecords []*model.MainElementAttachElementRecord
	var err error
	if goodID == 0 {
		attachRecords, err = dao.MainElementAttachElementRecordDao.GetByMainElementName(mainElementName)
	} else {
		attachRecords, err = dao.MainElementAttachElementRecordDao.GetByGoodID(goodID)
	}
	if err != nil {
		logger.Error("Fail to finish MainElementAttachElementRecordDao.GetByMainElementName", zap.Error(err))
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

