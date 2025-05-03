package output

import (
	"os"
	"path/filepath"
)

type Writer interface {
	Write(data []byte) error
}

type ConsoleWriter struct{}

type FileWriter struct {
	Path string
}

func NewWriter(output string) Writer {
	switch output {
	case "", "-":
		return &ConsoleWriter{}
	default:
		return &FileWriter{Path: output}
	}
}

func (w *ConsoleWriter) Write(data []byte) error {
	_, err := os.Stdout.Write(data)
	os.Stdout.WriteString("\n")
	return err
}

func (w FileWriter) Write(data []byte) error {
	dir := filepath.Dir(w.Path)
	if err := os.MkdirAll(dir, os.ModePerm&0755); err != nil {
		return err
	}
	return os.WriteFile(w.Path, data, os.ModePerm&0644)
}
