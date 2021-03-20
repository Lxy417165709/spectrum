package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
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
