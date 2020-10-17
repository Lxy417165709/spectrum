package access

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"net/http"
	"reflect"
	"spectrum/common/ers"
	"spectrum/common/logger"
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
	},
}

func DistributeRequest(c *gin.Context) {

	// 1. 解析请求
	var request model.Request
	logger.Info("Success to get request")
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Fail to get request's data", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	logger.Info("Success to get request's data", zap.Any("data", request))

	// 2. 获取 Rpc 单元
	rpcUnit := objectFunctionToRpcUnit[request.Object][request.Function]
	if rpcUnit == nil {
		logger.Error("Fail to get object",
			zap.Any("requestObjectName", request.Object))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: "无法获得 RPC 单元，请联系 小悦悦 解决",
		})
		return
	}

	// 3. 根据请求，调用相应函数
	var req, res = rpcUnit.ReqFunc(), rpcUnit.ResFunc()
	if err := mapstructure.Decode(request.Parameters, req); err != nil {
		logger.Error("Fail to decode parameters",
			zap.Any("parameters", request.Parameters),
			zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	if objectMap[request.Object] == nil {
		logger.Error("Fail to get object",
			zap.Any("requestObjectName", request.Object))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: "无法获得 RPC 对象，请联系 小悦悦 解决",
		})
		return
	}
	object := reflect.ValueOf(objectMap[request.Object])
	method := object.MethodByName(request.Function)
	returnValues := method.Call([]reflect.Value{
		reflect.ValueOf(context.Background()),
		reflect.ValueOf(req),
	})
	// 3. 解析返回值
	if err := mapstructure.Decode(returnValues[0].Interface(), res); err != nil {
		logger.Error("Fail to decode returnValues[0]",
			zap.Any("returnValues[0]", returnValues[0]),
			zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}

	// 4. 解析错误
	var returnErr ers.Error
	hasError := returnValues[1].Interface() != nil
	if hasError {
		if err := json.Unmarshal(
			[]byte(status.Convert(returnValues[1].Interface().(error)).Message()),
			&returnErr,
		); err != nil {
			logger.Error("Fail to unmarshal returnValues[1] to ers.Error",
				zap.Any("returnValues[1]", returnValues[1]),
				zap.Error(err))
			c.JSON(http.StatusBadRequest, model.Response{
				Err: err.Error(),
			})
			return
		}
	}

	// 4. 返回结果
	if hasError {
		c.JSON(http.StatusOK, model.Response{
			Msg: rpcUnit.FailMsg,
			Err: returnErr.Detail,
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			Msg:  rpcUnit.SuccessMsg,
			Data: res,
		})
	}
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Msg: "Running go http server success. :)",
	})
}
