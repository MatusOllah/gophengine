package scene

import (
	ge "github.com/MatusOllah/gophengine"
	"github.com/hajimehoshi/ebiten/v2"
)

type NopScene struct{}

var _ ge.Scene = (*NopScene)(nil)

func (s *NopScene) Init() error {
	return nil
}

func (s *NopScene) Close() error {
	return nil
}

func (s *NopScene) Draw(_ *ebiten.Image) {
}

func (s *NopScene) Update(_ float64) error {
	return nil
}
