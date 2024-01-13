package state

import (
	"bytes"
	"encoding/csv"
	"image/color"
	_ "image/png"
	"io/fs"
	"log/slog"
	"sync"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/anim"
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

var titleState *TitleState

type TitleState struct {
	ng                 *ge.Sprite
	mb                 *ge.MusicBeat
	once               *sync.Once
	randIntroText      []string
	introText          *ge.IntroText
	logoBl             *ge.Sprite
	gfDance            *ge.Sprite
	titleText          *ge.Sprite
	freakyMenu         *audio.Player
	freakyMenuTween    *gween.Tween
	danceLeft          bool
	flasher            *ge.Flasher
	blackScreenVisible bool
	skippedIntro       bool
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

	introText := records[ge.G.Rand.Intn(len(records))]
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
	logoBl.AnimController.SetAnim("bump", anim.NewAnimation(anim.MustGetImagesByPrefixFromFS(assets.FS, "images/logoBumpin", "logo bumpin"), anim.Dur24FPS))

	gfDance := ge.NewSprite(float64(ge.G.Width)*0.4, float64(ge.G.Height)*0.07)
	gfDance.AnimController.SetAnim("danceLeft", anim.NewAnimation(anim.MustGetImagesByIndicesFromFS(assets.FS, "images/gfDanceTitle", "gfDance", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}), anim.Dur24FPS))
	gfDance.AnimController.SetAnim("danceRight", anim.NewAnimation(anim.MustGetImagesByIndicesFromFS(assets.FS, "images/gfDanceTitle", "gfDance", []int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}), anim.Dur24FPS))

	titleText := ge.NewSprite(100, float64(ge.G.Height)*0.8)
	titleText.AnimController.SetAnim("idle", anim.NewAnimation(anim.MustGetImagesByPrefixFromFS(assets.FS, "images/titleEnter", "Press Enter to Begin"), anim.Dur24FPS))
	titleText.AnimController.SetAnim("press", anim.NewAnimation(anim.MustGetImagesByPrefixFromFS(assets.FS, "images/titleEnter", "ENTER PRESSED"), anim.Dur24FPS))
	titleText.AnimController.Play("idle")

	freakyMenuContent, err := fs.ReadFile(assets.FS, "music/freakyMenu.ogg")
	if err != nil {
		return nil, err
	}

	freakyMenuStream, err := vorbis.DecodeWithSampleRate(48000, bytes.NewReader(freakyMenuContent))
	if err != nil {
		return nil, err
	}

	freakyMenu, err := audio.CurrentContext().NewPlayer(audio.NewInfiniteLoop(freakyMenuStream, freakyMenuStream.Length()))
	if err != nil {
		return nil, err
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
		freakyMenu:         freakyMenu,
		freakyMenuTween:    gween.New(0, 0.7, 4, ease.Linear),
		danceLeft:          false,
		flasher:            flasher,
		blackScreenVisible: true,
		skippedIntro:       false,
	}

	titleState = ts

	return ts, nil
}

func (s *TitleState) Update(dt float64) error {
	s.once.Do(func() {
		slog.Info("(*sync.Once).Do")
		s.freakyMenu.Play()

		ge.G.Conductor.ChangeBPM(102)
	})

	// Conductor & MusicBeat (MusicBeatState)
	ge.G.Conductor.SongPosition = float64(s.freakyMenu.Position().Milliseconds())
	s.mb.Update()

	// freakyMenu Volume
	freakyMenuVolume, _ := s.freakyMenuTween.Update(float32(dt))
	s.freakyMenu.SetVolume(float64(freakyMenuVolume))

	// Title screen
	//TODO: press enter to begin screen
	s.titleText.AnimController.Update(dt)

	s.flasher.Update(dt)

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && !s.skippedIntro {
		s.skipIntro()
	}

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	//TODO: press enter to begin screen
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

func (ts *TitleState) skipIntro() {
	slog.Info("skipIntro")
	ts.skippedIntro = true
	ts.blackScreenVisible = false
	ts.ng.Img.Dispose()
	ts.flasher.Flash()
}

func titleState_BeatHit(curBeat int) {
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
