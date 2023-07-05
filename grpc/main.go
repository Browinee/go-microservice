package main

import (
	"flag"
	"fmt"
	"grpc/global"
	"grpc/handler"
	"grpc/initialize"
	"grpc/proto"
	"grpc/utils"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)




func main(){
  IP := flag.String("ip", "0.0.0.0","ip address")
	Port := flag.Int("port", 0, "port")
	flag.Parse()
	if *Port == 0 {
		*Port , _ = utils.GetFreePort()
	}

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitMysql()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	zap.S().Infof("Port", *Port)
	zap.S().Infof("IP", *IP)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
  // lis, err := net.Listen("tcp", "192.168.0.2:50051")
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
		GRPC:  fmt.Sprintf("192.168.1.103:%d", *Port),
		// GRPC: "192.168.0.2:50051",
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "600s",
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	uuid := uuid.New()

	// NOTE if id the same, it will override service in consul
	serviceId := uuid.String()
	registration.ID = serviceId
	registration.Port = *Port /* 50051 */
	registration.Tags = []string{"user", "service"}
	registration.Address = "192.168.1.103"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic("fail to register grpc"+err.Error())
	}


	go func() {
		err  = server.Serve(lis)
		if err != nil {
			panic("fail to start grpc"+err.Error())
		}
	}()
	// graceful quit
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent() .ServiceDeregister(serviceId);err != nil {
	zap.S().Info("Fail to deregister service")
}
zap.S().Infof("Deregister service: %s", serviceId)
}

