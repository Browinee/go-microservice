package initialize

import (
	"fmt"
	"go-api/global"
	"go-api/proto"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)


func InitSrcConn() {
		// get grpc service from consul
		cfg := api.DefaultConfig()
		consulInfo := global.ServerConfig.Consul
		cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

		client, err := api.NewClient(cfg)

		if err != nil {
			panic(err)
		}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserServiceInfo.Name))
	if err != nil {
		panic(err)
	}
	userSrvHost := ""
	userSrvPort := 0
	for _, value := range data{
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == ""{
		zap.S().Fatal("[InitSrcConn] fail to connect UserSrvClient")
	}

		userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
		if err != nil {
			zap.S().Errorw("[GetUserList] fail", "msg", err.Error())
		}
		global.UserSrvClient = proto.NewUserClient(userConn)
}