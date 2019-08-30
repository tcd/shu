package cmd

import (
	"log"

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
		for _, shu := range shoes.I {
			log.Println(shu)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
