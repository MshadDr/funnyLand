package cmd

import (
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "httpserver",
	Aliases: []string{"serv"},
	Short:   "Run http server",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
