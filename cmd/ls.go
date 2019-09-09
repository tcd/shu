package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tcd/shu/internal/shu"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all tracked issues",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		shoes, err := shu.GetIssues()
		if err != nil {
			log.Fatal(err)
		}
		shoes.SortByRepo()

		table, _ := cmd.Flags().GetBool("table")
		if table {
			shu.Table(shoes)
			os.Exit(0)
		}

		url, _ := cmd.Flags().GetBool("url")
		if url {
			for _, shu := range shoes.I {
				fmt.Println(shu.IssueURL())
			}
			os.Exit(0)
		}

		for _, shu := range shoes.I {
			fmt.Println(shu)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().Bool("table", false, "Print issues in table format.")
	lsCmd.Flags().Bool("url", false, "Print issue urls")
}
