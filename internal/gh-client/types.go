package gh_client

import (
	"context"
	"log"

	"github.com/google/go-github/v79/github"
)

// Client defines the interface for interacting with the GitHub API.
type Client interface {
	SearchCode(ctx context.Context, query, user string, perPage, page int) (*github.CodeSearchResult, error)
}

type client struct {
	*github.Client
}

// NewClient creates a new GitHub client with optional authentication.
func NewClient(ghToken string) Client {
	c := github.NewClient(nil)
	if ghToken != "" {
		c = c.WithAuthToken(ghToken)
	} else {
		log.Println("GitHub client created without authentication.")
	}
	return &client{Client: c}
}
