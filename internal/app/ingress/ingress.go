package ingress

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/NamanZelawat/go_test_api/proto/image"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
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

	s.Serve(lis)
	// go func() {
	// 	log.Fatalln(s.Serve(lis))
	// }()

	// 	flag.Parse()
	// 	// Set up a connection to the server.
	// 	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 	if err != nil {
	// 		log.Fatalf("did not connect: %v", err)
	// 	}
	// 	defer conn.Close()
	// 	c := pb.NewGreeterClient(conn)
	//
	// 	// Contact the server and print out its response.
	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 	defer cancel()
	// 	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	// 	if err != nil {
	// 		log.Fatalf("could not greet: %v", err)
	// 	}
	// 	log.Printf("Greeting: %s", r.GetMessage())
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
