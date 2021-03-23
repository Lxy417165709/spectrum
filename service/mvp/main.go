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
	// 1. 解析传入参数
	flag.Parse()

	// 2. 初始化日志打印机
	utils.InitLogger()

	// 3. 初始化配置
	utils.InitConfigure(fmt.Sprintf("configure/%s.json", *confEnv))

	// 4. 初始化数据库
	dao.InitMainDB()

	// 5. 初始化RPC服务
	utils.InitRpc(env.Conf.Rpc.MvpServerPort, &controller.MvpServer{})
}
