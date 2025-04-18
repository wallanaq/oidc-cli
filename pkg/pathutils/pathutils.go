package pathutils

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetUserConfigDir returns the platform-specific user configuration directory
// for the oidc-cli application.
//
// On Unix systems, it returns:   $HOME/.config/oidc-cli
// On Windows systems, it returns: %AppData%\oidc-cli
//
// Returns an error if the user's home/config directory cannot be determined.
func GetUserConfigDir() (string, error) {

	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")

		if appData == "" {
			return "", os.ErrNotExist
		}

		return filepath.Join(appData, "oidc-cli"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", "oidc-cli"), nil

}
