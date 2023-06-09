package initialize

import (
	"fmt"
	"go-api/global"
	"go-api/proto"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitSrcConn() {
	consulInfo := global.ServerConfig.Consul
	a := fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserConfig.Name)
	zap.S().Info("-------aaaaa: %+v ", a)
	userConn, err := grpc.Dial(
		a,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
 zap.S().Infof("consulInfo-------%+v", consulInfo)
	if err != nil {
		zap.S().Fatalf("[InitSrvConn] Fail to connect UserSrvClient %+v:", err.Error())
	}
	userSrcClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrcClient


}
func InitSrcConn_Deprecated() {
		// get grpc service from consul
		cfg := api.DefaultConfig()
		consulInfo := global.ServerConfig.Consul
		cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

		client, err := api.NewClient(cfg)

		if err != nil {
			panic(err)
		}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserConfig.Name))
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