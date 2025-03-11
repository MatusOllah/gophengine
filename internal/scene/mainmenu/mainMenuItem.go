package mainmenu

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/dialog"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
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

func (i *MainMenuItem) Update() {
	i.Sprite.AnimController.Update()
}

func (i *MainMenuItem) Draw(img *ebiten.Image) {
	if i.Sprite.Visible {
		i.Sprite.AnimController.Draw(img, i.Sprite.Position)
	}
}
