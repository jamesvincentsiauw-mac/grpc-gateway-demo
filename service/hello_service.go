package service

import (
	"context"
	pb "github.com/jamesvincentsiauw-mac/grpc-gateway-demo/proto/gen/api"
)

type HelloService struct {
}

func (h *HelloService) Hello(_ context.Context, message *pb.HelloMessage) (*pb.HelloResponse, error) {
	helloMessage := "Hello " + message.GetMessage()
	response := pb.HelloResponse{Result: helloMessage}

	return &response, nil
}
