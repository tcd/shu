package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/shu/internal/shu"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <issue url>",
	Short: "Add a new issue to track",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		issueURL := args[0]

		newIssue, err := shu.FromLink(issueURL)
		if err != nil {
			log.Fatal(err)
		}

		shoes, err := shu.GetIssues()
		if err != nil {
			log.Fatal(err)
		}

		err = shoes.Add(newIssue)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("New issue added: ", newIssue)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
