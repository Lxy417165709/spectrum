package main

import (
	"test/common/env"
	"test/common/utils"
	"test/service/email/controller"
)


func main() {
	utils.InitLogger()
	utils.InitConfigure()
	utils.InitRpc(env.Conf.Rpc.EmailServerPort,&controller.EmailServer{})
}
