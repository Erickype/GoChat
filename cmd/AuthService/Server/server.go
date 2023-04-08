package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

const (
	port     = 50051
	certPath = "./certificate/server.crt"
	keyPath  = "./certificate/server.key"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile(certPath, keyPath)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(grpc.Creds(creds))
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
