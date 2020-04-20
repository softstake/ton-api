package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tonradar/ton-api/config"
	pb "github.com/tonradar/ton-api/proto"
	"github.com/tonradar/ton-api/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.GetConfig()
	log.Println(cfg)

	server, err := server.NewTonApiServer(cfg)
	if err != nil {
		log.Fatalf("failed to init api: %v", err)
	}

	port := fmt.Sprintf(":%d", cfg.ListenPort)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	rpcserv := grpc.NewServer(grpc.UnaryInterceptor(server.ServerUnaryInterceptor))

	pb.RegisterTonApiServer(rpcserv, pb.TonApiServer(server))
	reflection.Register(rpcserv)

	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}
