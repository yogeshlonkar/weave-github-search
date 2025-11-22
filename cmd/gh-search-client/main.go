package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "weave-github-search/api/gh-search/v1"
)

func main() {
	grpc_port := "50051"
	if port := os.Getenv("GRPC_PORT"); port != "" {
		grpc_port = port
	}
	serverAddr := "localhost:" + grpc_port
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGithubSearchServiceClient(conn)

	for {
		searchGitHub(c)
	}
}

// searchGitHub prompts the user for a search query and optional GitHub username,
// performs a search using the provided gRPC client, and prints the results.
// Example input: "some_code @username"
func searchGitHub(c pb.GithubSearchServiceClient) {
	fmt.Println("Enter search query with optional user (e.g. some_code @username)")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	var q string
	var user string
	for index, part := range strings.Split(input, " @") {
		if index == 0 {
			q = part
			continue
		}
		if part != "" {
			user = part
			break
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	q = strings.TrimSpace(q)
	user = strings.TrimSpace(user)
	resp, err := c.Search(ctx, &pb.SearchRequest{SearchTerm: q, User: user})
	if err != nil {
		log.Fatalf("failed to search: %v", err)
	}

	for index, result := range resp.GetResults() {
		fmt.Printf("%d. repo: %s | url: %s\n", index+1, result.GetRepo(), result.GetFileUrl())
	}
	if len(resp.GetResults()) == 0 {
		fmt.Println("no results found")
		fmt.Println()
	}
}
