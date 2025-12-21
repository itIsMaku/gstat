package configuration

import (
	"os"
	"testing"
)

const testConfigFile = "test_config.json"

func Test(t *testing.T) {
	err := CreateConfig(testConfigFile)
	if err != nil {
		t.Fatalf("Failed to create configuration: %v", err)
	}

	defer os.Remove(testConfigFile)

	config, err := LoadConfig(testConfigFile)
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	if config.Interval.Seconds <= 0 {
		t.Fatalf("Invalid interval seconds: %d", config.Interval.Seconds)
	}
}
