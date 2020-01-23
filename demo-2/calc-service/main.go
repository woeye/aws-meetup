package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/woeye/aws-demo/calc-service/calc_service"
)

type server struct {
	calc_service.UnimplementedCalcServiceServer
}

func main() {
	s := server{}

	grpcServer := grpc.NewServer()
	calc_service.RegisterCalcServiceServer(grpcServer, &s)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Could not open listener on port 8000: %v\n", err)
	}
	log.Printf("Starting server on port: 8000")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}

// Multiply implements the gRPC method
func (s *server) Multiply(ctx context.Context, in *calc_service.MultiplyRequest) (*calc_service.MultiplyResponse, error) {
	return &calc_service.MultiplyResponse{
		Result: in.GetA() * in.GetB(),
	}, nil
}
