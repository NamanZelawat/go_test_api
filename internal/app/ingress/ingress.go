package ingress

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/NamanZelawat/go_test_api/proto/image"
	"google.golang.org/grpc"
)

func init() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(s, &server{})
	// Serve gRPC server

	log.Println("Serving gRPC on 0.0.0.0:8080")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Serving gRPC on 0.0.0.0:8080")
}

type server struct {
	pb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	filetype := http.DetectContentType(in.InputField)
	fmt.Println(filetype)
	return &pb.HelloReply{Message: "I am the world"}, nil
}
