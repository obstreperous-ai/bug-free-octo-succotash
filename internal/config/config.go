package config

import (
	"os"
	"path/filepath"
)

// Config holds the application configuration
type Config struct {
	Verbose bool
	Output  string
}

// New creates a new Config with default values
func New() *Config {
	return &Config{
		Verbose: false,
		Output:  "",
	}
}

// GetConfigDir returns the user's configuration directory
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(home, ".config", "cliutil")
	return configDir, nil
}
