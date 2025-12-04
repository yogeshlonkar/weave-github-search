package services

import (
	"context"
	"testing"

	"github.com/google/go-github/v79/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	pb "weave-github-search/api/gh-search/v1"
	_github "weave-github-search/internal/github"
)

type mockGhc struct {
	mock.Mock
	_github.Client
}

func (m *mockGhc) SearchCode(ctx context.Context, query, user string) (*github.CodeSearchResult, error) {
	args := m.Called(ctx, query, user)
	return args.Get(0).(*github.CodeSearchResult), args.Error(1)
}

func TestServer_Search(t *testing.T) {
	// mocking
	mockedGhc := new(mockGhc)
	total := 2
	mockedGhc.On("SearchCode", mock.Anything, "golang", "testuser").Return(&github.CodeSearchResult{
		Total: &total,
		CodeResults: []*github.CodeResult{
			{
				HTMLURL:    github.Ptr("example.com/file1"),
				Repository: &github.Repository{FullName: github.Ptr("x/repo1")},
			},
			{
				HTMLURL:    github.Ptr("example.com/file2"),
				Repository: &github.Repository{FullName: github.Ptr("x/repo2")},
			},
		},
	}, nil)

	// test server
	server := &Server{GithubClient: mockedGhc}
	req := &pb.SearchRequest{
		SearchTerm: "golang",
		User:       "testuser",
	}

	resp, err := server.Search(context.Background(), req)

	// assertions
	assertT := assert.New(t)
	assertT.NoError(err)
	assertT.Equal(&pb.SearchResponse{
		Results: []*pb.SearchResult{
			{FileUrl: "example.com/file1", Repo: "x/repo1"},
			{FileUrl: "example.com/file2", Repo: "x/repo2"},
		},
	}, resp)
	mockedGhc.AssertExpectations(t)
}
