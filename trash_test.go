package go2trash_test

import (
	"os"
	"path/filepath"
	"testing"

	go2trash "github.com/rafshawn/go2trash"
)

func TestMoveToTrash(t *testing.T) {
	// Create a temporary file to trash
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_file.txt")

	if err := os.WriteFile(testFile, []byte("move this file to trash"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Verify file exists before trashing
	if _, err := os.Stat(testFile); err != nil {
		t.Fatalf("Test file should exist before trashing: %v", err)
	}

	// Move to OS Trash
	if err := go2trash.MoveToTrash(testFile); err != nil {
		t.Fatalf("MoveToTrash failed: %v", err)
	}

	// Verify file is gone from original location
	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Fatalf("File should not exist at original location after trashing, but got: %v", err)
	}

	t.Logf("File successfully sent to OS Trash")
}

func TestMoveToTrash_Noexist(t *testing.T) {
	err := go2trash.MoveToTrash("/tmp/nonexistent_file.txt")
	if err == nil {
		t.Fatal("Expected an error when trashing a non-existent file, got nil")
	}
	t.Logf("Correctly returned error for non-existent file: %v", err)
}
