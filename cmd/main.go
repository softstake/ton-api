package main

import (
	"fmt"
	"github.com/tonradar/ton-api/config"
	pb "github.com/tonradar/ton-api/proto"
	"github.com/tonradar/ton-api/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	cfg := config.GetConfig("config")
	fmt.Println(cfg)

	server, err := server.NewTonApiServer(cfg)
	if err != nil {
		log.Fatalf("failed to init api: %v", err)
	}

	listener, err := net.Listen("tcp", ":5400")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	rpcserv := grpc.NewServer()

	pb.RegisterTonApiServer(rpcserv, pb.TonApiServer(server))
	reflection.Register(rpcserv)

	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}
