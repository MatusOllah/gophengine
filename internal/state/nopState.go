package state

import (
	ge "github.com/MatusOllah/gophengine"
	"github.com/hajimehoshi/ebiten/v2"
)

type NopState struct{}

var _ ge.State = (*NopState)(nil)

func (s *NopState) Draw(_ *ebiten.Image) {
}

func (s *NopState) Update(_ float64) error {
	return nil
}
