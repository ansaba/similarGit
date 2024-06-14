package git

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v62/github"
	"golang.org/x/oauth2"
)

func fetchSDKIssues(owner, repo string) ([]*github.Issue, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable not set")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	var maxIssues int = 10
	var issues []*github.Issue
	opts := &github.IssueListByRepoOptions{State: "all", ListOptions: github.ListOptions{PerPage: maxIssues}}
	for {
		issuebatch, resp, err := client.Issues.ListByRepo(context.Background(), owner, repo, opts)
		if err != nil {
			return nil, err
		}
		issues = append(issues, issuebatch...)

		if len(issues) >= maxIssues || resp.NextPage == 0 {
			break // Stop if we have enough issues or reached the last page
		}
		opts.Page = resp.NextPage
	}

	if len(issues) > maxIssues {
		issues = issues[:maxIssues] // Truncate to the desired number of issues
	}

	return issues, nil
}
