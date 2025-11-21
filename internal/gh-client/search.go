package gh_client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/go-github/v79/github"
)

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
