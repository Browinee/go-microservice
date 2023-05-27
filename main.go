package main

import (
	"go-microservice/handler"
	"go-microservice/proto"
	"net"

	"google.golang.org/grpc"
)
func main(){
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
  lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		panic("fail to listen"+err.Error())
	}

	err  = server.Serve(lis)
	if err != nil {
		panic("fail to start grpc"+err.Error())
	}
}