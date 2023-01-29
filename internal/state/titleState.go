package state

import (
	"bytes"
	"image/color"
	_ "image/png"
	"os"
	"path/filepath"

	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ztrue/tracerr"
)

type TitleState struct {
	*MusicBeatState
	inited    bool
	logoBl    *ge.Sprite
	gfDance   *ge.Sprite
	danceLeft bool
	//TitleText text.Text
}

func NewTitleState() (*TitleState, error) {
	logoBl := ge.NewSprite(-150, -100)
	logoBlImg, _, err := ebitenutil.NewImageFromFile(ge.GetAsset("images/blavacik.png"))
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	logoBl.Image = logoBlImg

	return &TitleState{
		inited:    false,
		logoBl:    logoBl,
		danceLeft: false,
	}, nil
}

func (s *TitleState) Update(dt float64) error {
	if !s.inited {
		content, err := os.ReadFile(ge.GetAsset(filepath.Join("music", "freakyMenu.ogg")))
		if err != nil {
			return tracerr.Wrap(err)
		}

		stream, err := vorbis.Decode(ge.G.AudioContext, bytes.NewReader(content))
		if err != nil {
			return tracerr.Wrap(err)
		}

		player, err := ge.G.AudioContext.NewPlayer(audio.NewInfiniteLoop(stream, stream.Length()))
		if err != nil {
			return tracerr.Wrap(err)
		}

		player.Play()

		ge.G.Conductor.ChangeBPM(102)
	}
	s.inited = true

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	if !s.inited {
		screen.Fill(color.Black)
	}

	s.logoBl.Draw(screen)
}
