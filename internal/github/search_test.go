package github

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-github/v79/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
)

func TestClient_SearchCode(t *testing.T) {
	assertT := assert.New(t)

	// create client
	ghToken := "random-generated-token"
	c, err := NewClient(context.Background(), ghToken)
	assertT.NoError(err, "NewClient should not return an error")

	// mocking
	total := 2
	expected := github.CodeSearchResult{
		Total: &total,
		CodeResults: []*github.CodeResult{
			{
				HTMLURL:    github.Ptr("example.com/file1"),
				Repository: &github.Repository{Name: github.Ptr("repo1")},
			},
			{
				HTMLURL:    github.Ptr("example.com/file2"),
				Repository: &github.Repository{Name: github.Ptr("repo2")},
			},
		},
	}
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatchHandler(
			mock.GetSearchCode,
			http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				_, err := w.Write(mock.MustMarshal(expected))
				if err != nil {
					t.Fatalf("failed to write mock response: %v", err)
				}
			}),
		),
	)
	c.(*client).Client = github.NewClient(mockedHTTPClient)

	query := "example search"
	user := ""
	results, err := c.SearchCode(context.Background(), query, user)

	// assertions
	assertT.NoError(err, "SearchCode should not return an error")
	assertT.Equal(&expected, results, "SearchCode should return expected results")
}
