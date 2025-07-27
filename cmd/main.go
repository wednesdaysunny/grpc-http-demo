package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "grpc-gateway-demo/proto/user"
	"grpc-gateway-demo/server"
)

const (
	grpcPort = "9091"
	httpPort = "9090"
)

func main() {
	log.Println("Starting gRPC-Gateway demo server...")

	userServer := server.NewUserServer()

	go startGRPCServer(userServer)

	time.Sleep(1 * time.Second)

	go startHTTPServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")
}

func startGRPCServer(userServer *server.UserServer) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", grpcPort, err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userServer)

	reflection.Register(s)

	log.Printf("gRPC server listening on :%s", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func startHTTPServer() {
	mux, err := server.NewGatewayServer(grpcPort)
	if err != nil {
		log.Fatalf("Failed to create gateway server: %v", err)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)

	httpMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "ok"}`)
	})

	log.Printf("HTTP server listening on :%s", httpPort)
	log.Printf("Swagger UI available at http://localhost:%s/swagger/", httpPort)
	log.Printf("Health check available at http://localhost:%s/health", httpPort)

	if err := http.ListenAndServe(":"+httpPort, httpMux); err != nil {
		log.Fatalf("Failed to serve HTTP server: %v", err)
	}
}
