package main

import (
	"fmt"
	"grpcserver/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalcServer(s, services.NewCalcServer())
	reflection.Register(s)

	fmt.Println("gRPC server is starting on port 50051")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
