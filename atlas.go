package main

import (
	"encoding/xml"
	"image"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"github.com/ztrue/tracerr"
)

type TextureAtlas struct {
	XMLName     xml.Name     `xml:"TextureAtlas"`
	ImagePath   string       `xml:"imagePath,attr"`
	SubTextures []SubTexture `xml:"SubTexture"`
}

type SubTexture struct {
	Name        string `xml:"name,attr"`
	X           int    `xml:"x,attr"`
	Y           int    `xml:"y,attr"`
	Width       int    `xml:"width,attr"`
	Height      int    `xml:"height,attr"`
	FrameX      int    `xml:"frameX,attr"`
	FrameY      int    `xml:"frameY,attr"`
	FrameWidth  int    `xml:"frameWidth,attr"`
	FrameHeight int    `xml:"frameHeight,attr"`
}

func ParseAtlas(rawXML []byte) (*TextureAtlas, error) {
	type tempSubTexture struct {
		Name        string `xml:"name,attr"`
		X           string `xml:"x,attr"`
		Y           string `xml:"y,attr"`
		Width       string `xml:"width,attr"`
		Height      string `xml:"height,attr"`
		FrameX      string `xml:"frameX,attr"`
		FrameY      string `xml:"frameY,attr"`
		FrameWidth  string `xml:"frameWidth,attr"`
		FrameHeight string `xml:"frameHeight,attr"`
	}

	type tempTextureAtlas struct {
		XMLName     xml.Name         `xml:"TextureAtlas"`
		ImagePath   string           `xml:"imagePath,attr"`
		SubTextures []tempSubTexture `xml:"SubTexture"`
	}

	tempTa := tempTextureAtlas{}

	if err := xml.Unmarshal(rawXML, &tempTa); err != nil {
		return nil, tracerr.Wrap(err)
	}

	ta := TextureAtlas{
		XMLName:   tempTa.XMLName,
		ImagePath: tempTa.ImagePath,
	}

	for _, st := range tempTa.SubTextures {
		x, err := strconv.Atoi(st.X)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		y, err := strconv.Atoi(st.Y)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		width, err := strconv.Atoi(st.Width)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		height, err := strconv.Atoi(st.Height)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		frameX, err := strconv.Atoi(st.FrameX)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		frameY, err := strconv.Atoi(st.FrameY)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		frameWidth, err := strconv.Atoi(st.FrameWidth)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		frameHeight, err := strconv.Atoi(st.FrameHeight)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		ta.SubTextures = append(ta.SubTextures, SubTexture{
			Name:        st.Name,
			X:           x,
			Y:           y,
			Width:       width,
			Height:      height,
			FrameX:      frameX,
			FrameY:      frameY,
			FrameWidth:  frameWidth,
			FrameHeight: frameHeight,
		})
	}

	return &ta, nil
}

func (ta *TextureAtlas) SetupAnimations(img *ebiten.Image) (map[string]*ganim8.Animation, error) {
	anims := make(map[string]*ganim8.Animation)

	var sizes []*image.Rectangle

	for _, st := range ta.SubTextures {

		name := st.Name[:len(st.Name)-4]

		index, err := strconv.Atoi(st.Name[len(st.Name)-4:])
		if err != nil {
			return nil, err
		}

		rect := &image.Rectangle{
			Min: image.Pt(st.X, st.Y),
			Max: image.Pt(st.Width, st.Height),
		}

		size := &image.Rectangle{
			Min: image.Pt(st.FrameX, st.FrameY),
			Max: image.Pt(st.FrameWidth, st.FrameHeight),
		}
	}
}
