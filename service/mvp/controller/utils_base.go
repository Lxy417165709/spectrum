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
func getExpenseInfo(chargeableObj model.Chargeable) *pb.ExpenseInfo {
	id := chargeableObj.GetID()
	favors := getFavors(chargeableObj)

	// todo: 这里有部分和 getGood、getDesk 冗余了
	switch chargeableObj.(type) {
	case *model.Good:
		mainElement := getMainElement(id, "")
		attachElements := getAttachElements(id, "")
		return chargeableObj.(*model.Good).GetExpenseInfo(mainElement, attachElements, favors)
	case *model.Desk:
		desk, err := dao.DeskDao.Get(id)
		if err != nil {
			// todo: log
			return nil
		}
		space, err := dao.SpaceDao.Get(desk.SpaceName, desk.SpaceNum)
		if err != nil {
			// todo: log
			return nil
		}
		return chargeableObj.(*model.Desk).GetExpenseInfo(space.Price, favors)
	default:
		panic("unfix type")
	}
}
