package server

import (
	pb "github.com/jamesvincentsiauw-mac/grpc-gateway-demo/proto/gen/api"
	"github.com/jamesvincentsiauw-mac/grpc-gateway-demo/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

var helloService = service.HelloService{}
var operationService = service.OperationService{}

func StartGrpcServer() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Listen gRPC port failed: ", err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, &helloService)
	pb.RegisterOperationServiceServer(server, &operationService)

	log.Println("Start gRPC Server on 0.0.0.0:9090")
	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("Start gRPC Server failed: ", err)
	}
}
