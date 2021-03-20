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

func getExpenseInfo(chargeableObj model.Chargeable) *pb.ExpenseInfo {
	switch chargeableObj.(type) {
	case *model.Good:
		return getPbGood(chargeableObj.(*model.Good), "todo").ExpenseInfo
	case *model.Desk:
		return getPbDesk(chargeableObj.(*model.Desk)).ExpenseInfo
	default:
		panic("unfix type")
	}
}
