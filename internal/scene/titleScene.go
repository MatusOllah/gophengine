package scene

import (
	"encoding/csv"
	"image/color"
	_ "image/png"
	"io/fs"
	"log/slog"
	"sync"
	"time"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/MatusOllah/gophengine/internal/audioutil"
	"github.com/MatusOllah/gophengine/internal/controls"
	gefx "github.com/MatusOllah/gophengine/internal/effects"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/funkin"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/internal/scene/title"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/vorbis"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
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
	freakyMenuStreamer beep.StreamSeekCloser
	freakyMenuFormat   beep.Format
	freakyMenu         *effects.Volume
	freakyMenuTween    *gween.Tween
	freakyMenuBPM      int
	danceLeft          bool
	flasher            *gefx.Flasher
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

	introText := records[ctx.Rand.IntN(len(records))]
	slog.Info("got intro text", "introText", introText)

	return introText, nil
}

func getFreakyMenuMeta(ctx *context.Context) (bpm int, tween *gween.Tween, err error) {
	path := "music/freakyMenu.meta.hcl"
	b, err := fs.ReadFile(ctx.AssetsFS, path)
	if err != nil {
		return
	}

	var v struct {
		BPM   int `hcl:"BPM"`
		Tween struct {
			Begin    float32 `hcl:"Begin"`
			End      float32 `hcl:"End"`
			Duration float32 `hcl:"Duration"`
		} `hcl:"Tween,block"`
	}
	if err = hclsimple.Decode(path, b, nil, &v); err != nil {
		return
	}

	return v.BPM, gween.New(v.Tween.Begin, v.Tween.End, v.Tween.Duration, ease.Linear), nil
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

	s.goLogo = engine.NewSprite(int((float64(s.ctx.Width)/2)-300), int(float64(s.ctx.Height)*0.33))
	goLogoImg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/go_logo.png")
	if err != nil {
		return err
	}
	s.goLogo.Img = goLogoImg
	s.goLogo.Visible = false

	s.ebitenLogo = engine.NewSprite(int((float64(s.ctx.Width) / 2)), int(float64(s.ctx.Height)*0.33))
	ebitenLogoImg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/ebiten_logo.png")
	if err != nil {
		return err
	}
	s.ebitenLogo.Img = ebitenLogoImg
	s.ebitenLogo.Visible = false

	logoBl := engine.NewSprite(-150, -100)
	ac, err := animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/logoBumpin/logoBumpin.anim.hcl", "logoBumpin")
	if err != nil {
		return err
	}
	logoBl.AnimController = ac
	s.logoBl = logoBl

	gfDance := engine.NewSprite(int(float64(s.ctx.Width)*0.4), int(float64(s.ctx.Height)*0.07))
	ac, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/gfDanceTitle/gfDanceTitle.anim.hcl", "gfDanceTitle")
	if err != nil {
		return err
	}
	gfDance.AnimController = ac
	s.gfDance = gfDance

	titleText := engine.NewSprite(100, int(float64(s.ctx.Height)*0.8))
	ac, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/titleEnter/titleEnter.anim.hcl", "titleEnter")
	if err != nil {
		return err
	}
	titleText.AnimController = ac
	s.titleText = titleText

	// FreakyMenu
	freakyMenuFile, err := s.ctx.AssetsFS.Open("music/freakyMenu.ogg")
	if err != nil {
		return err
	}
	defer freakyMenuFile.Close()

	s.freakyMenuStreamer, s.freakyMenuFormat, err = vorbis.Decode(freakyMenuFile)
	if err != nil {
		return err
	}

	s.freakyMenu = &effects.Volume{
		Streamer: audio.MustLoop2(s.freakyMenuStreamer),
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	s.freakyMenuBPM, s.freakyMenuTween, err = getFreakyMenuMeta(s.ctx)
	if err != nil {
		return err
	}

	mb := funkin.NewMusicBeat(s.ctx.Conductor)
	mb.BeatHitFunc = titleState_BeatHit
	s.mb = mb

	s.flasher = gefx.NewFlasher(s.ctx.Width, s.ctx.Height, 1)

	introText, err := title.NewIntroText(s.ctx.AssetsFS)
	if err != nil {
		return err
	}
	s.introText = introText

	s.once = &sync.Once{}
	s.danceLeft = false
	s.blackScreenVisible = true
	s.skippedIntro = false
	s.transitioning = false
	s.errCh = make(chan error, 1)

	titleSceneInstance = s

	return nil
}

func (s *TitleScene) Close() error {
	return nil
}

func (s *TitleScene) Update(dt float64) error {
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
		s.ctx.AudioMixer.Music.Add(s.freakyMenu)

		s.ctx.Conductor.ChangeBPM(s.freakyMenuBPM)
	})

	// Conductor & MusicBeat (MusicBeatState)
	s.ctx.Conductor.SongPosition = float64(s.freakyMenuFormat.SampleRate.D(s.freakyMenuStreamer.Position()).Milliseconds())
	s.mb.Update()

	// freakyMenu Volume
	freakyMenuVolume, _ := s.freakyMenuTween.Update(float32(dt))
	speaker.Lock()
	s.freakyMenu.Volume = float64(freakyMenuVolume)
	speaker.Unlock()

	// Title screen
	if s.skippedIntro {
		s.gfDance.AnimController.Update()
		s.logoBl.AnimController.Update()
		s.titleText.AnimController.Update()
	}

	s.flasher.Update(dt)

	if s.ctx.InputHandler.ActionIsJustPressed(controls.ActionAccept) && !s.transitioning && s.skippedIntro {
		s.titleText.AnimController.Play("press")
		s.flasher.Flash()

		if err := audioutil.PlaySoundFromFS(s.ctx, s.ctx.AssetsFS, "sounds/confirmMenu.ogg", -0.3); err != nil {
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
