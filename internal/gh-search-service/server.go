package gh_search_service

import (
	"context"
	"log"

	pb "weave-github-search/api/gh-search/v1"
)

type Server struct {
	pb.UnimplementedGithubSearchServiceServer
}

func (s *Server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Printf("Search term: %v", req.GetSearchTerm())
	log.Printf("User: %v", req.GetUser())

	return &pb.SearchResponse{}, nil
}
