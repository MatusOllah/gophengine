//go:build !js

package dialog

import "github.com/ncruces/zenity"

func Error(msg string) {
	if err := zenity.Error(msg, zenity.Title("GophEngine error")); err != nil {
		panic(err)
	}
}

func Warning(msg string) {
	if err := zenity.Warning(msg, zenity.Title("GophEngine warning")); err != nil {
		panic(err)
	}
}
