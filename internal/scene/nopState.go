package scene

import (
	ge "github.com/MatusOllah/gophengine"
	"github.com/hajimehoshi/ebiten/v2"
)

type NopScene struct{}

var _ ge.State = (*NopScene)(nil)

func (s *NopScene) Draw(_ *ebiten.Image) {
}

func (s *NopScene) Update(_ float64) error {
	return nil
}
