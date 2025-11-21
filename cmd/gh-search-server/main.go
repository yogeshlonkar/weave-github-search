package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	pb "weave-github-search/api/gh-search/v1"
	gh_client "weave-github-search/internal/gh-client"
	ghss "weave-github-search/internal/gh-search-service"
)

// main function sets up and starts the gRPC server for the GitHub Search Service.
// It listens on a specified port (default 50051) and handles graceful shutdown on termination signals.
// The server registers the GitHub Search Service implementation from the internal package.
func main() {
	grpc_port := "50051"
	if port := os.Getenv("GRPC_PORT"); port != "" {
		grpc_port = port
	}

	listner, err := net.Listen("tcp", ":"+grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	ghc := gh_client.NewClient(os.Getenv("GITHUB_TOKEN"))
	pb.RegisterGithubSearchServiceServer(server, &ghss.Server{Ghc: ghc})

	fmt.Println("gRPC Server starting on port " + grpc_port)

	go gracefulShutdown(server)

	if err := server.Serve(listner); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// gracefulShutdown handles graceful shutdown of the gRPC server,
// it waits for a termination signal and attempts to gracefully stop the server
// within a specified timeout period. If the server does not stop gracefully
// within the timeout, it forces a stop.
func gracefulShutdown(server *grpc.Server) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	log.Println("Initiating graceful shutdown...")
	timer := time.AfterFunc(10*time.Second, func() {
		log.Println("Server couldn't stop gracefully in time. Doing force stop.")
		server.Stop()
	})
	defer timer.Stop()
	server.GracefulStop()
	fmt.Println("Server stopped gracefully.")
}
