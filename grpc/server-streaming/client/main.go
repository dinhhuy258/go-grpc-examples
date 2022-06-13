package main

import (
	"context"
	"countdown/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewCountdownServiceClient(conn)

	stream, err := client.Countdown(context.Background(), &proto.CountdownRequest{
		Timer: 10,
	})
	if err != nil {
		log.Fatalf("Error while starting countdown: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving message: %v", err)
		}

		println(msg.Count)
	}
}
