package main

import (
	"io"
	"log"
	"net"
	"sum/proto"

	"google.golang.org/grpc"
)

// GrpcServer is a gRPC server
type GrpcServer struct {
	proto.UnimplementedSumServiceServer
}

// Sum implements the Sum RPC
func (s *GrpcServer) Sum(sumServer proto.SumService_SumServer) error {
	var sum int64
	for {
		req, err := sumServer.Recv()
		if err == io.EOF {
			// Client close the stream
			break
		}

		if err != nil {
			log.Printf("Error while receiving: %v", err)
			return err
		}

    log.Printf("Received: %v", req)
		sum += req.GetNum()
	}

	return sumServer.SendAndClose(&proto.SumResponse{
		Result: sum,
	})
}

func main() {
	// Create listener on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create grpc server to serve the client requests
	grpcServer := grpc.NewServer()
	proto.RegisterSumServiceServer(grpcServer, &GrpcServer{})

	log.Printf("Starting server at %v", l.Addr())

	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
