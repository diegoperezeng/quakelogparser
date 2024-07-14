package utils

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	content := "test content"
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Unable to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("Unable to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Unable to close temp file: %v", err)
	}

	result, err := ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	if result != content {
		t.Fatalf("Expected content to be %s, got %s", content, result)
	}
}
