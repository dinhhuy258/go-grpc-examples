package main

import (
	"countdown/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

// GrpcServer is a gRPC server that implements the CountdownServiceServer interface.
type GrpcServer struct {
	proto.UnimplementedCountdownServiceServer
}

// Countdown implements the CountdownServiceServer interface.
func (s *GrpcServer) Countdown(req *proto.CountdownRequest, stream proto.CountdownService_CountdownServer) error {
	log.Printf("Countdown request: %v", req.Timer)

	for i := req.Timer; i > 0; i-- {
		stream.Send(&proto.CountdownResponse{
			Count: i,
		})

		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterCountdownServiceServer(grpcServer, &GrpcServer{})

	log.Printf("Starting server at %v", listener.Addr())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
