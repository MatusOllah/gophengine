package main

import (
	"bytes"
	"image/color"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/yohamta/ganim8/v2"
	"github.com/ztrue/tracerr"
)

type TitleState struct {
	*MusicBeatState
	Inited    bool
	LogoBl    *ganim8.Animation
	GfDance   *ganim8.Animation
	DanceLeft bool
	//TitleText text.Text
}

func NewTitleState() (*TitleState, error) {
	/*
		logoBlImg, _, err := ebitenutil.NewImageFromFile(logoBumpinPngPath)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		logoBlAtlasContent, err := os.ReadFile(logoBumpinXmlPath)

		logoBlAtlas, err := ParseAtlas(logoBlAtlasContent)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		var logoBlRects []*image.Rectangle

		for _, st := range logoBlAtlas.SubTextures {
			logoBlRects = append(logoBlRects, &image.Rectangle{
				Min: image.Pt(st.X, st.Y),
				Max: image.Pt(st.Width, st.Height),
			})
		}

		logoBl := ganim8.New(logoBlImg, logoBlRects, nil, ganim8.Nop)
	*/

	return &TitleState{
		Inited: false,
		//LogoBl:    logoBl,
		DanceLeft: false,
	}, nil
}

func (s *TitleState) Update(dt float64) error {
	if !s.Inited {
		content, err := os.ReadFile(GetAsset(filepath.Join("music", "freakyMenu.ogg")))
		if err != nil {
			return tracerr.Wrap(err)
		}

		stream, err := vorbis.Decode(g.AudioContext, bytes.NewReader(content))
		if err != nil {
			return tracerr.Wrap(err)
		}

		player, err := g.AudioContext.NewPlayer(audio.NewInfiniteLoop(stream, stream.Length()))
		if err != nil {
			return tracerr.Wrap(err)
		}

		player.Play()

		conductor.ChangeBPM(102)
	}
	s.Inited = true

	return nil
}

func (s *TitleState) Draw(screen *ebiten.Image) {
	if !s.Inited {
		screen.Fill(color.Black)
	}
}
