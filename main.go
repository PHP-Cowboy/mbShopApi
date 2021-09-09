package main

import (
	"fmt"
	"go.uber.org/zap"
	"mbShopApi/initialize"
)

func main() {
	port := 9000
	//初始化logger
	initialize.InitLogger()
	//初始化router
	r := initialize.InitRouter()

	zap.S().Infof("启动服务器，端口：%d", port)

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
