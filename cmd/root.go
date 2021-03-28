package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "coffee",
	Version:      "1.0.0",
	Short:        "CLI application for the coffee shop",
	Long:         "CLI application for the coffee shop",
	SilenceUsage: true,
}

// Execute _
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
