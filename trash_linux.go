//go:build linux

package go2trash

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// MoveToTrash sends a file or directory to the FreeDesktop.org Trash.
//
// Creates a .trashinfo metadata file alongside the trashed file to
// display and restore the file.
func MoveToTrash(filePath string) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("go2trash: cannot resolve absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); err != nil {
		return fmt.Errorf("go2trash: file does not exist: %w", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("go2trash: cannot find home directory: %w", err)
	}

	trash := filepath.Join(home, ".local", "share", "Trash")
	filesDir := filepath.Join(trash, "files")
	infoDir := filepath.Join(trash, "info")

	// check if trash subdirs exist. if doesn't exist, mkdir (chmod 0700), else error.
	if err := os.MkdirAll(filesDir, 0700); err != nil {
		return fmt.Errorf("go2trash: cannot create trash files dir: %w", err)
	}
	if err := os.MkdirAll(infoDir, 0700); err != nil {
		return fmt.Errorf("go2trash: cannot create trash info dir: %w", err)
	}

	fileName := filepath.Base(absPath)
	trashName := fileName

	infoCheck := filepath.Join(infoDir, trashName+".trashinfo")
	if _, err := os.Stat(infoCheck); err == nil {
		trashName = appendTime(fileName)
	}

	trashFile := filepath.Join(filesDir, trashName)
	trashInfo := filepath.Join(infoDir, trashName+".trashinfo")

	infoContent := fmt.Sprintf(
		"[Trash Info]\nPath=%s\nDeletionDate=%s\n",
		absPath,
		time.Now().Format("20060102T03:04:05"),
	)

	f, err := os.OpenFile(trashInfo, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("go2trash: cannot create trashinfo: %w", err)
	}
	f.WriteString(infoContent)
	f.Close()

	// move file to the trash files directory
	if err := os.Rename(absPath, trashFile); err != nil {
		os.Remove(trashInfo) // rollback on failure
		return fmt.Errorf("go2trash: cannot move file to trash: %w", err)
	}

	return nil
}

// Add timestamp to avoid collision in trash dir
func appendTime(name string) string {
	ext := filepath.Ext(name)
	fileName := strings.TrimSuffix(name, ext)

	t := time.Now()
	timeStr := t.Format("03.04.05 PM")

	return fileName + " " + timeStr + ext
}
