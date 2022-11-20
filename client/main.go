package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpc_simple "hasanustun/grpc-simple/calculator"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := grpc_simple.NewChatServiceClient(conn)

	response, err := c.Add(context.Background(), &grpc_simple.CalcRequest{Num1: 2, Num2: 3})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %f", response.Res)

	responseSub, err := c.Subtract(context.Background(), &grpc_simple.CalcRequest{Num1: 2, Num2: 3})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %f", responseSub.Res)

	responseMul, err := c.Multiply(context.Background(), &grpc_simple.CalcRequest{Num1: 2, Num2: 3})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %f", responseMul.Res)

	responseDiv, err := c.Divide(context.Background(), &grpc_simple.CalcRequest{Num1: 2, Num2: 3})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %f", responseDiv.Res)

}
