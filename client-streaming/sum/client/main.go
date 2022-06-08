package main

import (
	"context"
	"log"
	"sum/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Create grpc server
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create client
	c := proto.NewSumServiceClient(conn)

	// Create stream to send numbers
	stream, err := c.Sum(context.Background())

	// Send request to grpc server
	for i := 0; i < 10; i++ {
		log.Printf("Sending number %d", i)
		numberRequest := proto.NumberRequest{
			Num: int64(i),
		}

		stream.Send(&numberRequest)

		time.Sleep(time.Second)
	}

	// Close the stream and receive the result
  res, err := stream.CloseAndRecv()
  if err != nil {
    log.Fatalf("Error while receiving response: %v", err)
  }

  log.Printf("Sum result: %d", res.Result)
}
