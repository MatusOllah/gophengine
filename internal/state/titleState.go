package state

import (
	"encoding/csv"
	"image/color"
	_ "image/png"
	"log/slog"
	"sync"
	"time"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

var titleState *TitleState

// TitleState is the intro and "press enter to begin" screen.
type TitleState struct {
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
}

func getRandIntroText() ([]string, error) {
	f, err := assets.FS.Open("data/introText.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	introText := records[ge.G.Rand.IntN(len(records))]
	slog.Info("got intro text", "introText", introText)

	return introText, nil
}

func NewTitleState() (*TitleState, error) {
	it, err := getRandIntroText()
	if err != nil {
		return nil, err
	}

	ng := ge.NewSprite((float64(ge.G.Width)/2)-150, float64(ge.G.Height)*0.52)
	ngImg, _, err := ebitenutil.NewImageFromFileSystem(assets.FS, "images/newgrounds_logo.png")
	if err != nil {
		return nil, err
	}
	ng.Img = ngImg
	ng.Visible = false

	logoBl := ge.NewSprite(-150, -100)
	ac, err := animhcl.LoadAnimsFromFS(assets.FS, "images/logoBumpin/logoBumpin.anim.hcl")
	if err != nil {
		return nil, err
	}
	logoBl.AnimController = ac

	gfDance := ge.NewSprite(float64(ge.G.Width)*0.4, float64(ge.G.Height)*0.07)
	ac, err = animhcl.LoadAnimsFromFS(assets.FS, "images/gfDanceTitle/gfDanceTitle.anim.hcl")
	if err != nil {
		return nil, err
	}
	gfDance.AnimController = ac

	titleText := ge.NewSprite(100, float64(ge.G.Height)*0.8)
	ac, err = animhcl.LoadAnimsFromFS(assets.FS, "images/titleEnter/titleEnter.anim.hcl")
	if err != nil {
		return nil, err
	}
	titleText.AnimController = ac

	freakyMenuFile, err := assets.FS.Open("music/freakyMenu.ogg")
	if err != nil {
		return nil, err
	}
	defer freakyMenuFile.Close()

	freakyMenuStreamer, freakyMenuFormat, err := vorbis.Decode(freakyMenuFile)
	if err != nil {
		return nil, err
	}

	freakyMenu := &effects.Volume{
		Streamer: ge.Resample(freakyMenuFormat.SampleRate, beep.Loop(-1, freakyMenuStreamer)),
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	mb := ge.NewMusicBeat()
	mb.BeatHitFunc = titleState_BeatHit

	flasher, err := ge.NewFlasher(1)
	if err != nil {
		return nil, err
	}

	ts := &TitleState{
		randIntroText:      it,
		introText:          ge.NewIntroText(),
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
	}

	titleState = ts

	return ts, nil
}

func (s *TitleState) Update(dt float64) error {
	s.once.Do(func() {
		slog.Info("(*sync.Once).Do")
		ge.G.Mixer.Add(s.freakyMenu)

		ge.G.Conductor.ChangeBPM(102)
	})

	// Conductor & MusicBeat (MusicBeatState)
	ge.G.Conductor.SongPosition = float64(s.freakyMenuFormat.SampleRate.D(s.freakyMenuStreamer.Position()).Milliseconds())
	s.mb.Update()

	// freakyMenu Volume
	freakyMenuVolume, _ := s.freakyMenuTween.Update(float32(dt))
	speaker.Lock()
	s.freakyMenu.Volume = float64(freakyMenuVolume)
	speaker.Unlock()

	// Title screen
	s.gfDance.AnimController.UpdateWithDelta(dt)
	s.logoBl.AnimController.UpdateWithDelta(dt)
	s.titleText.AnimController.UpdateWithDelta(dt)

	s.flasher.Update(dt)

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && !s.transitioning && s.skippedIntro {
		s.titleText.AnimController.Play("press")
		s.flasher.Flash()

		if err := ge.PlaySoundFromFS(assets.FS, "sounds/confirmMenu.ogg", -0.3); err != nil {
			return err
		}

		slog.Info("pressed enter, transitioning")
		s.transitioning = true

		time.AfterFunc(2*time.Second, func() {
			panic("main menu not implemented yet")
		})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && !s.skippedIntro {
		s.skipIntro()
	}

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	s.gfDance.AnimController.Draw(screen, s.gfDance.DrawImageOptions())
	s.logoBl.AnimController.Draw(screen, s.logoBl.DrawImageOptions())
	s.titleText.AnimController.Draw(screen, s.titleText.DrawImageOptions())

	if s.blackScreenVisible {
		screen.Fill(color.Black)
	}

	if !s.skippedIntro {
		s.introText.Draw(screen)
		s.ng.Draw(screen)
	}

	s.flasher.Draw(screen)
}

func (s *TitleState) skipIntro() {
	if s.skippedIntro {
		return
	}
	s.skippedIntro = true
	slog.Info("skipIntro")
	s.blackScreenVisible = false
	s.ng.Img.Dispose()
	s.flasher.Flash()
}

func titleState_BeatHit(curBeat int) {
	titleState.logoBl.AnimController.Play("bump")
	titleState.danceLeft = !titleState.danceLeft

	if titleState.danceLeft {
		titleState.gfDance.AnimController.Play("danceRight")
	} else {
		titleState.gfDance.AnimController.Play("danceLeft")
	}

	switch curBeat {
	case 1:
		titleState.introText.CreateText(
			"ninjamuffin99",
			"phantomArcade",
			"kawaisprite",
			"evilsk8er",
			"MatusOllah",
		)
	case 3:
		titleState.introText.AddText("present")
	case 4:
		titleState.introText.DeleteText()
	case 5:
		titleState.introText.CreateText("In association", "with")
	case 7:
		titleState.introText.AddText("newgrounds")
		titleState.ng.Visible = true
	case 8:
		titleState.ng.Visible = false
		titleState.introText.DeleteText()
	case 9:
		titleState.introText.CreateText(titleState.randIntroText[0])
	case 11:
		titleState.introText.AddText(titleState.randIntroText[1])
	case 12:
		titleState.introText.DeleteText()
	case 13:
		titleState.introText.AddText("Friday")
	case 14:
		titleState.introText.AddText("Night")
	case 15:
		titleState.introText.AddText("Funkin")
	case 16:
		titleState.introText.DeleteText()
		titleState.skipIntro()
	}
}
