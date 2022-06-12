package main

import (
	"context"
	"feed/proto"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	client := proto.NewFeedServiceClient(conn)
	stream, err := client.Feed(context.Background())
	if err != nil {
		log.Fatalf("failed to call Feed: %v", err)
	}

	waitChannel := make(chan struct{})

	// Go routine to send data to server
	go func() {
		for i := 0; i < 5; i++ {
			err := stream.Send(&proto.FeedRequest{
				Feed: "Request " + fmt.Sprint(i),
			})
			if err != nil {
				log.Fatalf("failed to send request: %v", err)
			}

			time.Sleep(time.Second)
		}

		// Close the write stream
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("failed to close stream: %v", err)
		}
	}()

	// Go routine to receive data from server
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitChannel)
				return
			}

			if err != nil {
				log.Fatalf("failed to receive response: %v", err)
			}

			println("Received response: " + resp.GetFeed())
		}
	}()

	<-waitChannel
}
