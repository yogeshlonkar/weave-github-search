package gh_client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/go-github/v79/github"
)

// SearchCode searches GitHub code based on the provided query and user.
// If a user is specified, the search is limited to that user's repositories.
// It returns the search results or an error if the search fails.
func (c *client) SearchCode(ctx context.Context, query, user string, perPage, page int) (*github.CodeSearchResult, error) {
	if user != "" {
		query += " user:" + user
	}
	query = url.QueryEscape(query)
	opts := &github.SearchOptions{Sort: "created", Order: "asc", ListOptions: github.ListOptions{PerPage: perPage, Page: page}}

	codeResult, resp, err := c.Client.Search.Code(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API returned status code %d", resp.StatusCode)
	}

	return codeResult, nil
}
