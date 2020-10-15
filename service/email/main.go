package main

import (
	"spectrum/common/env"
	"spectrum/common/utils"
	"spectrum/service/email/controller"
)

const confFilePath = "configure/alpha.json"

func main() {
	utils.InitLogger()
	utils.InitConfigure(confFilePath)
	utils.InitRpc(env.Conf.Rpc.EmailServerPort, &controller.EmailServer{})
}
