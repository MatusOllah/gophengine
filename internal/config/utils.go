package config

import (
	"errors"
	"os"
)

// exists checks if file exists.
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}
	panic("unreachable")
}
