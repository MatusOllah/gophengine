package scene

import (
	"fmt"
	"image/color"
	"log/slog"
	"strings"

	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/MatusOllah/gophengine/internal/controls"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/scene/storymenu"
	"github.com/MatusOllah/gophengine/pkg/context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StoryMenuScene struct {
	ctx              *context.Context
	scoreTextFace    *text.GoTextFace
	txtWeekTitleFace *text.GoTextFace
	curWeek          int
	txtTracklistFace *text.GoTextFace
	//TODO: menu characters
	movedBack    bool
	selectedWeek bool
	yellowBG     *ebiten.Image
	blackBar     *ebiten.Image
	grpWeekText  *engine.Group[*storymenu.MenuItem]
	leftArrow    *engine.Sprite
	difficulty   *engine.Sprite
	rightArrow   *engine.Sprite
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

	i := 0
	for _, week := range s.ctx.Weeks.Iter() {
		item, err := storymenu.NewMenuItem(s.ctx, 0, s.blackBar.Bounds().Dy()+s.yellowBG.Bounds().Dy()+10, i, week)
		if err != nil {
			return fmt.Errorf("StoryModeScene: failed to load menu item for week %s: %w", week.ID, err)
		}
		item.Sprite.Position.Y += (item.Bounds.Dy() * i)
		item.Sprite.Position.X = (engine.GameWidth - item.Bounds.Dx()) / 2 // centers image on x axis
		s.grpWeekText.Add(item)
		i++
	}

	s.leftArrow = engine.NewSprite(s.grpWeekText.Get(0).Sprite.Position.X+s.grpWeekText.Get(0).Bounds.Dx()+10, s.grpWeekText.Get(0).Sprite.Position.Y+10)
	s.leftArrow.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/storymenu/ui/anim.hcl", "left_arrow")
	if err != nil {
		return err
	}

	s.difficulty = engine.NewSprite(s.leftArrow.Position.X+60, s.leftArrow.Position.Y)
	s.difficulty.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/storymenu/ui/anim.hcl", "difficulty")
	if err != nil {
		return err
	}

	s.rightArrow = engine.NewSprite(s.difficulty.Position.X+320, s.leftArrow.Position.Y)
	s.rightArrow.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/storymenu/ui/anim.hcl", "right_arrow")
	if err != nil {
		return err
	}

	return nil
}

func (s *StoryMenuScene) Close() error {
	return nil
}

func (s *StoryMenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	s.grpWeekText.Draw(screen)

	s.leftArrow.AnimController.Draw(screen, s.leftArrow.Position)
	s.difficulty.AnimController.Draw(screen, s.difficulty.Position)
	s.rightArrow.AnimController.Draw(screen, s.rightArrow.Position)

	screen.DrawImage(s.yellowBG, nil)
	screen.DrawImage(s.blackBar, nil)

	{
		op := &text.DrawOptions{}
		op.GeoM.Translate(10, 10)
		text.Draw(screen, "SCORE: 49324858", s.scoreTextFace, op)
	}
	{
		week, _ := s.ctx.Weeks.GetIndex(s.curWeek)
		txt := strings.ToUpper(week.Name)

		width, _ := text.Measure(txt, s.txtWeekTitleFace, 0)

		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(engine.GameWidth)-(width+10), 10)
		op.ColorScale.ScaleAlpha(0.7)
		op.PrimaryAlign = text.AlignStart

		text.Draw(screen, txt, s.txtWeekTitleFace, op)
	}
}

func (s *StoryMenuScene) Update() error {
	s.grpWeekText.Update()

	s.leftArrow.AnimController.Update()
	s.difficulty.AnimController.Update()
	s.rightArrow.AnimController.Update()

	if !s.movedBack {
		if !s.selectedWeek {
			if s.ctx.InputHandler.ActionIsJustPressed(controls.ActionUp) {
				if err := s.changeWeek(-1); err != nil {
					return err
				}
			}
			if s.ctx.InputHandler.ActionIsJustPressed(controls.ActionDown) {
				if err := s.changeWeek(1); err != nil {
					return err
				}
			}
		}
	}

	if s.ctx.InputHandler.ActionIsJustPressed(controls.ActionBack) && !s.movedBack && !s.selectedWeek {
		if err := audio.PlaySoundFromFS(s.ctx.AssetsFS, "sounds/cancelMenu.ogg", 0, s.ctx.AudioMixer.SFX); err != nil {
			return err
		}
		s.movedBack = true
		return s.ctx.SceneCtrl.SwitchScene(&MainMenuScene{ctx: s.ctx})
	}

	return nil
}

func (s *StoryMenuScene) changeWeek(delta int) error {
	s.curWeek += delta

	if s.curWeek >= s.ctx.Weeks.Len() {
		s.curWeek = 0
	}
	if s.curWeek < 0 {
		s.curWeek = s.ctx.Weeks.Len() - 1
	}

	slog.Info("selected week", "i", s.curWeek)

	if err := audio.PlaySoundFromFS(s.ctx.AssetsFS, "sounds/scrollMenu.ogg", 0, s.ctx.AudioMixer.SFX); err != nil {
		return err
	}

	for i, item := range s.grpWeekText.Iterate() {
		item.TargetY = float64(i - s.curWeek)
	}

	s.updateText()

	return nil
}

func (s *StoryMenuScene) updateText() {
	//TODO: update track list and characters
}
