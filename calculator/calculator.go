package grpc_simple

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) Add(ctx context.Context, in *CalcRequest) (*CalcResponse, error) {
	log.Printf("Receive message body from client: %f", in.Num1)
	return &CalcResponse{Res: in.Num1 + in.Num2}, nil
}

func (s *Server) Subtract(ctx context.Context, in *CalcRequest) (*CalcResponse, error) {
	log.Printf("Receive message body from client: %f", in.Num1)
	return &CalcResponse{Res: in.Num1 - in.Num2}, nil
}

func (s *Server) Multiply(ctx context.Context, in *CalcRequest) (*CalcResponse, error) {
	log.Printf("Receive message body from client: %f", in.Num1)
	return &CalcResponse{Res: in.Num1 * in.Num2}, nil
}

func (s *Server) Divide(ctx context.Context, in *CalcRequest) (*CalcResponse, error) {
	log.Printf("Receive message body from client: %f", in.Num1)
	return &CalcResponse{Res: in.Num1 / in.Num2}, nil
}
