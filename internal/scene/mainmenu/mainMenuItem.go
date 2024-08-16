package mainmenu

import (
	"log/slog"

	ge "github.com/MatusOllah/gophengine"
	"github.com/ncruces/zenity"
)

type MainMenuItem struct {
	Name     string
	Sprite   *ge.Sprite
	OnSelect func(*MainMenuItem) error
}

func NopOnSelectFunc(i *MainMenuItem) error {
	slog.Warn(i.Name + " not implemented yet!")
	return zenity.Warning(i.Name + " not implemented yet!")
}
