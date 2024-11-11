package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GetConfigDir returns the path to the configuration directory for the application.
// This function checks the default locations for configuration storage based
// on the operating system and user environment.
//
// Returns:
//
//	string - the configuration directory path
//	error - if unable to determine the configuration directory
func GetConfigDir() (string, error) {

	var basePath string

	switch runtime.GOOS {
	case "windows":

		basePath = os.Getenv("APPDATA")

		if basePath == "" {
			return "", fmt.Errorf("environment variable is not set: APPDATA")
		}

	case "linux", "darwin":

		homeDir, err := os.UserHomeDir()

		if err != nil {
			return "", fmt.Errorf("could not get user home directory: %w", err)
		}

		basePath = filepath.Join(homeDir, ".config")

	default:
		return "", fmt.Errorf("unsupported platform %s", runtime.GOOS)
	}

	configDir := filepath.Join(basePath, "jencli")

	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("could not create configuration directory: %w", err)
	}

	return configDir, nil

}
