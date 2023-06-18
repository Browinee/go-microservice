package main

import (
	"grpc/handler"
	"grpc/initialize"
	"grpc/proto"
	"net"

	"google.golang.org/grpc"
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

	err  = server.Serve(lis)
	if err != nil {
		panic("fail to start grpc"+err.Error())
	}
}

