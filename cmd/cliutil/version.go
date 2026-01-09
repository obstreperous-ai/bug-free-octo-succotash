package cliutil

import (
	"fmt"

	"github.com/obstreperous-ai/bug-free-octo-succotash/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Display the version and build information for this application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version.Version)
		fmt.Printf("Build Date: %s\n", version.BuildDate)
		fmt.Printf("Commit: %s\n", version.GitCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
