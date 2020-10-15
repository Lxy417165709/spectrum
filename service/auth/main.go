package main

import (
	"spectrum/common/env"
	"spectrum/common/utils"
	"spectrum/service/auth/controller"
	"spectrum/service/auth/dao"
)

const confFilePath = "configure/alpha.json"

func main() {
	utils.InitLogger()
	utils.InitConfigure(confFilePath)
	dao.InitMainDB()
	utils.InitRpc(env.Conf.Rpc.AuthServerPort, &controller.AuthServer{})
}
