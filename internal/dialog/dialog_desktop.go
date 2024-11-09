//go:build !wasm

package dialog

import "github.com/ncruces/zenity"

func Error(msg string) {
	zenity.Error(msg, zenity.Title("GophEngine error"))
}

func Warning(msg string) {
	zenity.Warning(msg, zenity.Title("GophEngine warning"))
}
