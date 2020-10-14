package main

import (
	"test/common/env"
	"test/common/utils"
	"test/service/auth/controller"
	"test/service/auth/dao"
)

const confFilePath = "configure/alpha.json"

func main() {
	utils.InitLogger()
	utils.InitConfigure(confFilePath)
	dao.InitMainDB()
	utils.InitRpc(env.Conf.Rpc.AuthServerPort, &controller.AuthServer{})
}
