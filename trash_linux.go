package ostrash

import (
	"fmt"
	"os"
	"path/filepath"
)

func TrashItem(filePath string) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("go2trash: cannot resolve absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); err != nil {
		return fmt.Errorf("go2trash: file does not exist: %w", err)
	}

	home, err := filepath.Abs(filePath)
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

	return nil
}
