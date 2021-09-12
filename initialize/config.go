package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"mbShopApi/global"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfig() {
	data := GetEnvInfo("ENV")
	var configFileName string
	configFileNamePrefix := "config"
	configFileName = fmt.Sprintf("./%s-pro.yaml", configFileNamePrefix)
	if data == "local" {
		configFileName = fmt.Sprintf("./%s-debug.yaml", configFileNamePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	fmt.Println(global.ServerConfig)

	//go func() {
	//	v.WatchConfig()
	//	v.OnConfigChange(func(in fsnotify.Event) {
	//		zap.S().Info("配置修改：", in.Name)
	//		err = v.ReadInConfig()
	//		if err != nil {
	//			panic(err)
	//		}
	//		err = v.Unmarshal(global.ServerConfig)
	//		if err != nil {
	//			panic(err)
	//		}
	//		zap.S().Info(global.ServerConfig)
	//	})
	//}()
	//
	//time.Sleep(time.Second * 3)
}
