package shu

// GetIssues fetches GitHubIssues from a user's issues.json file.
func GetIssues() (Shoes, error) {
	var shoes Shoes
	issues, err := readFromFile(dataPath())
	if err != nil {
		return shoes, err
	}
	shoes.I = issues
	return shoes, nil
}

// SaveIssues writes a slice of GitHubIssue to a user's issues.json file.
func SaveIssues(issues []GitHubIssue) error {
	return writeToFile(issues, dataPath())
}
