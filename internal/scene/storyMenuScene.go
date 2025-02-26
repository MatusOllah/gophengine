package scene

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
Week Names (some names changed):

  - Week #1 - Daddy Dearest
  - Week #2 - Spooky Month
  - Week #3 - Pico
  - Week #4 - Mommy Mearest
  - Week #5 - Happy and Merry
  - Week #6 - dating simulator ft. moawling
*/
type StoryMenuScene struct {
	ctx *context.Context
}

var _ engine.Scene = (*StoryMenuScene)(nil)

func NewStoryMenuScene(ctx *context.Context) *StoryMenuScene {
	return &StoryMenuScene{ctx: ctx}
}

func (s *StoryMenuScene) Init() error {
	return nil
}

func (s *StoryMenuScene) Close() error {
	return nil
}

func (s *StoryMenuScene) Draw(_ *ebiten.Image) {

}

func (s *StoryMenuScene) Update() error {
	return nil
}
