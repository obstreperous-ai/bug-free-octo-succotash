// Package version provides version information for the application.
// The version variables are set at build time using ldflags.
package version

var (
	// Version is the current version of the application
	Version = "dev"
	// BuildDate is the date the application was built
	BuildDate = "unknown"
	// GitCommit is the git commit hash
	GitCommit = "unknown"
)
