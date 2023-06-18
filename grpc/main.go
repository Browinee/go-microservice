package main

import (
	"grpc/handler"
	"grpc/initialize"
	"grpc/proto"
	"net"

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
  lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic("fail to listen"+err.Error())
	}
  grpc_health_v1.RegisterHealthServer(server, health.NewServer() )
	err  = server.Serve(lis)
	if err != nil {
		panic("fail to start grpc"+err.Error())
	}
}

