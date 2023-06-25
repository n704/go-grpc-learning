package main

import (
	"context"
	"log"
	"net"

	pb "github.com/n704/go-grpc-learning/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	pb.CalcualatorServiceServer
}

// implement the Add method
func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	a, b := in.A, in.B
	log.Printf("Received a: %d, b: %d", a, b)
	result := a + b
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	// listen to port 8080
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new gRPC server
	s := grpc.NewServer()

	// register the calcualator service
	pb.RegisterCalcualatorServiceServer(s, &server{})

	log.Println("Server is running on port: ", port)
	// serve the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
