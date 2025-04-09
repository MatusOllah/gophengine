package config

import (
	"errors"
	"os"
)

// fileExists checks if file exists.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return false
	}
}
