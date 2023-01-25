package main

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Update(dt float64) error
	Draw(screen *ebiten.Image)
}
