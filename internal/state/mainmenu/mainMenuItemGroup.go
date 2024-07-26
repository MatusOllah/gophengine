package mainmenu

import (
	"log/slog"
	"time"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/audioutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type mainMenuItemGroup struct {
	items       []*mainMenuItem
	curSelected int
	isSelected  bool
	flicker     *ge.Flicker
	itemFlicker *ge.Flicker
	magenta     *ge.Sprite
	ctx         *context.Context
}

func newMainMenuItemGroup(ctx *context.Context, items []*mainMenuItem, magenta *ge.Sprite) *mainMenuItemGroup {
	for i, item := range items {
		item.Sprite.Position.Y = 60 + (i * 160)
	}

	return &mainMenuItemGroup{
		items:       items,
		curSelected: 0,
		isSelected:  false,
		flicker:     ge.NewFlicker(magenta, 1100*time.Millisecond, 150*time.Millisecond),
		itemFlicker: ge.NewFlicker(nil, time.Second, 60*time.Millisecond),
		magenta:     magenta,
		ctx:         ctx,
	}
}

func (g *mainMenuItemGroup) Draw(screen *ebiten.Image) {
	for _, item := range g.items {
		if item.Sprite.Visible {
			item.Sprite.AnimController.Draw(screen, item.Sprite.Position)
		}
	}
}

func (g *mainMenuItemGroup) Update(dt float64) error {
	g.items[g.curSelected].Sprite.AnimController.Play("selected")

	for _, item := range g.items {
		item.Sprite.AnimController.UpdateWithDelta(dt)
	}

	if err := g.flicker.Update(); err != nil {
		return err
	}
	if err := g.itemFlicker.Update(); err != nil {
		return err
	}

	// Prevent the user from selecting when transitioning
	if g.isSelected {
		return nil
	}

	// Handle keyboard input for up navigation
	if g.ctx.InputHandler.ActionIsJustPressed(ge.ActionUp) {
		g.items[g.curSelected].Sprite.AnimController.Play("idle") // deselect old item

		if g.curSelected > 0 {
			if err := audioutil.PlaySoundFromFS(g.ctx, g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0); err != nil {
				return err
			}

			g.curSelected--
		}

		slog.Info("highlighted menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)
	}

	// Handle keyboard input for down navigation
	if g.ctx.InputHandler.ActionIsJustPressed(ge.ActionDown) {
		g.items[g.curSelected].Sprite.AnimController.Play("idle") // deselect old item

		if g.curSelected < len(g.items)-1 {
			if err := audioutil.PlaySoundFromFS(g.ctx, g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0); err != nil {
				return err
			}

			g.curSelected++
		}

		slog.Info("highlighted menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)
	}

	// Handle mouse wheel input
	_, yOffset := ebiten.Wheel()
	if yOffset != 0 {
		g.items[g.curSelected].Sprite.AnimController.Play("idle") // deselect old item

		if yOffset < 0 && g.curSelected < len(g.items)-1 {
			if err := audioutil.PlaySoundFromFS(g.ctx, g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0); err != nil {
				return err
			}
			g.curSelected++
		} else if yOffset > 0 && g.curSelected > 0 {
			if err := audioutil.PlaySoundFromFS(g.ctx, g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0); err != nil {
				return err
			}
			g.curSelected--
		}

		slog.Info("highlighted menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)
	}

	// Handle item selection keyboard
	if g.ctx.InputHandler.ActionIsJustPressed(ge.ActionAccept) {
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

		g.flicker.Flicker()

		for i := range g.items {
			if g.curSelected == i {
				continue
			}
			slog.Debug("hiding item", "i", i)
			g.items[i].Sprite.Visible = false // TODO: tween animation
		}

		g.itemFlicker.Sprite = g.items[g.curSelected].Sprite
		g.itemFlicker.OnCompleteCallback = func() error {
			if err := g.items[g.curSelected].OnSelect(g.items[g.curSelected]); err != nil {
				return err
			}

			return nil
		}
		g.itemFlicker.Flicker()

	}

	return nil
}
