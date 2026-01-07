package cliutil

import (
	"fmt"
	"os"

	"github.com/obstreperous-ai/bug-free-octo-succotash/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cliutil",
	Short: "A command line utility for Linux",
	Long: `bug-free-octo-succotash is a command line utility for Linux.
Built with Go and designed for extensibility.`,
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to cliutil! Use --help to see available commands.")
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
}
