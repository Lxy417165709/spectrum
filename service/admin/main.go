package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"spectrum/service/admin/model"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Msg:"Running go http server success. :)",
	})
}

func main() {
	r := gin.Default()
	r.GET("/test", Test)
	if err := r.Run(fmt.Sprintf(":%d", 9000)); err != nil {
		logs.Error("Running go http server failed. :|")
		return
	}
}
