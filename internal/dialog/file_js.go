//go:build js

package dialog

import "errors"

func SelectFileOpen(title string, filename string, filters []FileFilter) (string, error) {
	return "", errors.New("file open dialog unsupported on js/wasm")
}

func SelectFileSave(title string, filename string, filters []FileFilter) (string, error) {
	return "", errors.New("file save dialog unsupported on js/wasm")
}
