package http

import (
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		reachable bool
	}{
		{"http-reachable", "https://store.rcore.cz", true},
		{"http-non-reachable", "https://store.rcore.wtf", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Check(tt.url)
			if result.Reachable != tt.reachable {
				t.Fatalf("Test failed, expected reachable=%v, got reachable=%v", tt.reachable, result.Reachable)
			}
		})
	}
}
