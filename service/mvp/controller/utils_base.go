package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

func getElementNames(className string) []string {
	var isLiving = make(map[string]bool)
	elements, err := dao.ElementDao.GetByClassName(className)
	var elementNames []string
	if err != nil {
		logger.Error("Fail to finish ElementDao.GetByClassName", zap.Error(err))
		return nil
	}
	for _, element := range elements {
		if isLiving[element.Name] {
			continue
		}
		isLiving[element.Name] = true
		elementNames = append(elementNames, element.Name)
	}
	return elementNames
}

// 以下只用到了 chargeableObj  的 GetID,GetName
// todo: 这个函数设计得不太好...
func getExpenseInfo(chargeableObj model.Chargeable) *pb.ExpenseInfo {
	switch chargeableObj.(type) {
	case *model.Good:
		return getPbGood(chargeableObj.(*model.Good)).ExpenseInfo
	case *model.Desk:
		return getPbDesk(chargeableObj.(*model.Desk), false).ExpenseInfo
	default:
		panic("unfix type")
	}
}
