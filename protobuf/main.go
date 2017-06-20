package main

import (
	"bytes"
	"log"
	"net/http"

	pb "github.com/ariefrahmansyah/learn-grpc/gateway"
	"github.com/golang/protobuf/proto"
)

func main() {
	payment := &pb.PaymentRequest{
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

	data, err := proto.Marshal(payment)
	if err != nil {
		log.Fatalln("Failed to encode payment data:", err)
	}

	log.Println(data)

	payment2 := &pb.PaymentRequest{}

	if err := proto.Unmarshal(data, payment2); err != nil {
		log.Fatalln("Failed to decode payment data:", err)
	}

	log.Println(payment2)

	if _, err := http.Post("http://httpbin.org/post", "", bytes.NewBuffer(data)); err != nil {
		log.Fatalln("Failed to send payment data:", err)
	}
}
