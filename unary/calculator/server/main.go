package main

import (
	"calculator/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GrpcServer is a gRPC server
type GrpcServer struct {
	proto.UnimplementedCalculatorServiceServer
}

// Sum is a function that adds two numbers
func (s *GrpcServer) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	result := req.Num1 + req.Num2

	response := &proto.SumResponse{
		Result: result,
	}

	return response, nil
}

// Sub is a function that subtracts two numbers
func (s *GrpcServer) Sub(
	ctx context.Context,
	req *proto.SubRequest,
) (*proto.SubResponse, error) {
	result := req.Num1 - req.Num2

	response := &proto.SubResponse{
		Result: result,
	}

	return response, nil
}

func main() {
  l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &GrpcServer{})

	log.Printf("Starting server at %v", l.Addr())

	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
