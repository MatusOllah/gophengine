package gophengine

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	X, Y  int
	Image *ebiten.Image
}

func NewSprite(x, y int) *Sprite {
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
	op.GeoM.Translate(float64(s.X), float64(s.Y))

	screen.DrawImage(s.Image, op)
}
