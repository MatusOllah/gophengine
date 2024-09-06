package scene

import (
	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/scene/optionsui"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type OptionsScene struct {
	ctx        *context.Context
	bg         *ebiten.Image
	ui         *ebitenui.UI
	shouldExit bool
}

func NewOptionsScene(ctx *context.Context) *OptionsScene {
	return &OptionsScene{ctx: ctx}
}

var _ ge.Scene = (*OptionsScene)(nil)

func (s *OptionsScene) Init() error {
	s.shouldExit = false

	bg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/menuDesat.png")
	if err != nil {
		return err
	}
	s.bg = bg

	ui, err := optionsui.MakeUI(s.ctx, &s.shouldExit)
	if err != nil {
		return err
	}
	s.ui = ui

	return nil
}

func (s *OptionsScene) Close() error {
	return nil
}

func (s *OptionsScene) Draw(screen *ebiten.Image) {
	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts.GeoM.Scale(1.1, 1.1)
	screen.DrawImage(s.bg, bgOpts)

	s.ui.Draw(screen)
}

func (s *OptionsScene) Update(_ float64) error {
	if s.shouldExit {
		return s.ctx.SceneCtrl.SwitchScene(&MainMenuScene{ctx: s.ctx})
	}

	s.ui.Update()

	return nil
}
