package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	shu "github.com/tcd/shu/shu"
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

		table, _ := cmd.Flags().GetBool("table")
		if table {
			shu.Table(shoes)
			os.Exit(0)
		}

		for _, shu := range shoes.I {
			log.Println(shu)
		}
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().Bool("table", false, "Print issues in table format.")
}
