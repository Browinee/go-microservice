package main

import (
	"fmt"
	"grpc/global"
	"grpc/handler"
	"grpc/initialize"
	"grpc/proto"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)




func main(){

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitMysql()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
  lis, err := net.Listen("tcp", "192.168.0.2:50051")
	if err != nil {
		panic("fail to listen"+err.Error())
	}
  grpc_health_v1.RegisterHealthServer(server, health.NewServer() )

	// register service

  cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	check := &api.AgentServiceCheck{
		GRPC: "192.168.0.2:50051",
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "600s",
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = 50051
	registration.Tags = []string{"user", "service"}
	registration.Address = "192.168.0.2"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic("fail to register grpc"+err.Error())
	}
	err  = server.Serve(lis)
	if err != nil {
		panic("fail to start grpc"+err.Error())
	}
}

