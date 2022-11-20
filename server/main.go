package main

import (
	"fmt"
	grpc_simple "hasanustun/grpc-simple/calculator"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Calculator!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc_simple.Server{}

	grpcServer := grpc.NewServer()

	grpc_simple.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
