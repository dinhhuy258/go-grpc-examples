package main

import (
	"calculator/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	c := proto.NewCalculatorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sumResponse, err := c.Sum(ctx, &proto.SumRequest{
		Num1: 1,
		Num2: 2,
	})
	if err != nil {
		log.Fatalf("failed to call Sum: %v", err)
	}

	log.Printf("Sum result: %d", sumResponse.Result)

  subResponse, err := c.Sub(ctx, &proto.SubRequest{
		Num1: 1,
		Num2: 2,
	})
	if err != nil {
		log.Fatalf("failed to call Sub: %v", err)
	}

	log.Printf("Sub result: %d", subResponse.Result)
}
