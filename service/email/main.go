package main

import (
	"test/common/env"
	"test/common/utils"
	"test/service/email/controller"
)

const confFilePath = "configure/alpha.json"

func main() {
	utils.InitLogger()
	utils.InitConfigure(confFilePath)
	utils.InitRpc(env.Conf.Rpc.EmailServerPort, &controller.EmailServer{})
}
