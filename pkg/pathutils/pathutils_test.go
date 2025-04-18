package pathutils

import "testing"

func TestGetUserConfigDir(t *testing.T) {

	path, err := GetUserConfigDir()
	if err != nil {
		t.Fatal("Expected config path, got error:", err)
	}

	if path == "" {
		t.Fatal("Expected config path, got empty string")
	}

	t.Log("Config path:", path)
}
