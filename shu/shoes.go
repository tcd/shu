package shu

import "fmt"

// Shoes is a wrapper around a slice of GitHubIssue.
type Shoes struct {
	I []GitHubIssue
}

// Save writes issues to a user's issues.json file.
func (shoes Shoes) Save() error {
	return writeToFile(shoes.I, dataPath())
}

// Add an issue & save to a user's issues.json file.
func (shoes *Shoes) Add(newShoes ...GitHubIssue) (err error) {
	for _, shu := range newShoes {
		err = shoes.add(shu)
		if err != nil {
			return
		}
	}
	return
}

func (shoes *Shoes) add(issue GitHubIssue) error {
	for _, shu := range shoes.I {
		if issue.IssueURL() == shu.IssueURL() {
			return fmt.Errorf("issue already tracked, %q", shu)
		}
	}
	shoes.I = append(shoes.I, issue)
	return shoes.Save()
}
