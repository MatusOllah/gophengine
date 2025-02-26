package scene

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
	// TODO: move freakyMenu music to some MusicManager struct
	scoreTextFace    *text.GoTextFace
	txtWeekTitleFace *text.GoTextFace
}

var _ engine.Scene = (*StoryMenuScene)(nil)

func NewStoryMenuScene(ctx *context.Context) *StoryMenuScene {
	return &StoryMenuScene{ctx: ctx}
}

func (s *StoryMenuScene) loadFont(path string, size float64) (*text.GoTextFace, error) {
	f, err := s.ctx.AssetsFS.Open(path)
	if err != nil {
		return nil, err
	}

	src, err := text.NewGoTextFaceSource(f)
	if err != nil {
		return nil, err
	}

	return &text.GoTextFace{Source: src, Size: size}, nil
}

func (s *StoryMenuScene) Init() (err error) {
	s.scoreTextFace, err = s.loadFont("fonts/better-vcr-tweaked.ttf", 32)
	if err != nil {
		return
	}

	s.txtWeekTitleFace = s.scoreTextFace // they're the same font and size

	return nil
}

func (s *StoryMenuScene) Close() error {
	return nil
}

func (s *StoryMenuScene) Draw(screen *ebiten.Image) {
	{
		op := &text.DrawOptions{}
		op.GeoM.Translate(10, 10)
		text.Draw(screen, "SCORE: 49324858", s.scoreTextFace, op)
	}
	{
		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(engine.GameWidth)*0.7, 10)
		op.ColorScale.ScaleAlpha(0.7)
		op.PrimaryAlign = text.AlignEnd
		text.Draw(screen, "test week name", s.txtWeekTitleFace, op)
	}
}

func (s *StoryMenuScene) Update() error {
	return nil
}
