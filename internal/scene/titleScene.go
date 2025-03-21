package scene

import (
	"encoding/csv"
	"image/color"
	_ "image/png"
	"log/slog"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/MatusOllah/gophengine/internal/audio/music"
	"github.com/MatusOllah/gophengine/internal/controls"
	"github.com/MatusOllah/gophengine/internal/effects"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/funkin"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/internal/scene/title"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var titleSceneInstance *TitleScene

var _ engine.Scene = (*TitleScene)(nil)

// TitleScene is the intro and "press enter to begin" screen.
type TitleScene struct {
	ctx                *context.Context
	goLogo             *engine.Sprite
	ebitenLogo         *engine.Sprite
	mb                 *funkin.MusicBeat
	once               *sync.Once
	randIntroText      []string
	introText          *title.IntroText
	logoBl             *engine.Sprite
	gfDance            *engine.Sprite
	titleText          *engine.Sprite
	danceLeft          bool
	flasher            *effects.Flasher
	blackScreenVisible bool
	skippedIntro       bool
	transitioning      bool
	errCh              chan error
}

func getRandIntroText(ctx *context.Context) ([]string, error) {
	f, err := ctx.AssetsFS.Open("data/introText.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comment = '#'

	_, err = r.Read()
	if err != nil {
		return nil, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	introText := records[rand.IntN(len(records))]
	slog.Info("got intro text", "introText", introText)

	return introText, nil
}

func NewTitleScene(ctx *context.Context) *TitleScene {
	return &TitleScene{ctx: ctx}
}

func (s *TitleScene) Init() error {
	it, err := getRandIntroText(s.ctx)
	if err != nil {
		return err
	}
	s.randIntroText = it

	s.goLogo = engine.NewSprite(int((float64(engine.GameWidth)/2)-300), int(float64(engine.GameHeight)*0.33))
	goLogoImg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/go_logo.png")
	if err != nil {
		return err
	}
	s.goLogo.Img = goLogoImg
	s.goLogo.Visible = false

	s.ebitenLogo = engine.NewSprite(int((float64(engine.GameWidth) / 2)), int(float64(engine.GameHeight)*0.33))
	ebitenLogoImg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/ebiten_logo.png")
	if err != nil {
		return err
	}
	s.ebitenLogo.Img = ebitenLogoImg
	s.ebitenLogo.Visible = false

	logoBl := engine.NewSprite(-150, -100)
	ac, err := animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/logoBumpin/anim.hcl", "logoBumpin")
	if err != nil {
		return err
	}
	logoBl.AnimController = ac
	s.logoBl = logoBl

	gfDance := engine.NewSprite(int(float64(engine.GameWidth)*0.4), int(float64(engine.GameHeight)*0.07))
	ac, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/gfDanceTitle/anim.hcl", "gfDanceTitle")
	if err != nil {
		return err
	}
	gfDance.AnimController = ac
	s.gfDance = gfDance

	titleText := engine.NewSprite(100, int(float64(engine.GameHeight)*0.8))
	ac, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/titleEnter/anim.hcl", "titleEnter")
	if err != nil {
		return err
	}
	titleText.AnimController = ac
	s.titleText = titleText

	// FreakyMenu
	s.ctx.FreakyMenu, err = music.New(s.ctx.AssetsFS)
	if err != nil {
		return err
	}

	mb := funkin.NewMusicBeat(s.ctx.Conductor)
	mb.BeatHitFunc = titleState_BeatHit
	s.mb = mb

	s.flasher = effects.NewFlasher(engine.GameWidth, engine.GameHeight, 1)

	introText, err := title.NewIntroText(s.ctx.AssetsFS)
	if err != nil {
		return err
	}
	s.introText = introText

	s.once = &sync.Once{}
	s.blackScreenVisible = true
	s.errCh = make(chan error, 1)

	titleSceneInstance = s

	return nil
}

func (s *TitleScene) Close() error {
	return nil
}

func (s *TitleScene) Update() error {
	select {
	case err := <-s.errCh:
		if err != nil {
			return err
		}
	default:
		// Continue with update routine
	}

	s.once.Do(func() {
		slog.Debug("s.once.Do")
		s.ctx.AudioMixer.Music.Add(s.ctx.FreakyMenu)

		s.ctx.Conductor.ChangeBPM(s.ctx.FreakyMenu.BPM())
	})

	// Conductor & MusicBeat (MusicBeatState)
	s.ctx.Conductor.SongPosition = float64(s.ctx.FreakyMenu.Format().SampleRate.D(s.ctx.FreakyMenu.Position()).Milliseconds())
	s.mb.Update()

	// freakyMenu Volume
	s.ctx.FreakyMenu.Update()

	// Title screen
	if s.skippedIntro {
		s.gfDance.AnimController.Update()
		s.logoBl.AnimController.Update()
		s.titleText.AnimController.Update()
	}

	s.flasher.Update()

	if s.ctx.InputHandler.ActionIsJustPressed(controls.ActionAccept) && !s.transitioning && s.skippedIntro {
		s.titleText.AnimController.Play("press")
		s.flasher.Flash()

		if err := audio.PlaySoundFromFS(s.ctx.AssetsFS, "sounds/confirmMenu.ogg", -0.3, s.ctx.AudioMixer.SFX); err != nil {
			return err
		}

		slog.Info("pressed enter, transitioning")
		s.transitioning = true

		// Bit janky solution with the error channel but whatever
		time.AfterFunc(2*time.Second, func() {
			titleSceneInstance.errCh <- titleSceneInstance.ctx.SceneCtrl.SwitchScene(NewMainMenuScene(s.ctx))
		})
	}

	if s.ctx.InputHandler.ActionIsJustPressed(controls.ActionAccept) && !s.skippedIntro {
		s.skipIntro()
	}

	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	if s.skippedIntro {
		s.gfDance.AnimController.Draw(screen, s.gfDance.Position)
		s.logoBl.AnimController.Draw(screen, s.logoBl.Position)
		s.titleText.AnimController.Draw(screen, s.titleText.Position)
	}

	if s.blackScreenVisible {
		screen.Fill(color.Black)
	}

	if !s.skippedIntro {
		s.introText.Draw(screen)
		s.goLogo.Draw(screen)
		s.ebitenLogo.Draw(screen)
	}

	s.flasher.Draw(screen)
}

func (s *TitleScene) skipIntro() {
	if s.skippedIntro {
		return
	}
	s.skippedIntro = true
	slog.Info("skipIntro")
	s.blackScreenVisible = false
	s.goLogo.Img.Deallocate()
	s.ebitenLogo.Img.Deallocate()
	s.flasher.Flash()
}

func titleState_BeatHit(curBeat int) {
	titleSceneInstance.logoBl.AnimController.Play("bump")
	titleSceneInstance.danceLeft = !titleSceneInstance.danceLeft

	if titleSceneInstance.danceLeft {
		titleSceneInstance.gfDance.AnimController.Play("danceRight")
	} else {
		titleSceneInstance.gfDance.AnimController.Play("danceLeft")
	}

	switch curBeat {
	case 1:
		titleSceneInstance.introText.CreateText(
			"ninjamuffin99",
			"phantomArcade",
			"kawaisprite",
			"evilsk8er",
			"MatusOllah",
		)
	case 3:
		titleSceneInstance.introText.AddText(i18n.L("Present"))
	case 4:
		titleSceneInstance.introText.DeleteText()
	case 5:
		titleSceneInstance.introText.CreateText(i18n.L("MadeWith"))
	case 7:
		titleSceneInstance.introText.AddText("")
		titleSceneInstance.introText.AddText("+")
		titleSceneInstance.goLogo.Visible = true
		titleSceneInstance.ebitenLogo.Visible = true
	case 8:
		titleSceneInstance.goLogo.Visible = false
		titleSceneInstance.ebitenLogo.Visible = false
		titleSceneInstance.introText.DeleteText()
	case 9:
		titleSceneInstance.introText.CreateText(titleSceneInstance.randIntroText[0])
	case 11:
		titleSceneInstance.introText.AddText(titleSceneInstance.randIntroText[1])
	case 12:
		titleSceneInstance.introText.DeleteText()
	case 13:
		titleSceneInstance.introText.AddText("Friday")
	case 14:
		titleSceneInstance.introText.AddText("Night")
	case 15:
		titleSceneInstance.introText.AddText("Funkin")
	case 16:
		titleSceneInstance.introText.DeleteText()
		titleSceneInstance.skipIntro()
	}
}
