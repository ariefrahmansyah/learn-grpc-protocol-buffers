package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/ariefrahmansyah/learn-grpc/gateway"
)

const (
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGatewayClient(conn)

	// Make payment to the payment gateway and print out its response.
	request := &pb.PaymentRequest{
		PaymentId:         1,
		Merchant:          "tokopedia",
		MerchantPaymentId: "TKP1",
		Amount:            float64(100000),
		User: &pb.PaymentRequest_User{
			Id:    1,
			Name:  "Arief",
			Email: "ariefrahmansyah@hotmail.com",
		},
	}

	r, err := c.MakePayment(context.Background(), request)
	if err != nil {
		log.Fatalf("could not make payment: %v", err)
	}

	log.Printf("Make Payment Response: %s", r.String())
}
