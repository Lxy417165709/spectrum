package main

import (
	"flag"
	"fmt"
	"spectrum/common/env"
	"spectrum/common/utils"
	"spectrum/service/mvp/controller"
	"spectrum/service/mvp/dao"
)

var confEnv = flag.String("env", "local", "")

func main() {
	flag.Parse()
	confFilePath := fmt.Sprintf("configure/%s.json", *confEnv)
	utils.InitLogger()
	utils.InitConfigure(confFilePath)
	dao.InitMainDB()
	utils.InitRpc(env.Conf.Rpc.MvpServerPort, &controller.MvpServer{})
}
