package mainmenu

import (
	"log/slog"

	ge "github.com/MatusOllah/gophengine"
	"github.com/ncruces/zenity"
)

type mainMenuItem struct {
	Name     string
	Sprite   *ge.Sprite
	OnSelect func(*mainMenuItem) error
}

func NopOnSelectFunc(i *mainMenuItem) error {
	slog.Warn(i.Name + " not implemented yet!")
	return zenity.Warning(i.Name + " not implemented yet!")
}
