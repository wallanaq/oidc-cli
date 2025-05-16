package output

import (
	"os"
	"path/filepath"
)

type FileWriter struct {
	Path string
}

func (w FileWriter) Write(data []byte) error {
	dir := filepath.Dir(w.Path)
	if err := os.MkdirAll(dir, os.ModePerm&0755); err != nil {
		return err
	}
	return os.WriteFile(w.Path, data, os.ModePerm&0644)
}
