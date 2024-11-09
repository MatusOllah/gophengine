//go:build js

package browser

import "syscall/js"

func OpenURL(url string) error {
	js.Global().Get("window").Get("open").Invoke(url)
	return nil
}
