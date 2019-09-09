package shu

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/mitchellh/go-homedir"
)

// dataPath returns the absolute path to a user's "issues.json" file.
func dataPath() string {
	// dataDir := os.Getenv("XDG_DATA_HOME")
	// if dataDir != "" {
	// 	path = filepath.Join(dataDir, "shu")
	// } else {
	// }
	home, err := homedir.Dir()
	if err != nil {
		log.Println(err)
	}
	dataDir := filepath.Join(home, ".local/share/shu")

	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0777)
		if err != nil {
			log.Println("Error creating data directory: " + err.Error())
		}
	}

	dataFile := filepath.Join(dataDir, "issues.json")
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		err = writeToFile([]GitHubIssue{}, dataFile)
		if err != nil {
			log.Printf("Error creating projects file: %s\n", err)
		}
	}

	return dataFile
}

// GetGroups returns regular expression matches
func getGroups(regEx, testString string) (groupsMap map[string]string) {
	compiledRegEx := regexp.MustCompile(regEx)
	match := compiledRegEx.FindStringSubmatch(testString)
	groupsMap = make(map[string]string)
	for i, name := range compiledRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			groupsMap[name] = match[i]
		}
	}
	return groupsMap
}
