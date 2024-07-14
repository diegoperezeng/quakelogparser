package utils

import (
	"testing"
)

func TestReadFile(t *testing.T) {

	content := "19:30 InitGame:"

	result, err := ReadFile("../../files/file_for_test.log")
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	if result != content {
		t.Fatalf("Expected content to be %s, got %s", content, result)
	}
}
