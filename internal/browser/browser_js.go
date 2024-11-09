//go:build js

package browser

import "syscall/js"

// OpenURL opens a new browser window pointing to url.
func OpenURL(url string) error {
	js.Global().Get("window").Get("open").Invoke(url)
	return nil
}
