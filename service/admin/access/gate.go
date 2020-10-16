package access

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/common/rpc"
	"spectrum/service/admin/model"
)

var objectMap = map[string]interface{}{
	"mvp": rpc.MvpClient,
}

var objectFunctionToReq = map[string]map[string]interface{}{
	"mvp": {
		"AddOptionClass":      &pb.AddOptionClassReq{},
		"AddGood":             &pb.AddGoodReq{},
		"GetAllOptionClasses": &pb.GetAllOptionClassesReq{},
	},
}

var objectFunctionToRes = map[string]map[string]interface{}{
	"mvp": {
		"AddOptionClass":      &pb.AddOptionClassRes{},
		"AddGood":             &pb.AddGoodRes{},
		"GetAllOptionClasses": &pb.GetAllOptionClassesRes{},
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

	// 2. 根据请求，调用相应函数
	req := objectFunctionToReq[request.Object][request.Function]
	if err := mapstructure.Decode(request.Parameters, req); err != nil {
		logger.Error("Fail to decode parameters",
			zap.Any("parameters", request.Parameters),
			zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	if objectMap[request.Object] == nil{
		logger.Error("Fail to get object",
			zap.Any("requestObjectName", request.Object))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: "无法获得 RPC 对象",
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
	res := objectFunctionToRes[request.Object][request.Function]
	if err := mapstructure.Decode(returnValues[0].Interface(), res); err != nil {
		logger.Error("Fail to decode returnValues[0]",
			zap.Any("returnValues[0]", returnValues[0]),
			zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	var returnErr error
	if err := mapstructure.Decode(returnValues[1].Interface(), &returnErr); err != nil {
		logger.Error("Fail to decode returnValues[1]",
			zap.Any("returnValues[1]", returnValues[1]),
			zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}

	// 4. 返回结果
	if returnErr != nil{
		c.JSON(http.StatusOK, model.Response{
			Msg:  "Fail to request ;(",
			Err:  returnErr.Error(),
		})
	}else{
		c.JSON(http.StatusOK, model.Response{
			Msg:  "Success to request ;)",
			Data: res,
		})
	}
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Msg: "Running go http server success. :)",
	})
}
