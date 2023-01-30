package gophengine

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	X, Y  float64
	Image *ebiten.Image
}

func NewSprite(x, y float64) *Sprite {
	return &Sprite{
		X: x,
		Y: y,
	}
}

func (s *Sprite) Update(dt float64) error {
	return nil
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)

	screen.DrawImage(s.Image, op)
}
