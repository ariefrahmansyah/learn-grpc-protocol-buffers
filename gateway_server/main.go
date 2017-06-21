package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/ariefrahmansyah/learn-grpc/gateway"
)

const (
	port = ":8080"
)

// server is used to implement gateway.MakePayment.
type server struct{}

// MakePayment implements gateway.MakePayment
func (s *server) MakePayment(ctx context.Context, in *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Println("Receive request:", in)
	return &pb.PaymentResponse{Status: 200, Message: "Payment Success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("server is listening at port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
