package scene

import (
	"fmt"
	"image/color"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/audioutil"
	"github.com/MatusOllah/gophengine/internal/controls"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/scene/storymenu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StoryMenuScene struct {
	ctx *context.Context
	// TODO: move freakyMenu music to some MusicManager struct
	scoreTextFace    *text.GoTextFace
	txtWeekTitleFace *text.GoTextFace
	yellowBG         *ebiten.Image
	blackBar         *ebiten.Image
	grpWeekText      *engine.Group[*storymenu.MenuItem]
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
		return err
	}

	s.txtWeekTitleFace = s.scoreTextFace // they're the same font and size

	s.yellowBG = ebiten.NewImage(engine.GameWidth, 400)
	s.yellowBG.Fill(color.NRGBA{0xF9, 0xCF, 0x51, 0xFF})

	s.blackBar = ebiten.NewImage(engine.GameWidth, 56)
	s.blackBar.Fill(color.Black)

	s.grpWeekText = engine.NewGroup[*storymenu.MenuItem]()

	for i, week := range s.ctx.Weeks {
		item, err := storymenu.NewMenuItem(s.ctx, 0, s.blackBar.Bounds().Dy()+s.yellowBG.Bounds().Dy()+10, i, week)
		if err != nil {
			return fmt.Errorf("StoryModeScene: failed to load menu item for week %s: %w", week.ID, err)
		}
		item.Sprite.Position.Y += (item.Bounds.Dy() * i)
		item.Sprite.Position.X = (engine.GameWidth - item.Bounds.Dx()) / 2 // centers image on x axis
		s.grpWeekText.Add(item)
	}

	return nil
}

func (s *StoryMenuScene) Close() error {
	return nil
}

func (s *StoryMenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	screen.DrawImage(s.yellowBG, nil)
	screen.DrawImage(s.blackBar, nil)

	{
		op := &text.DrawOptions{}
		op.GeoM.Translate(10, 10)
		text.Draw(screen, "SCORE: 49324858", s.scoreTextFace, op)
	}
	{
		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(engine.GameWidth)*0.7, 10)
		op.ColorScale.ScaleAlpha(0.7)
		op.PrimaryAlign = text.AlignStart
		text.Draw(screen, "test week name", s.txtWeekTitleFace, op)
	}

	s.grpWeekText.Draw(screen)
}

func (s *StoryMenuScene) Update() error {
	if s.ctx.InputHandler.ActionIsJustPressed(controls.ActionBack) {
		if err := audioutil.PlaySoundFromFS(s.ctx, s.ctx.AssetsFS, "sounds/cancelMenu.ogg", 0); err != nil {
			return err
		}
		return s.ctx.SceneCtrl.SwitchScene(&MainMenuScene{ctx: s.ctx})
	}

	s.grpWeekText.Update()

	return nil
}
