package main

import (
	"test/common/env"
	"test/common/utils"
	"test/service/auth/controller"
)


func main() {
	utils.InitLogger()
	utils.InitConfigure()
	utils.InitRpc(env.Conf.Rpc.AuthServerPort, &controller.AuthServer{})
}
