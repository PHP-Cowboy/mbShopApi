package main

import (
	"fmt"
	"go.uber.org/zap"
	"mbShopApi/global"
	"mbShopApi/initialize"
)

func main() {
	//初始化logger
	initialize.InitLogger()
	//初始化配置
	initialize.InitConfig()

	//初始化router
	r := initialize.InitRouter()

	port := global.ServerConfig.Port

	fmt.Println()

	zap.S().Infof("启动服务器，端口：%d", port)

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
