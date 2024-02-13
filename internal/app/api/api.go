package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	helloworldpb "github.com/NamanZelawat/go_test_api/proto/image"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	gwmux := runtime.NewServeMux()

	log.Println("Starting to establish the grpc conn")
	conn, err := grpc.Dial(
		"ingress:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	defer conn.Close()

	log.Println("Dialing done")
	if err != nil {
		log.Println("Failed to dial server:", err)
	}

	// Register Greeter
	err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
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
	handleFileUpload := func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		log.Println("Calling the function")

		f, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get file: %s", err.Error()), http.StatusBadRequest)
			return
		}
		defer f.Close()

		buff := make([]byte, 512)

		_, err = f.Read(buff)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get 'attachment': %s", err.Error()), http.StatusBadRequest)
			return
		}

		filetype := http.DetectContentType(buff)

		c := helloworldpb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := c.SayHello(ctx, &helloworldpb.HelloRequest{InputField: buff})
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

	gwmux.HandlePath("POST", "/image", handleFileUpload)
}
