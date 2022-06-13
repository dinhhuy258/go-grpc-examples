package main

import (
	"feed/proto"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GrpcServer is a gRPC server that implements FeedServiceServer.
type GrpcServer struct {
	proto.UnimplementedFeedServiceServer
}

// Feed receives a stream of FeedMessages and echoes them back.
func (s *GrpcServer) Feed(stream proto.FeedService_FeedServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("failed to receive request: %v", err)
			return err
		}

		log.Printf("Received request: %v", req.Feed)
		stream.Send(&proto.FeedResponse{Feed: fmt.Sprintf("Server send response: %v", req.Feed)})
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterFeedServiceServer(s, &GrpcServer{})

	log.Printf("Starting server at %v", l.Addr())

	if err = s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
