package gh_search_service

import (
	"context"
	"testing"
	pb "weave-github-search/api/gh-search/v1"
)

func TestServer_Search(t *testing.T) {
	server := &Server{}
	req := &pb.SearchRequest{
		SearchTerm: "golang",
		User:       "testuser",
	}

	resp, err := server.Search(context.Background(), req)
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if resp == nil {
		t.Fatalf("Expected non-nil response")
	}
}
