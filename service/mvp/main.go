package main

import (
	"spectrum/common/env"
	"spectrum/common/utils"
	"spectrum/service/mvp/controller"
	"spectrum/service/mvp/dao"
)

const confFilePath = "/Users/a123/Desktop/lxy_test/go/grpc/spectrum/configure/alpha.json"

func main() {
	utils.InitLogger()
	utils.InitConfigure(confFilePath)
	dao.InitMainDB()
	utils.InitRpc(env.Conf.Rpc.MvpServerPort, &controller.MvpServer{})
}
