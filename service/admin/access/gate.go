package gate

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"spectrum/common/logger"
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
	c.JSON(http.StatusOK, model.Response{
		Msg: "Success to request ;)",
	})
}
