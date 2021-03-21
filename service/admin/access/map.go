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
		"GetAllGoodClasses": {
			ReqFunc: func() interface{} {
				return &pb.GetAllGoodClassesReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetAllGoodClassesRes{}
			},
			SuccessMsg: "获取商品类成功",
			FailMsg:    "获取商品类失败",
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
		"AddGoodClass": {
			ReqFunc: func() interface{} {
				return &pb.AddGoodClassReq{}
			},
			ResFunc: func() interface{} {
				return &pb.AddGoodClassRes{}
			},
			SuccessMsg: "添加商品类成功",
			FailMsg:    "添加商品类失败",
		},
		"GetAllGoods": {
			ReqFunc: func() interface{} {
				return &pb.GetAllGoodsReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetAllGoodsRes{}
			},
			SuccessMsg: "获取商品成功",
			FailMsg:    "获取商品失败",
		},
		"AddElement": {
			ReqFunc: func() interface{} {
				return &pb.AddElementReq{}
			},
			ResFunc: func() interface{} {
				return &pb.AddElementRes{}
			},
			SuccessMsg: "添加成功",
			FailMsg:    "添加失败",
		},
		"GetAllGoodOptions": {
			ReqFunc: func() interface{} {
				return &pb.GetAllGoodOptionsReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetAllGoodOptionsRes{}
			},
			SuccessMsg: "获取成功",
			FailMsg:    "获取失败",
		},
		"GetAllDeskClasses": {
			ReqFunc: func() interface{} {
				return &pb.GetAllDeskClassesReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetAllDeskClassesRes{}
			},
			SuccessMsg: "获取成功",
			FailMsg:    "获取失败",
		},
		"GetAllDesks": {
			ReqFunc: func() interface{} {
				return &pb.GetAllDesksReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetAllDesksRes{}
			},
			SuccessMsg: "获取成功",
			FailMsg:    "获取失败",
		},
		"AddDeskClass": {
			ReqFunc: func() interface{} {
				return &pb.AddDeskClassReq{}
			},
			ResFunc: func() interface{} {
				return &pb.AddDeskClassRes{}
			},
			SuccessMsg: "添加成功",
			FailMsg:    "添加失败",
		},
		"AddDesk": {
			ReqFunc: func() interface{} {
				return &pb.AddDeskReq{}
			},
			ResFunc: func() interface{} {
				return &pb.AddDeskRes{}
			},
			SuccessMsg: "添加成功",
			FailMsg:    "添加失败",
		},
		"OrderDesk": {
			ReqFunc: func() interface{} {
				return &pb.OrderDeskReq{}
			},
			ResFunc: func() interface{} {
				return &pb.OrderDeskRes{}
			},
			SuccessMsg: "点桌成功",
			FailMsg:    "点桌失败",
		},
		"OrderGood": {
			ReqFunc: func() interface{} {
				return &pb.OrderGoodReq{}
			},
			ResFunc: func() interface{} {
				return &pb.OrderGoodRes{}
			},
			SuccessMsg: "点单成功",
			FailMsg:    "点单失败",
		},
		"DeleteElementSizeInfo": {
			ReqFunc: func() interface{} {
				return &pb.DeleteElementSizeInfoReq{}
			},
			ResFunc: func() interface{} {
				return &pb.DeleteElementSizeInfoRes{}
			},
			SuccessMsg: "删除成功",
			FailMsg:    "删除失败",
		},
		"GetOrder": {
			ReqFunc: func() interface{} {
				return &pb.GetOrderReq{}
			},
			ResFunc: func() interface{} {
				return &pb.GetOrderRes{}
			},
			SuccessMsg: "获取成功",
			FailMsg:    "获取失败",
		},
	},
	//"AddOptionClass": {
	//	ReqFunc: func() interface{} {
	//		return &pb.AddOptionClassReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.AddOptionClassRes{}
	//	},
	//	SuccessMsg: "添加选项类成功",
	//	FailMsg:    "添加选项类失败",
	//},
	//"AddGood": {
	//	ReqFunc: func() interface{} {
	//		return &pb.AddGoodReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.AddGoodRes{}
	//	},
	//	SuccessMsg: "添加商品成功",
	//	FailMsg:    "添加商品失败",
	//},
	//"GetAllOptionClasses": {
	//	ReqFunc: func() interface{} {
	//		return &pb.GetAllOptionClassesReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.GetAllOptionClassesRes{}
	//	},
	//	SuccessMsg: "获取所有选项类成功",
	//	FailMsg:    "获取所有选项类失败",
	//},
	//"DelOption": {
	//	ReqFunc: func() interface{} {
	//		return &pb.DelOptionReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.DelOptionRes{}
	//	},
	//	SuccessMsg: "删除选项成功",
	//	FailMsg:    "删除选项失败",
	//},
	//"GetAllGoods": {
	//	ReqFunc: func() interface{} {
	//		return &pb.GetAllGoodsReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.GetAllGoodsRes{}
	//	},
	//	SuccessMsg: "获取所有商品成功",
	//	FailMsg:    "获取所有商品失败",
	//},
	//"DelOptionClass": {
	//	ReqFunc: func() interface{} {
	//		return &pb.DelOptionClassReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.DelOptionClassRes{}
	//	},
	//	SuccessMsg: "删除选项类成功",
	//	FailMsg:    "删除选项类失败",
	//},
	//"GetClasses": {
	//	ReqFunc: func() interface{} {
	//		return &pb.GetAllGoodClassesReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.GetAllGoodClassesRes{}
	//	},
	//	SuccessMsg: "获得所有商品类成功",
	//	FailMsg:    "获得所有商品类失败",
	//},
	//"AddGoodClass": {
	//	ReqFunc: func() interface{} {
	//		return &pb.AddGoodClassReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.AddGoodClassRes{}
	//	},
	//	SuccessMsg: "添加商品类成功",
	//	FailMsg:    "添加商品类失败",
	//},
	//"DelGoodClass": {
	//	ReqFunc: func() interface{} {
	//		return &pb.DelGoodClassReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.DelGoodClassRes{}
	//	},
	//	SuccessMsg: "删除商品类成功",
	//	FailMsg:    "删除商品类失败",
	//},
	//"Order": {
	//	ReqFunc: func() interface{} {
	//		return &pb.OrderReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.OrderRes{}
	//	},
	//	SuccessMsg: "订单创建成功",
	//	FailMsg:    "订单创建失败",
	//},
	//"GetOrderLog": {
	//	ReqFunc: func() interface{} {
	//		return &pb.GetOrderLogReq{}
	//	},
	//	ResFunc: func() interface{} {
	//		return &pb.GetOrderLogRes{}
	//	},
	//	SuccessMsg: "获取订单信息成功",
	//	FailMsg:    "获取订单信息失败",
	//},

}
