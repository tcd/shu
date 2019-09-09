package cmd

import (
	"os"

	"github.com/spf13/cobra"
	shu "github.com/tcd/shu/internal/server"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run the shu web interface",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		if port != "" {
			shu.ServeToPort(port)
			os.Exit(0)
		}

		shu.Serve()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().StringP("port", "p", "", "Port on which to run the server")
}
