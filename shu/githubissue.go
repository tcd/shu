package shu

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/google/go-github/github"
)

const (
	apiBase = "https://api.github.com"
	webBase = "https://github.com"
)

// GitHubIssue models a GitHub issue.
type GitHubIssue struct {
	RepoOwner string    `json:"repoOwner"`
	RepoName  string    `json:"repoName"`
	Number    int       `json:"number"`
	Title     string    `json:"title"`
	Open      bool      `json:"open"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (shu GitHubIssue) String() string {
	return fmt.Sprintf("%s/%s #%d: %s", shu.RepoOwner, shu.RepoName, shu.Number, shu.Title)
}

// Repo returns the repo's owner/name.
func (shu GitHubIssue) Repo() string {
	return shu.RepoOwner + "/" + shu.RepoName
}

// RepoURL for a issue.
func (shu GitHubIssue) RepoURL() string {
	return fmt.Sprintf("https://github.com/%s/%s", shu.RepoOwner, shu.RepoName)
}

// IssueURL for an issue.
func (shu GitHubIssue) IssueURL() string {
	return fmt.Sprintf("https://github.com/%s/%s/issues/%d", shu.RepoOwner, shu.RepoName, shu.Number)
}

// FromLink returns a new GitHubIssue from a link to an existing issue.
func FromLink(link string) (GitHubIssue, error) {
	var ghi GitHubIssue
	regex := `https://github.com/(?P<owner>(?:\w|-)+)/(?P<repo>(?:\w|-)+)/issues/(?P<number>[0-9]+)`

	matches := getGroups(regex, link)
	if len(matches) != 3 {
		return ghi, errors.New("Error parsing values from link")
	}
	ghi.RepoOwner = matches["owner"]
	ghi.RepoName = matches["repo"]
	ghi.Number, _ = strconv.Atoi(matches["number"])

	issue, err := FetchIssue(ghi.RepoOwner, ghi.RepoName, ghi.Number)
	if err != nil {
		return ghi, err
	}
	ghi.Title = issue.GetTitle()
	ghi.CreatedAt = issue.GetCreatedAt()
	ghi.UpdatedAt = issue.GetUpdatedAt()
	if issue.GetState() != "closed" {
		ghi.Open = true
	}

	return ghi, nil
}

// FetchIssue gets information on a GitHub issue.
func FetchIssue(owner, repo string, number int) (*github.Issue, error) {
	client := github.NewClient(nil)
	issue, _, err := client.Issues.Get(context.Background(), owner, repo, number)
	if err != nil {
		return &github.Issue{}, err
	}
	return issue, err
}

// writeToFile an array of GitHubIssues to a json file.
func writeToFile(issues []GitHubIssue, path string) error {
	bytes, err := json.MarshalIndent(issues, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// readFromFile returns a slice of GitHubIssue from an issues.json file.
func readFromFile(path string) ([]GitHubIssue, error) {
	var issues []GitHubIssue
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return issues, err
	}
	err = json.Unmarshal(bytes, &issues)
	if err != nil {
		return issues, err
	}
	return issues, nil
}
