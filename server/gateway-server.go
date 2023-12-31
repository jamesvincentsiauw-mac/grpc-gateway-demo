package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/jamesvincentsiauw-mac/grpc-gateway-demo/proto/gen/api"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func StartGwServer() {
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:9090",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server: ", err)
	}

	mux := runtime.NewServeMux()
	err = pb.RegisterHelloServiceHandler(context.Background(), mux, conn)
	err = pb.RegisterOperationServiceHandler(context.Background(), mux, conn)

	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}

	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("Start gRPC Gateway Server on http://0.0.0.0:8090")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("Start Gateway Server failed: ", err)
	}
}
