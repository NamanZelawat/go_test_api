package server

import (
	"context"
	"io"
	"log"
	"os"

	pb "github.com/NamanZelawat/go_test_api/proto/image"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// CreateServer creates a grpc sever serving on 8080
func CreateServer() *grpc.Server {
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(s, &server{})
	// Serve gRPC server

	log.Println("Starting to serve gRPC on 0.0.0.0:8080")

	return s
}

func (s *server) SayHello(stream pb.Greeter_SayHelloServer) error {
	// filetype := http.DetectContentType(in.InputField)
	// fmt.Println(filetype)

	minioURL := os.Getenv("MINIO_URL")
	messageURL := os.Getenv("MESSAGE_URL")

	fo, err := os.Create("output.png")
	if err != nil {
		log.Println("Error creating file", err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.HelloReply{Message: "I am the world"})
			break
		}

		if _, err := fo.Write(chunk.InputField); err != nil {
			panic(err)
		}
	}

	s3Client, err := minio.New(minioURL, &minio.Options{
		Creds:  credentials.NewStaticV4("root", "password", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := s3Client.FPutObject(context.Background(), "image", "test.png", "output.png", minio.PutObjectOptions{
		ContentType: "application/png",
	}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully uploaded image")

	nc, err := nats.Connect(messageURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Simple Publisher
	subject := "example"
	msg := []byte("Hello NATS!")
	nc.Publish(subject, msg)
	log.Printf("Published message: %s\n", msg)
	return nil
}
