package github_client

import (
	"context"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type Client struct {
	Github GithubServices
	Org    string
	Orgs   []string
}

func (c *Client) ID() string {
	return c.Org
}

func (c Client) withOrg(org string) *Client {
	return &Client{
		Github: c.Github,
		Org:    org,
		Orgs:   c.Orgs,
	}
}

func NewClients(config Config) ([]*Client, error) {
	client, err := newClient(config)
	if err != nil {
		return nil, err
	}
	return []*Client{client}, nil
}

func newClient(config Config) (*Client, error) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.AccessToken})
	tc := oauth2.NewClient(context.Background(), ts)
	c := github.NewClient(tc)
	return &Client{
		Github: GithubServices{
			Teams:         c.Teams,
			Billing:       c.Billing,
			Repositories:  c.Repositories,
			Organizations: c.Organizations,
			Issues:        c.Issues,
			Actions:       c.Actions,
			Dependabot:    c.Dependabot,
		},
		Orgs: config.Orgs,
	}, nil
}
