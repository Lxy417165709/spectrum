package controller

import (
	"spectrum/common/ers"
	"spectrum/common/pb"
)

func CheckAddGoodParameter(req *pb.AddGoodReq) (*pb.Good, string, error) {
	// 1. 判断 good
	good := req.Good
	if good == nil {
		return nil, "", ers.New("商品为空。")
	}
	if good.MainElement == nil {
		return nil, "", ers.New("商品没有主元素。")
	}
	if len(good.MainElement.SizeInfos) == 0 {
		return nil, "", ers.New("商品没有默认选项。")
	}
	if good.MainElement.SelectedIndex < 0 {
		return nil, "", ers.New("商品默认选项索引非法，不能小于0。")
	}
	if good.MainElement.SelectedIndex >= int32(len(good.MainElement.SizeInfos)) {
		return nil, "", ers.New("商品默认选项索引非法，不能超过主元素可选尺寸数组的最大索引。")
	}

	// 2. 判断 className
	className := req.ClassName
	if errResult := CheckIsValidGoodClassName(className); errResult != nil {
		return nil, "", errResult
	}

	// 3. 返回
	return good, req.ClassName, nil
}

func CheckAddGoodClassParameter(req *pb.AddGoodClassReq) (*pb.GoodClass, error) {
	goodClass := req.GoodClass
	if errResult := CheckIsValidGoodClassName(goodClass.Name); errResult != nil {
		return nil, errResult
	}
	if errResult := CheckIsValidPictureStorePath(goodClass.PictureStorePath); errResult != nil {
		return nil, errResult
	}
	return goodClass, nil
}

func CheckIsValidPictureStorePath(path string) error {
	if path == "" {
		return ers.New("图片路径为空。")
	}
	return nil
}

func CheckIsValidGoodClassName(goodClassName string) error {
	if goodClassName == "" {
		return ers.New("商品类名为空。")
	}
	return nil
}
