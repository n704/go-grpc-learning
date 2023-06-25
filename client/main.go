package main

import (
	"context"
	"log"

	pb "github.com/n704/go-grpc-learning/proto"
	"google.golang.org/grpc"
)

func main() {
	// create a connection to the server
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	// create a client instance
	c := pb.NewCalcualatorServiceClient(conn)

	// create an AddRequest instance
	req := &pb.AddRequest{
		A: 10,
		B: 20,
	}

	// call the Add method
	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Add RPC: %v", err)
	}

	log.Printf("Response from Add: %v", res.Result)
}
