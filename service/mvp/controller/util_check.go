package controller

import (
	"spectrum/common/ers"
	"spectrum/common/pb"
	"spectrum/service/mvp/model"
	"spectrum/service/mvp/utils"
	"strconv"
	"time"
)

// --------------------------------------------------- 请求参数校验层 ---------------------------------------------------
func checkAddGoodParameter(req *pb.AddGoodReq) (*pb.Good, string, error) {
	// 1. 判断 good
	good := req.Good
	if errResult := checkIsValidGood(good); errResult != nil {
		return nil, "", errResult
	}

	// 2. 判断 className
	className := req.ClassName
	if errResult := checkIsValidGoodClassName(className); errResult != nil {
		return nil, "", errResult
	}

	// 3. 返回
	return good, req.ClassName, nil
}

func checkAddGoodClassParameter(req *pb.AddGoodClassReq) (*pb.GoodClass, error) {
	goodClass := req.GoodClass
	if errResult := checkIsValidGoodClassName(goodClass.Name); errResult != nil {
		return nil, errResult
	}
	if errResult := checkIsValidPictureStorePath(goodClass.PictureStorePath); errResult != nil {
		return nil, errResult
	}
	return goodClass, nil
}

func checkOrderGoodParameter(req *pb.OrderGoodReq) (int64, []*pb.Good, error) {
	if req.OrderID <= 0 {
		return 0, nil, ers.New("订单ID 必须大于 0。当前订单ID为 %d。",req.OrderID)
	}
	for index, good := range req.Goods {
		if errResult := checkIsValidGood(good); errResult != nil {
			return 0, nil, ers.New("第 %d 个商品非法。%s", index+1, errResult.Error())
		}
	}
	return req.OrderID, req.Goods, nil
}

// --------------------------------------------------- 结构校验层 ---------------------------------------------------
func checkIsValidGood(good *pb.Good) error {
	// 1. 判断商品整体
	if good == nil {
		return ers.New("商品为空。")
	}

	// 2. 判断商品ID
	if good.Id < 0 {
		return ers.New("商品ID 必须大于等于 0。")
	}

	// 3. 判断商品主元素
	if errResult := checkIsValidElement(good.MainElement); errResult != nil {
		return ers.New("商品主元素非法，%s", errResult.Error())
	}

	// 4. 判断商品附属元素
	for _, attachElement := range good.AttachElements {
		if errResult := checkIsValidElement(attachElement); errResult != nil {
			return ers.New("商品第 %d 个附属元素非法。%s", errResult.Error())
		}
	}

	// 5. 判断商品花费信息
	if errResult := checkIsValidExpenseInfo(good.ExpenseInfo); errResult != nil {
		return errResult
	}

	// 6. 判断商品优惠
	for _, favor := range good.Favors {
		if errResult := checkIsValidFavor(favor); errResult != nil {
			return ers.New("商品第 %d 个优惠非法。%s", errResult.Error())
		}
	}

	// 7. 返回
	return nil
}

func checkIsValidFavor(favor *pb.Favor) error {
	if favor == nil {
		return ers.New("优惠结构 不能为空。")
	}
	if !isValidFavorType(favor.FavorType) {
		return ers.New("优惠类型 只能为 %s。", utils.GetValidFavorTypesString())
	}
	if _, errResult := model.GetFavor(favor); errResult != nil {
		return errResult
	}
	return nil
}

func checkIsValidExpenseInfo(info *pb.ExpenseInfo) error {
	if info == nil {
		return ers.New("花费信息 为空。")
	}
	if info.CheckOutAt > time.Now().Unix() {
		return ers.New("结账时间 不能大于当前时间。")
	}
	if info.CheckOutAt == 0 {
		return ers.New("结账时间 不能为空。")
	}
	if info.Expense < 0 {
		return ers.New("实际花费 不能小于 0。")
	}
	if info.NonFavorExpense < 0 {
		return ers.New("未优惠花费 不能小于 0。")
	}
	if info.Expense > info.NonFavorExpense {
		return ers.New("实际花费 不能大于 未优惠花费。")
	}
	return nil
}

func checkIsValidElement(element *pb.Element) error {
	if element == nil {
		return ers.New("元素 为空。")
	}
	if element.Id < 0 {
		return ers.New("元素ID 必须大于等于 0。")
	}
	if element.Name == "" {
		return ers.New("元素名 不能为空。")
	}
	if !isValidElementType(element.Type) {
		return ers.New("元素类型 只能为 %s。", utils.GetValidElementTypesString())
	}
	if errResult := checkIsValidSelectedIndex(int(element.SelectedIndex), len(element.SizeInfos)); errResult != nil {
		return ers.New("元素尺寸选择索引非法，%s", errResult.Error())
	}
	if len(element.SizeInfos) == 0 {
		return ers.New("元素 没有规格选项。")
	}
	for index, sizeInfo := range element.SizeInfos {
		if errResult := checkIsValidSizeInfo(sizeInfo); errResult != nil {
			return ers.New("第 %d 个尺寸非法。%s", index+1, errResult.Error())
		}
	}
	return nil
}

func checkIsValidSizeInfo(sizeInfo *pb.SizeInfo) error {
	if sizeInfo == nil {
		return ers.New("尺寸 为空。")
	}
	if sizeInfo.Id < 0 {
		return ers.New("尺寸ID 必须大于等于 0。")
	}
	if sizeInfo.Size == "" {
		return ers.New("尺寸名 为空。")
	}
	if errResult := checkIsValidPictureStorePath(sizeInfo.PictureStorePath); errResult != nil {
		if sizeInfo.PictureStorePath == "" {
			return ers.New("尺寸图片 非法, %s", errResult.Error())
		}
	}
	price, err := strconv.ParseFloat(sizeInfo.Price, 64)
	if err != nil {
		return ers.New("尺寸价格 只能为小数。")
	}
	if price < 0 {
		return ers.New("尺寸价格 必须大于等于 0。")
	}
	return nil
}

// --------------------------------------------------- 字段校验层 ---------------------------------------------------
func checkIsValidSelectedIndex(selectedIndex int, sizeInfoLength int) error {
	if selectedIndex < 0 {
		return ers.New("选项索引 不能小于0。")
	}
	if selectedIndex >= sizeInfoLength {
		return ers.New("选项索引 不能大于可选尺寸数组的最大索引。")
	}
	return nil
}

func checkIsValidPictureStorePath(path string) error {
	if path == "" {
		return ers.New("图片路径为空。")
	}
	return nil
}

func checkIsValidGoodClassName(goodClassName string) error {
	if goodClassName == "" {
		return ers.New("商品类名为空。")
	}
	return nil
}

func isValidElementType(elementType pb.ElementType) bool {
	for _, validElementType := range utils.ValidElementTypes {
		if elementType == validElementType {
			return true
		}
	}
	return false
}

func isValidFavorType(favorType pb.FavorType) bool {
	for _, validFavorType := range utils.ValidFavorTypes {
		if validFavorType == favorType {
			return true
		}
	}
	return false
}
