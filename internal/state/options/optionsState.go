package options

import (
	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type OptionsState struct {
	bg *ebiten.Image
	ui *ebitenui.UI
}

func NewOptionsState(ctx *context.Context) (*OptionsState, error) {
	bg, _, err := ebitenutil.NewImageFromFileSystem(ctx.AssetsFS, "images/menuDesat.png")
	if err != nil {
		return nil, err
	}

	ui, err := makeUI(ctx)
	if err != nil {
		return nil, err
	}

	return &OptionsState{
		bg: bg,
		ui: ui,
	}, nil
}

var _ ge.State = (*OptionsState)(nil)

func (s *OptionsState) Draw(screen *ebiten.Image) {
	bgOpts := &ebiten.DrawImageOptions{}
	bgOpts.GeoM.Scale(1.1, 1.1)
	screen.DrawImage(s.bg, bgOpts)

	s.ui.Draw(screen)
}

func (s *OptionsState) Update(_ float64) error {
	s.ui.Update()

	return nil
}
