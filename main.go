package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mbShopApi/global"
	"mbShopApi/initialize"
	validator2 "mbShopApi/validator"
)

func main() {
	//初始化logger
	initialize.InitLogger()
	//初始化配置
	initialize.InitConfig()

	zap.S().Info(global.ServerConfig)

	//初始化router
	r := initialize.InitRouter()

	//初始化验证器
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		err := v.RegisterValidation("mobile", validator2.ValidateMobile)
		if err != nil {
			panic(err)
		}
	}

	port := global.ServerConfig.Port

	zap.S().Infof("启动服务器，端口：%d", port)

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
