package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/admin/access"
)

func main() {
	var port = 9000
	r := gin.Default()
	r.GET("/test", access.Test)
	r.POST("/distributor", access.DistributeRequest)
	logger.Info("Ready to run go http server", zap.Any("port", port))
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		logger.Error("Fail to run go http server",
			zap.Any("port", port),
			zap.Error(err))
		return
	}
}
