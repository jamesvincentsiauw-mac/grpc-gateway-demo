package main

import "github.com/jamesvincentsiauw-mac/grpc-gateway-demo/server"

func main() {
	go server.StartGrpcServer()
	server.StartGwServer()
}
