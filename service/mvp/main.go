package main

import (
	"test/common/env"
	"test/common/utils"
	"test/service/mvp/controller"
	"test/service/mvp/dao"
)

const confFilePath = "configure/alpha.json"

func main() {
	utils.InitLogger()
	utils.InitConfigure(confFilePath)
	dao.InitMainDB()
	utils.InitRpc(env.Conf.Rpc.MvpServerPort, &controller.MvpServer{})
}
