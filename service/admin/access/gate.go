package access

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"net/http"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/common/rpc"
	"spectrum/service/admin/model"
)

func DistributeRequest(c *gin.Context) {
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
	//if _,err := rpc.MvpClient.AddGood(context.Background(), &pb.AddGoodReq{
	//	GoodName: "草莓益菌多",
	//	Price:    15,
	//	Type:     1,
	//}); err != nil {
	//	logger.Error("Fail to add good", zap.Error(err))
	//	c.JSON(http.StatusBadRequest, model.Response{
	//		Err: err.Error(),
	//	})
	//	return
	//}

	aocr := pb.AddOptionClassReq{}
	if err := mapstructure.Decode(request.Parameters, &aocr); err != nil {
		logger.Error("Fail to decode parameters",
			zap.Any("parameters", request.Parameters),
			zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	if _, err := rpc.MvpClient.AddOptionClass(context.Background(), &aocr); err != nil {
		logger.Error("Fail to add option class",
			zap.Any("AddOptionClassReq", aocr),
			zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Msg: "Success to request ;)",
	})
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Msg: "Running go http server success. :)",
	})
}
