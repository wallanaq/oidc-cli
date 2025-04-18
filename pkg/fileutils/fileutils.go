package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
)

// Save writes the given data to the specified file path.
// It ensures the parent directory exists by creating it if necessary.
//
// If the file already exists, it will be overwritten.
// The file is written with permission 0644, and the parent directories
// are created with permission 0755.
//
// Returns an error if the directory or file could not be created.
func Save(path string, data []byte) error {

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}

	return nil
}
