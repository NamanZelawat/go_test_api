package ingress

import (
	"log"
	"net"

	"github.com/NamanZelawat/go_test_api/internal/pkg/server"
)

func init() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := server.CreateServer()

	log.Println("Starting to serve gRPC on 0.0.0.0:8080")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
