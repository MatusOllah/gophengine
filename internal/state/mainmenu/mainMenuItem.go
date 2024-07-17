package mainmenu

import (
	"log/slog"

	ge "github.com/MatusOllah/gophengine"
)

type mainMenuItem struct {
	Name     string
	Sprite   *ge.Sprite
	OnSelect func(*mainMenuItem) error
}

func NopOnSelectFunc(i *mainMenuItem) error {
	slog.Warn(i.Name + " not implemented yet!")
	return nil
}
