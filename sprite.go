package main

import "github.com/hajimehoshi/ebiten/v2"

type sprite struct {
	x, y  int
	image *ebiten.Image
}

func newSprite(x, y int) *sprite {
	return &sprite{
		x: x,
		y: y,
	}
}

func (s *sprite) update(dt float64) error {
	return nil
}

func (s *sprite) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x), float64(s.y))

	screen.DrawImage(s.image, op)
}
