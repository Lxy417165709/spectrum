package access

import (
	"spectrum/common/pb"
	"spectrum/common/rpc"
	"spectrum/service/admin/model"
)

var objectMap = map[string]interface{}{
	"mvp": rpc.MvpClient,
}

var objectFunctionToRpcUnit = map[string]map[string]*model.RpcUnit{
	"mvp": {
		"AddOptionClass": {
			ReqFunc: func() interface{} {
				return &pb.AddOptionClassReq{}
			},
			ResFunc: func() interface{} {
				return &pb.AddOptionClassRes{}
			},
			SuccessMsg: "添加选项类成功",
			FailMsg:    "添加选项类失败",
		},
		"AddGood": {
			ReqFunc: func() interface{} {
				return &pb.AddGoodReq{}
			},
			ResFunc: func() interface{} {
				return &pb.AddGoodRes{}
			},
			SuccessMsg: "添加商品成功",
			FailMsg:    "添加商品失败",
		},
		"GetAllOptionClasses": {
			ReqFunc: func() interface{} {
				return &pb.GetAllOptionClassesReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetAllOptionClassesRes{}
			},
			SuccessMsg: "获取所有选项类成功",
			FailMsg:    "获取所有选项类失败",
		},
		"DelOption": {
			ReqFunc: func() interface{} {
				return &pb.DelOptionReq{}
			},
			ResFunc: func() interface{} {
				return &pb.DelOptionRes{}
			},
			SuccessMsg: "删除选项成功",
			FailMsg:    "删除选项失败",
		},
		"GetAllGoods": {
			ReqFunc: func() interface{} {
				return &pb.GetAllGoodsReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetAllGoodsRes{}
			},
			SuccessMsg: "获取所有商品成功",
			FailMsg:    "获取所有商品失败",
		},
		"DelOptionClass":{
			ReqFunc: func() interface{} {
				return &pb.DelOptionClassReq{}
			},
			ResFunc: func() interface{} {
				return &pb.DelOptionClassRes{}
			},
			SuccessMsg: "删除选项类成功",
			FailMsg:    "删除选项类失败",
		},
		"GetAllGoodClasses":{
			ReqFunc: func() interface{} {
				return &pb.GetAllGoodClassesReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetAllGoodClassesRes{}
			},
			SuccessMsg: "获得所有商品类成功",
			FailMsg:    "获得所有商品类失败",
		},
		"AddGoodClass":{
			ReqFunc: func() interface{} {
				return &pb.AddGoodClassReq{}
			},
			ResFunc: func() interface{} {
				return &pb.AddGoodClassRes{}
			},
			SuccessMsg: "添加商品类成功",
			FailMsg:    "添加商品类失败",
		},
	},
}
