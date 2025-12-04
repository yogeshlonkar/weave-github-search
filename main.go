package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	pb "weave-github-search/api/gh-search/v1"
	"weave-github-search/internal/github"
	"weave-github-search/internal/services"
)

// main function sets up and starts the gRPC server for the GitHub Search Service.
// It listens on a specified port (default 50051) and handles graceful shutdown on termination signals.
// The server registers the GitHub Search Service implementation from the internal package.
func main() {
	grpcPort := "50051"
	if port := os.Getenv("GRPC_PORT"); port != "" {
		grpcPort = port
	}

	listener, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatal("GITHUB_TOKEN environment variable is not set")
	}

	ghc, err := github.NewClient(ctx, githubToken)
	if err != nil {
		log.Fatalf("failed to create GitHub client: %v", err)
	}
	pb.RegisterGithubSearchServiceServer(server, &services.Server{GithubClient: ghc})

	fmt.Println("gRPC Server starting on port " + grpcPort)

	go gracefulShutdown(ctx, server)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// gracefulShutdown handles graceful shutdown of the gRPC server,
// it waits for a termination signal and attempts to gracefully stop the server
// within a specified timeout period. If the server does not stop gracefully
// within the timeout, it forces a stop.
func gracefulShutdown(ctx context.Context, server *grpc.Server) {
	<-ctx.Done()
	log.Println("initiating graceful shutdown...")
	timer := time.AfterFunc(10*time.Second, func() {
		log.Println("server couldn't stop gracefully in time. Doing force stop.")
		server.Stop()
	})
	defer timer.Stop()
	server.GracefulStop()
	fmt.Println("server stopped gracefully.")
}
