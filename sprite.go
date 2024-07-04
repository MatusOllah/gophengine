package gophengine

import (
	"image"

	"github.com/MatusOllah/gophengine/internal/anim"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Position       image.Point
	Img            *ebiten.Image
	Visible        bool
	AnimController *anim.AnimController
}

func NewSprite(x, y int) *Sprite {
	return &Sprite{
		Position:       image.Pt(x, y),
		Visible:        true,
		AnimController: anim.NewAnimController(),
	}
}

func (s *Sprite) DrawImageOptions() *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Position.X), float64(s.Position.Y))
	return op
}

func (s *Sprite) Draw(img *ebiten.Image) {
	if !s.Visible {
		return
	}

	img.DrawImage(s.Img, s.DrawImageOptions())
}

func (s *Sprite) DrawWithOptions(img *ebiten.Image, opts *ebiten.DrawImageOptions) {
	if !s.Visible {
		return
	}

	img.DrawImage(s.Img, opts)
}
