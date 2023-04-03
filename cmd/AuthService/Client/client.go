package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	port     = 50051
	certPath = "./certificate/server.crt"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	log.Println(conn)
}
