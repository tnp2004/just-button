package main

import (
	"clientserver/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	calcClient := services.NewCalcClient(cc)
	calcService := services.NewCalcService(calcClient)

	calcService.Plus(2)

}
