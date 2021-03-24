package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"strconv"
	"time"
)

//func getExpenseInfo(chargeableObj model.Chargeable) *pb.ExpenseInfo {
//	switch chargeableObj.(type) {
//	case *model.Good:
//		return getPbGood(chargeableObj.(*model.Good), "todo").ExpenseInfo
//	case *model.Desk:
//		return getPbDesk(chargeableObj.(*model.Desk)).ExpenseInfo
//	default:
//		panic("unfix type")
//	}
//}

func getSelectedIndex(sizeInfos []*pb.SizeInfo, selectSizeID int64) int {
	for index, sizeInfo := range sizeInfos {
		if sizeInfo.Id == selectSizeID {
			return index
		}
	}
	logger.Warn("Can not get selected Index", zap.Any("sizeInfos", sizeInfos), zap.Any("selectSizeID", selectSizeID))
	return 0
}

func getDbSpaceClassByID(classID int64) *model.SpaceClass {
	spaceClass, errResult := dao.SpaceClassDao.Get(classID)
	if errResult != nil {
		// todo: log
		return nil
	}
	return spaceClass
}

func getDbElementClassByID(classID int64) *model.ElementClass {
	elementClass, errResult := dao.ElementClassDao.Get(classID)
	if errResult != nil {
		// todo: log
		return nil
	}
	return elementClass
}

func getDbSpaceClassByName(className string) *model.SpaceClass {
	spaceClass, errResult := dao.SpaceClassDao.GetByName(className)
	if errResult != nil {
		// todo: log
		return nil
	}
	return spaceClass
}

func getDbElementClassByName(className string) *model.ElementClass {
	elementClass, errResult := dao.ElementClassDao.GetByName(className)
	if errResult != nil {
		// todo: log
		return nil
	}
	return elementClass
}

func getDbPrice(priceString string) float64 {
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		logger.Error("Fail to finish strconv.ParseFloat", zap.String("priceString", priceString))
	}
	return price
}

func toTime(timestamp int64) time.Time {
	if timestamp != 0 {
		return time.Unix(timestamp, 0)
	}
	return model.NilTime
}
