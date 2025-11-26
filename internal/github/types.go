package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v79/github"
	"golang.org/x/oauth2"
)

// Client defines the interface for interacting with the GitHub API.
type Client interface {
	SearchCode(ctx context.Context, query, user string, perPage, page int) (*github.CodeSearchResult, error)
}

type client struct {
	*github.Client
}

// NewClient creates a new GitHub client with optional authentication.
func NewClient(ctx context.Context, ghToken string) (Client, error) {
	if ghToken == "" {
		return nil, fmt.Errorf("GitHub client can not be created without authentication")
	}

	c := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	))
	return &client{Client: github.NewClient(c)}, nil
}
