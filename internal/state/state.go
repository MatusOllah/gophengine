package state

import "github.com/hajimehoshi/ebiten/v2"

type Updater interface {
	Update(dt float64) error
}

type Drawer interface {
	Draw(screen *ebiten.Image)
}

type State interface {
	Updater
	Drawer
}
