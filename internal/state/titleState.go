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
	"github.com/rs/zerolog/log"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/ztrue/tracerr"
)

type TitleState struct {
	mb              *ge.MusicBeat
	inited          bool
	logoBl          *ge.Sprite
	gfDance         *ge.Sprite
	freakyMenu      *audio.Player
	freakyMenuTween *gween.Tween
	danceLeft       bool
}

func NewTitleState() (*TitleState, error) {
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

	return &TitleState{
		mb:              mb,
		inited:          false,
		logoBl:          logoBl,
		freakyMenu:      freakyMenu,
		freakyMenuTween: gween.New(0, 0.7, 4, ease.Linear),
		danceLeft:       false,
	}, nil
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
}

func titleState_BeatHit(beatHit int) {
	log.Info().Int("beatHit", beatHit).Msg("BeatHit")
}
