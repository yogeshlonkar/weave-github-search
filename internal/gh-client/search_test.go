package gh_client

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
	ghToken := "random-generated-token"
	c, err := NewClient(ghToken)
	assertT.NoError(err, "NewClient should not return an error")
	total := 2
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatchHandler(
			mock.GetSearchCode,
			http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Write(mock.MustMarshal(github.CodeSearchResult{
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
				}))
			}),
		),
	)
	c.(*client).Client = github.NewClient(mockedHTTPClient)

	query := "example search"
	user := ""
	perPage := 10
	page := 1

	results, err := c.SearchCode(context.Background(), query, user, perPage, page)
	assertT.NoError(err, "SearchCode should not return an error")
	assertT.NotNil(results, "Results should not be nil")
	assertT.Equal(total, results.GetTotal(), "Total results should match expected value")
	for _, result := range results.CodeResults {
		t.Logf("File URL: %s, Repo: %s", result.GetHTMLURL(), result.GetRepository())
	}
}
