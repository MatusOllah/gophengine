package gophengine

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Init() error
	Close() error
	Update(dt float64) error
	Draw(screen *ebiten.Image)
}
