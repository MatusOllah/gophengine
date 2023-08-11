package state

import (
	"bytes"
	"image/color"
	_ "image/png"
	"io/fs"

	"github.com/MatusOllah/gophengine/assets"
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/ztrue/tracerr"
)

var titleState *TitleState

type TitleState struct {
	ng              *ge.Sprite
	drawNg          bool
	mb              *ge.MusicBeat
	inited          bool
	text            []string
	logoBl          *ge.Sprite
	gfDance         *ge.Sprite
	freakyMenu      *audio.Player
	freakyMenuTween *gween.Tween
	danceLeft       bool
}

func NewTitleState() (*TitleState, error) {
	ng := ge.NewSprite((float64(ge.G.ScreenWidth)/2)-150, float64(ge.G.ScreenHeight)*0.52)
	ngImg, _, err := ebitenutil.NewImageFromFileSystem(assets.FS, "images/newgrounds_logo.png")
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	ng.Img = ngImg

	logoBl := ge.NewSprite(-150, -100)
	logoBlImg, _, err := ebitenutil.NewImageFromFileSystem(assets.FS, "images/logoBumpin.png")
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	logoBl.Img = logoBlImg

	freakyMenuPath := "music/freakyMenu.ogg"
	if ge.Options.PC4R {
		freakyMenuPath = "music/freakyMenu_pc4r.ogg"
	}
	freakyMenuContent, err := fs.ReadFile(assets.FS, freakyMenuPath)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	freakyMenuStream, err := vorbis.DecodeWithSampleRate(48000, bytes.NewReader(freakyMenuContent))
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	freakyMenu, err := ge.G.AudioContext.NewPlayer(audio.NewInfiniteLoop(freakyMenuStream, freakyMenuStream.Length()))
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	mb := ge.NewMusicBeat()
	mb.BeatHitFunc = titleState_BeatHit

	ts := &TitleState{
		ng:              ng,
		drawNg:          false,
		mb:              mb,
		inited:          false,
		logoBl:          logoBl,
		freakyMenu:      freakyMenu,
		freakyMenuTween: gween.New(0, 0.7, 4, ease.Linear),
		danceLeft:       false,
	}

	titleState = ts

	return ts, nil
}

func (s *TitleState) Update(dt float64) error {
	if !s.inited {
		s.freakyMenu.Play()

		ge.G.Conductor.ChangeBPM(102)

		s.inited = true
	}

	ge.G.Conductor.SongPosition = float64(s.freakyMenu.Current().Milliseconds())

	freakyMenuVolume, _ := s.freakyMenuTween.Update(float32(dt))
	s.freakyMenu.SetVolume(float64(freakyMenuVolume))

	s.mb.Update()

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	s.drawText(screen)

	if s.drawNg {
		s.ng.Draw(screen)
	}
}

func (ts *TitleState) drawText(img *ebiten.Image) {
	for i, s := range ts.text {
		ebitenutil.DebugPrintAt(img, s, img.Bounds().Dx()/2, (i*60)+200)
	}
}

func (s *TitleState) createText(text []string) {
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
		titleState.createText([]string{
			"ninjamuffin99",
			"phantomArcade",
			"kawaisprite",
			"evilsk8er",
		})
	case 3:
		titleState.addText("present")
	case 4:
		titleState.deleteText()
	case 5:
		titleState.createText([]string{"In association", "with"})
	case 7:
		titleState.addText("newgrounds")
		titleState.drawNg = true
	case 8:
		titleState.drawNg = false
		titleState.deleteText()
	case 9:
		titleState.createText([]string{"horalky"})
	case 11:
		titleState.addText("sedita")
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
		//TODO: skip intro
	}
}
