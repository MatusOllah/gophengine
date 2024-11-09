//go:build js

package dialog

import "syscall/js"

func Error(msg string) {
	js.Global().Get("alert").Invoke("Error: " + msg)
}

func Warning(msg string) {
	js.Global().Get("alert").Invoke("Warning: " + msg)

}
