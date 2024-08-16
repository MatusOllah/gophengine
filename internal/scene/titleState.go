package scene

import (
	"encoding/csv"
	"image/color"
	_ "image/png"
	"log/slog"
	"sync"
	"time"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/MatusOllah/gophengine/internal/audioutil"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

var titleSceneInstance *TitleScene

var _ ge.State = (*TitleScene)(nil)

// TitleScene is the intro and "press enter to begin" screen.
type TitleScene struct {
	ctx                *context.Context
	ng                 *ge.Sprite
	mb                 *ge.MusicBeat
	once               *sync.Once
	randIntroText      []string
	introText          *ge.IntroText
	logoBl             *ge.Sprite
	gfDance            *ge.Sprite
	titleText          *ge.Sprite
	freakyMenuStreamer beep.StreamSeekCloser
	freakyMenuFormat   beep.Format
	freakyMenu         *effects.Volume
	freakyMenuTween    *gween.Tween
	danceLeft          bool
	flasher            *ge.Flasher
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

func NewTitleScene(ctx *context.Context) (*TitleScene, error) {
	it, err := getRandIntroText(ctx)
	if err != nil {
		return nil, err
	}

	ng := ge.NewSprite(int((float64(ctx.Width)/2)-150), int(float64(ctx.Height)*0.52))
	ngImg, _, err := ebitenutil.NewImageFromFileSystem(ctx.AssetsFS, "images/newgrounds_logo.png")
	if err != nil {
		return nil, err
	}
	ng.Img = ngImg
	ng.Visible = false

	logoBl := ge.NewSprite(-150, -100)
	ac, err := animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/logoBumpin/logoBumpin.anim.hcl", "logoBumpin")
	if err != nil {
		return nil, err
	}
	logoBl.AnimController = ac

	gfDance := ge.NewSprite(int(float64(ctx.Width)*0.4), int(float64(ctx.Height)*0.07))
	ac, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/gfDanceTitle/gfDanceTitle.anim.hcl", "gfDanceTitle")
	if err != nil {
		return nil, err
	}
	gfDance.AnimController = ac

	titleText := ge.NewSprite(100, int(float64(ctx.Height)*0.8))
	ac, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/titleEnter/titleEnter.anim.hcl", "titleEnter")
	if err != nil {
		return nil, err
	}
	titleText.AnimController = ac

	freakyMenuFile, err := ctx.AssetsFS.Open("music/freakyMenu.ogg")
	if err != nil {
		return nil, err
	}
	defer freakyMenuFile.Close()

	freakyMenuStreamer, freakyMenuFormat, err := vorbis.Decode(freakyMenuFile)
	if err != nil {
		return nil, err
	}

	freakyMenu := &effects.Volume{
		Streamer: beep.Resample(ctx.AudioResampleQuality, freakyMenuFormat.SampleRate, ctx.SampleRate, beep.Loop(-1, freakyMenuStreamer)),
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	mb := ge.NewMusicBeat(ctx.Conductor)
	mb.BeatHitFunc = titleState_BeatHit

	flasher, err := ge.NewFlasher(ctx.Width, ctx.Height, 1)
	if err != nil {
		return nil, err
	}

	introText, err := ge.NewIntroText(ctx.AssetsFS)
	if err != nil {
		return nil, err
	}

	ts := &TitleScene{
		ctx:                ctx,
		randIntroText:      it,
		introText:          introText,
		ng:                 ng,
		mb:                 mb,
		once:               new(sync.Once),
		logoBl:             logoBl,
		gfDance:            gfDance,
		titleText:          titleText,
		freakyMenuStreamer: freakyMenuStreamer,
		freakyMenuFormat:   freakyMenuFormat,
		freakyMenu:         freakyMenu,
		freakyMenuTween:    gween.New(-10, -0.3, 4, ease.Linear), // 0 => 0.7
		danceLeft:          false,
		flasher:            flasher,
		blackScreenVisible: true,
		skippedIntro:       false,
		transitioning:      false,
		errCh:              make(chan error, 1),
	}

	titleSceneInstance = ts

	return ts, nil
}

func (s *TitleScene) Update(dt float64) error {
	s.once.Do(func() {
		slog.Info("(*sync.Once).Do")
		s.ctx.AudioMixer.Add(s.freakyMenu)

		s.ctx.Conductor.ChangeBPM(102)
	})

	select {
	case err := <-s.errCh:
		if err != nil {
			return err
		}
	default:
		// Continue with update routine
	}

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
		s.gfDance.AnimController.UpdateWithDelta(dt)
		s.logoBl.AnimController.UpdateWithDelta(dt)
		s.titleText.AnimController.UpdateWithDelta(dt)
	}

	s.flasher.Update(dt)

	if s.ctx.InputHandler.ActionIsJustPressed(ge.ActionAccept) && !s.transitioning && s.skippedIntro {
		s.titleText.AnimController.Play("press")
		s.flasher.Flash()

		if err := audioutil.PlaySoundFromFS(s.ctx, s.ctx.AssetsFS, "sounds/confirmMenu.ogg", -0.3); err != nil {
			return err
		}

		slog.Info("pressed enter, transitioning")
		s.transitioning = true

		// Bit janky solution with the error channel but whatever
		time.AfterFunc(2*time.Second, func() {
			mms, err := NewMainMenuScene(s.ctx)
			if err != nil {
				titleSceneInstance.errCh <- err
				return
			}
			titleSceneInstance.ctx.StateController.SwitchState(mms)
			titleSceneInstance.errCh <- nil
		})
	}

	if s.ctx.InputHandler.ActionIsJustPressed(ge.ActionAccept) && !s.skippedIntro {
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
		s.ng.Draw(screen)
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
	s.ng.Img.Deallocate()
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
		titleSceneInstance.introText.AddText(i18nutil.Localize(titleSceneInstance.ctx.Localizer, "IntroTextPresent"))
	case 4:
		titleSceneInstance.introText.DeleteText()
	case 5:
		titleSceneInstance.introText.CreateText(i18nutil.Localize(titleSceneInstance.ctx.Localizer, "IntroTextInAssoc"), i18nutil.Localize(titleSceneInstance.ctx.Localizer, "IntroTextWith"))
	case 7:
		titleSceneInstance.introText.AddText("newgrounds")
		titleSceneInstance.ng.Visible = true
	case 8:
		titleSceneInstance.ng.Visible = false
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
