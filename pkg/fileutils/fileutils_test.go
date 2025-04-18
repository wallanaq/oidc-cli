package fileutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSave(t *testing.T) {

	path := filepath.Join(os.TempDir(), "oidc-cli-test", "testfile.txt")
	content := []byte("hello")

	err := Save(path, content)
	if err != nil {
		t.Fatal("Failed to save file:", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal("Failed to read back saved file:", err)
	}

	if string(data) != "hello" {
		t.Errorf("Expected content 'hello', got '%s'", string(data))
	}

	_ = os.RemoveAll(filepath.Dir(path))

}
