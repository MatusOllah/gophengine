package mainmenu

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/dialog"
	"github.com/MatusOllah/gophengine/internal/engine"
)

type MainMenuItem struct {
	Name     string
	Sprite   *engine.Sprite
	OnSelect func(*MainMenuItem) error
}

func NopOnSelectFunc(i *MainMenuItem) error {
	slog.Warn(i.Name + " not implemented yet!")
	dialog.Warning(i.Name + " not implemented yet!")
	return nil
}
