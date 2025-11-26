package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v79/github"
)

// SearchCode searches GitHub code based on the provided query and user.
// If a user is specified, the search is limited to that user's repositories.
// It returns the search results or an error if the search fails.
func (c *client) SearchCode(ctx context.Context, query, user string, perPage, page int) (*github.CodeSearchResult, error) {
	if user != "" {
		query += " user:" + user
	}
	log.Printf("executing GitHub code search with query: %s\n", query)
	opts := &github.SearchOptions{
		TextMatch: true,
		Sort:      "created",
		Order:     "asc",
		ListOptions: github.ListOptions{
			PerPage: perPage,
			Page:    page,
		},
	}

	codeResult, resp, err := c.Client.Search.Code(ctx, query, opts)
	if err != nil {
		return nil, fmt.Errorf("error executing GitHub code search: %v", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API returned status code %d", resp.StatusCode)
	}

	return codeResult, nil
}
