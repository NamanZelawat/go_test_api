package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	pb "github.com/NamanZelawat/go_test_api/proto/image"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn *grpc.ClientConn
	err  error
)

func init() {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	gwmux := runtime.NewServeMux()

	log.Println("Starting to establish the grpc conn")
	conn, err = grpc.Dial(
		"ingress:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	defer conn.Close()

	log.Println("Dialing done")
	if err != nil {
		log.Println("Failed to dial server:", err)
	}

	// Register Greeter
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	reqRegister(conn, gwmux)

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	var wg sync.WaitGroup
	numberOfWorkers := 1
	wg.Add(numberOfWorkers)
	go serverThread(gwServer, &wg)
	log.Println("Server is running guys")
	wg.Wait()
}

func serverThread(gwServer *http.Server, wg *sync.WaitGroup) {
	log.Fatalln(gwServer.ListenAndServe())
	defer wg.Done()
}

func reqRegister(conn *grpc.ClientConn, gwmux *runtime.ServeMux) {
	gwmux.HandlePath("POST", "/image", handleFileUpload)
	gwmux.HandlePath("GET", "/image", handleGet)
}

func handleFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string) {
	log.Println("Calling the function")

	f, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get file: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer f.Close()

	log.Println("Making buffer")

	buff := make([]byte, 512)

	log.Println("Detecting file type")

	filetype := http.DetectContentType(buff)

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.SayHello(ctx)

	for {
		bytesRead, err := f.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading chunk from file: %v", err)
		}

		// Send the chunk to the server
		err = stream.Send(&pb.HelloRequest{InputField: buff[:bytesRead]})
		if err != nil {
			log.Fatalf("error sending chunk to server: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.GetMessage())

	fmt.Println(filetype)

	httpResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error, : %v", err)
	}
	w.Write(httpResp)
}

func handleGet(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Write([]byte("Hello brother!!"))
}
