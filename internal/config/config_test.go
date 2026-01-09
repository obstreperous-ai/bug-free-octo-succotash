package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := New()

	if cfg.Verbose != false {
		t.Errorf("Expected Verbose to be false, got %v", cfg.Verbose)
	}

	if cfg.Output != "" {
		t.Errorf("Expected Output to be empty, got %v", cfg.Output)
	}
}

func TestGetConfigDir(t *testing.T) {
	dir, err := GetConfigDir()
	if err != nil {
		t.Fatalf("GetConfigDir() returned error: %v", err)
	}

	home, _ := os.UserHomeDir()
	expected := filepath.Join(home, ".config", "succotash")

	if dir != expected {
		t.Errorf("GetConfigDir() = %v, want %v", dir, expected)
	}
}
