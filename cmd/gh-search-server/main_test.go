package main

import (
	"log"
	"os"
	"time"
)

func Example_main() {
	os.Setenv("GITHUB_TOKEN", "test_token")
	go main()
	time.Sleep(150 * time.Millisecond)
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Fatalf("failed to find process: %v", err)
	}
	p.Signal(os.Interrupt)
	time.Sleep(150 * time.Millisecond)
	// Output:
	// gRPC Server starting on port 50051
	// server stopped gracefully.
}

func Example_main_withPortEnv() {
	os.Setenv("GRPC_PORT", "60051")
	os.Setenv("GITHUB_TOKEN", "test_token")
	go main()
	time.Sleep(150 * time.Millisecond)
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Fatalf("failed to find process: %v", err)
	}
	p.Signal(os.Interrupt)
	time.Sleep(150 * time.Millisecond)
	// Output:
	// gRPC Server starting on port 60051
	// server stopped gracefully.
}
