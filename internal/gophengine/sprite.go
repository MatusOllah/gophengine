package gophengine

import (
	"github.com/MatusOllah/gophengine/internal/anim"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	X, Y           float64
	Img            *ebiten.Image
	Visible        bool
	AnimController *anim.AnimController
}

func NewSprite(x, y float64) *Sprite {
	return &Sprite{
		X:              x,
		Y:              y,
		Visible:        true,
		AnimController: anim.NewAnimController(),
	}
}

func (s *Sprite) DrawImageOptions() *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	return op
}

func (s *Sprite) Draw(img *ebiten.Image) {
	if !s.Visible {
		return
	}

	img.DrawImage(s.Img, s.DrawImageOptions())
}
