package version

import "testing"

func TestVersion(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected string
	}{
		{"dev version", "dev", "dev"},
		{"release version", "1.0.0", "1.0.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version = tt.version
			if Version != tt.expected {
				t.Errorf("Version = %v, want %v", Version, tt.expected)
			}
		})
	}
}
