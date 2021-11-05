package main

import (
	"log"
	"net"

	"github.com/DaviAraujoCC/gRPC-Golang/chat"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer list.Close()

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalln(err)
	}
}
