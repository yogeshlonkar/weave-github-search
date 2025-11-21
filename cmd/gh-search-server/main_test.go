package main

import (
	"log"
	"os"
	"time"
)

func Example_main() {
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
	// Server stopped gracefully.
}

func Example_main_withPortEnv() {
	os.Setenv("GRPC_PORT", "60051")
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
	// Server stopped gracefully.
}
