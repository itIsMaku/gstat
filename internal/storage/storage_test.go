package storage

import (
	"gstat/internal/protocol"
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	dir := "test_history_dir"
	defer os.RemoveAll(dir)

	CreateHistoryDirectory(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Fatal(dir + " does not exist")
	}

	result := Save(dir, protocol.Result{
		Target:    "example.com",
		Protocol:  protocol.HTTP,
		Reachable: true,
		Message:   "OK",
	})
	if !result {
		t.Fatal("Failed to save result")
	}

	Read(dir)
}
