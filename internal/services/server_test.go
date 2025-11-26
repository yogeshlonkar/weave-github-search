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

func (m *mockGhc) SearchCode(ctx context.Context, query, user string, perPage, page int) (*github.CodeSearchResult, error) {
	args := m.Called(ctx, query, user, perPage, page)
	return args.Get(0).(*github.CodeSearchResult), args.Error(1)
}

func TestServer_Search(t *testing.T) {
	// mocking
	mockedGhc := new(mockGhc)
	total := 2
	mockedGhc.On("SearchCode", mock.Anything, "golang", "testuser", 10, 1).Return(&github.CodeSearchResult{
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
	assertT.NotNil(resp)
	assertT.Equal(2, len(resp.Results))
	assertT.Equal("example.com/file1", resp.Results[0].FileUrl)
	assertT.Equal("x/repo1", resp.Results[0].Repo)
	assertT.Equal("example.com/file2", resp.Results[1].FileUrl)
	assertT.Equal("x/repo2", resp.Results[1].Repo)
	mockedGhc.AssertExpectations(t)
}
