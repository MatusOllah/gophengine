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
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

var titleState *TitleState

type TitleState struct {
	ng              *ge.Sprite
	mb              *ge.MusicBeat
	once            *sync.Once
	text            []string
	introText       []string
	logoBl          *ge.Sprite
	gfDance         *ge.Sprite
	freakyMenu      *audio.Player
	freakyMenuTween *gween.Tween
	danceLeft       bool
	flasher         *ge.Flasher
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
		introText:       it,
		ng:              ng,
		mb:              mb,
		once:            new(sync.Once),
		logoBl:          logoBl,
		freakyMenu:      freakyMenu,
		freakyMenuTween: gween.New(0, 0.7, 4, ease.Linear),
		danceLeft:       false,
		flasher:         flasher,
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

	ge.G.Conductor.SongPosition = float64(s.freakyMenu.Position().Milliseconds())

	freakyMenuVolume, _ := s.freakyMenuTween.Update(float32(dt))
	s.freakyMenu.SetVolume(float64(freakyMenuVolume))

	s.mb.Update()

	s.flasher.Update(dt)

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	s.drawText(screen)
	s.ng.Draw(screen)

	s.flasher.Draw(screen)
}

func (ts *TitleState) skipIntro() {
	slog.Info("skipIntro")
	//TODO: skip intro
	ts.ng.Img.Dispose()
	ts.flasher.Flash()
}

func (ts *TitleState) drawText(img *ebiten.Image) {
	for i, s := range ts.text {
		ebitenutil.DebugPrintAt(img, s, img.Bounds().Dx()/2, (i*60)+200)
	}
}

func (s *TitleState) createText(text ...string) {
	s.text = append(s.text, text...)
}

func (s *TitleState) addText(text string) {
	s.text = append(s.text, text)
}

func (s *TitleState) deleteText() {
	s.text = nil
}

func titleState_BeatHit(curBeat int) {
	switch curBeat {
	case 1:
		titleState.createText(
			"ninjamuffin99",
			"phantomArcade",
			"kawaisprite",
			"evilsk8er",
			"MatusOllah",
		)
	case 3:
		titleState.addText("present")
	case 4:
		titleState.deleteText()
	case 5:
		titleState.createText("In association", "with")
	case 7:
		titleState.addText("newgrounds")
		titleState.ng.Visible = true
	case 8:
		titleState.ng.Visible = false
		titleState.deleteText()
	case 9:
		titleState.createText(titleState.introText[0])
	case 11:
		titleState.addText(titleState.introText[1])
	case 12:
		titleState.deleteText()
	case 13:
		titleState.addText("Friday")
	case 14:
		titleState.addText("Night")
	case 15:
		titleState.addText("Funkin")
	case 16:
		titleState.deleteText()
		titleState.skipIntro()
	}
}
