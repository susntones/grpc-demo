package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"baishi/grpc-demo/proto/demo"
	"baishi/grpc-demo/server"
)

func main() {
	listener, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Println("net listen err ", err)
		return
	}
	s := grpc.NewServer()
	demo.RegisterDemoServer(s, server.NewDemoServer())

	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}
