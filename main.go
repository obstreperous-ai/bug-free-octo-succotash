package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// Version information - can be overridden at build time
	Version   = "0.1.0"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

func main() {
	// Define version flag following Linux conventions
	versionFlag := flag.Bool("version", false, "display version information")
	versionShort := flag.Bool("v", false, "display version information (short)")

	flag.Parse()

	// Check if version flag is set
	if *versionFlag || *versionShort {
		printVersion()
		os.Exit(0)
	}

	// Main application logic
	fmt.Println("bug-free-octo-succotash - An experimental command line utility")
	fmt.Println("Use --version or -v to display version information")
}

func printVersion() {
	fmt.Printf("bug-free-octo-succotash version %s\n", Version)
	fmt.Printf("Git commit: %s\n", GitCommit)
	fmt.Printf("Build date: %s\n", BuildDate)
}
