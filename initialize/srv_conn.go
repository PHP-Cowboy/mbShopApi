package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mbShopApi/global"
	"mbShopApi/proto"
)

func InitUserClient() {
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}

	srvHost := ""
	srvPort := 0

	for _, service := range data {
		srvHost = service.Address
		srvPort = service.Port
		break
	}

	if srvHost == "" {
		zap.S().Fatal("[InitUserClient] 连接 [用户服务失败]")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", srvHost, srvPort), grpc.WithInsecure())

	if err != nil {
		zap.S().Panic("err:" + err.Error())
	}

	global.UserClient = proto.NewUserClient(conn)
}
