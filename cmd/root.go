package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// main cli command
var rootCmd = &cobra.Command{
	Use:   "httpserver",
	Short: "run http server",
	Long:  `run http server using the cobra command`,
}

// Execute commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
