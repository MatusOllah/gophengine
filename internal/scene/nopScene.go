package scene

import (
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

type NopScene struct{}

var _ engine.Scene = (*NopScene)(nil)

func (s *NopScene) Init() error {
	return nil
}

func (s *NopScene) Close() error {
	return nil
}

func (s *NopScene) Draw(_ *ebiten.Image) {
}

func (s *NopScene) Update() error {
	return nil
}
