package services

import (
	"context"
	"fmt"
	"log"

	pb "weave-github-search/api/gh-search/v1"
	"weave-github-search/internal/github"
)

// TODO: these constants could be moved as parameters on SearchRequest in the proto file
const (
	defaultPageSize = 10
	pageNumber      = 1
)

// Server implements the GithubSearchService gRPC server.
type Server struct {
	pb.UnimplementedGithubSearchServiceServer
	GithubClient github.Client
}

// Search handles the SearchRequest and returns a SearchResponse.
func (s *Server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	term := req.GetSearchTerm()
	user := req.GetUser()

	codeSearchResult, err := s.GithubClient.SearchCode(ctx, term, user, defaultPageSize, pageNumber)
	if err != nil {
		return nil, fmt.Errorf("error searching code: %v", err)
	}
	log.Printf("found %d code results for term: %s, user: %s\n", codeSearchResult.GetTotal(), term, user)

	results := make([]*pb.SearchResult, 0)
	for _, codeResult := range codeSearchResult.CodeResults {
		result := &pb.SearchResult{
			FileUrl: codeResult.GetHTMLURL(),
			Repo:    codeResult.GetRepository().GetFullName(),
		}
		results = append(results, result)
	}
	return &pb.SearchResponse{Results: results}, nil
}
