package mainmenu

import (
	"log/slog"
	"time"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/MatusOllah/gophengine/internal/controls"
	"github.com/MatusOllah/gophengine/internal/effects"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

type MainMenuItemGroup struct {
	items       []*MainMenuItem
	group       *engine.Group[*MainMenuItem]
	curSelected int
	isSelected  bool
	flicker     *effects.Flicker
	itemFlicker *effects.Flicker
	magenta     *engine.Sprite
	ctx         *context.Context
}

func NewMainMenuItemGroup(ctx *context.Context, items []*MainMenuItem, magenta *engine.Sprite) *MainMenuItemGroup {
	for i, item := range items {
		item.Sprite.Position.Y = 60 + (i * 160)
	}

	return &MainMenuItemGroup{
		items:       items,
		group:       engine.NewGroup(items...),
		curSelected: 0,
		isSelected:  false,
		flicker:     effects.NewFlicker(magenta, 1100*time.Millisecond, 150*time.Millisecond),
		itemFlicker: effects.NewFlicker(nil, time.Second, 60*time.Millisecond),
		magenta:     magenta,
		ctx:         ctx,
	}
}

func (g *MainMenuItemGroup) Draw(screen *ebiten.Image) {
	g.group.Draw(screen)
}

func (g *MainMenuItemGroup) Update() error {
	g.items[g.curSelected].Sprite.AnimController.Play("selected")

	g.group.Update()

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
	if g.ctx.InputHandler.ActionIsJustPressed(controls.ActionUp) {
		g.items[g.curSelected].Sprite.AnimController.Play("idle") // deselect old item

		if g.curSelected > 0 {
			if err := audio.PlaySoundFromFS(g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0, g.ctx.AudioMixer.SFX); err != nil {
				return err
			}

			g.curSelected--
		}

		slog.Info("highlighted menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)
	}

	// Handle keyboard input for down navigation
	if g.ctx.InputHandler.ActionIsJustPressed(controls.ActionDown) {
		g.items[g.curSelected].Sprite.AnimController.Play("idle") // deselect old item

		if g.curSelected < len(g.items)-1 {
			if err := audio.PlaySoundFromFS(g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0, g.ctx.AudioMixer.SFX); err != nil {
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
			if err := audio.PlaySoundFromFS(g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0, g.ctx.AudioMixer.SFX); err != nil {
				return err
			}
			g.curSelected++
		} else if yOffset > 0 && g.curSelected > 0 {
			if err := audio.PlaySoundFromFS(g.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0, g.ctx.AudioMixer.SFX); err != nil {
				return err
			}
			g.curSelected--
		}

		slog.Info("highlighted menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)
	}

	// Handle item selection (keyboard)
	if g.ctx.InputHandler.ActionIsJustPressed(controls.ActionAccept) {
		slog.Info("selected menu item", "item", g.items[g.curSelected].Name, "i", g.curSelected)

		if g.items[g.curSelected].Name == "donate" {
			// skip flicker anim
			if err := g.items[g.curSelected].OnSelect(g.items[g.curSelected]); err != nil {
				return err
			}
			return nil
		}

		g.isSelected = true

		if err := audio.PlaySoundFromFS(g.ctx.AssetsFS, "sounds/confirmMenu.ogg", -0.3, g.ctx.AudioMixer.SFX); err != nil {
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
