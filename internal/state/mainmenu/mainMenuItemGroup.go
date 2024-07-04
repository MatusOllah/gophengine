package mainmenu

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/audioutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type mainMenuItemGroup struct {
	items       []*mainMenuItem
	curSelected int
	isSelected  bool
	ctx         *context.Context
}

func newMainMenuItemGroup(ctx *context.Context, items ...*mainMenuItem) *mainMenuItemGroup {
	for i, item := range items {
		item.Sprite.Position.Y = 60 + (i * 160)
	}

	return &mainMenuItemGroup{
		items:       items,
		curSelected: 0,
		isSelected:  false,
		ctx:         ctx,
	}
}

func (g *mainMenuItemGroup) Draw(screen *ebiten.Image) {
	for _, item := range g.items {
		item.Sprite.AnimController.Draw(screen, item.Sprite.Position)
	}
}

func (g *mainMenuItemGroup) Update(dt float64) error {
	g.items[g.curSelected].Sprite.AnimController.Play("selected")

	for _, item := range g.items {
		item.Sprite.AnimController.UpdateWithDelta(dt)
	}

	if !g.isSelected {
		if inpututil.IsKeyJustPressed(g.ctx.Controls.Up) {
			g.items[g.curSelected].Sprite.AnimController.Play("idle")

			if g.curSelected != 0 {
				if err := audioutil.PlaySoundFromFS(g.ctx, g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0); err != nil {
					return err
				}

				g.curSelected--
			}

			slog.Info("highlighted menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)
		}

		if inpututil.IsKeyJustPressed(g.ctx.Controls.Down) {
			g.items[g.curSelected].Sprite.AnimController.Play("idle")

			if g.curSelected != len(g.items)-1 {
				if err := audioutil.PlaySoundFromFS(g.ctx, g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0); err != nil {
					return err
				}

				g.curSelected++
			}

			slog.Info("highlighted menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)
		}

		if inpututil.IsKeyJustPressed(g.ctx.Controls.Accept) {
			slog.Info("selected menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)

			if g.items[g.curSelected].Name == "donate" {
				// skip flicker anim
				if err := g.items[g.curSelected].OnSelect(g.items[g.curSelected]); err != nil {
					return err
				}
				return nil
			}

			g.isSelected = true

			if err := audioutil.PlaySoundFromFS(g.ctx, g.ctx.AssetsFS, "sounds/confirmMenu.ogg", -0.3); err != nil {
				return err
			}

			//TODO: flicker

			if err := g.items[g.curSelected].OnSelect(g.items[g.curSelected]); err != nil {
				return err
			}
		}
	}

	return nil
}
