package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const (
	port = 50051
)

func main() {
	server := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
