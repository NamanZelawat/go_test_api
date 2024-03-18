package server

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/NamanZelawat/go_test_api/proto/image"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	kafka "github.com/segmentio/kafka-go"
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

	fo, err := os.Create("output.png")
	if err != nil {
		panic(err)
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

	s3Client, err := minio.New("minio:9000", &minio.Options{
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

	// to produce messages
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	log.Println("Successfully sent message")
	return nil
}
