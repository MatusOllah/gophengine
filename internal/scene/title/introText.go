package title

import (
	"image/color"
	"io/fs"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type IntroText struct {
	textLock sync.RWMutex
	text     []string

	fullFace    text.Face
	outlineFace text.Face
}

func NewIntroText(fsys fs.FS) (*IntroText, error) {
	fullFile, err := fsys.Open("fonts/phantom-full.ttf")
	if err != nil {
		return nil, err
	}
	defer fullFile.Close()

	fullFaceSrc, err := text.NewGoTextFaceSource(fullFile)
	if err != nil {
		return nil, err
	}

	outlineFile, err := fsys.Open("fonts/phantom-outline.ttf")
	if err != nil {
		return nil, err
	}
	defer outlineFile.Close()

	outlineFaceSrc, err := text.NewGoTextFaceSource(outlineFile)
	if err != nil {
		return nil, err
	}

	return &IntroText{
		text: []string{},
		fullFace: &text.GoTextFace{
			Source:    fullFaceSrc,
			Direction: text.DirectionLeftToRight,
			Size:      64,
		},
		outlineFace: &text.GoTextFace{
			Source:    outlineFaceSrc,
			Direction: text.DirectionLeftToRight,
			Size:      64,
		},
	}, nil
}

func (it *IntroText) Update(dt float64) error {
	return nil
}

func (it *IntroText) Draw(img *ebiten.Image) {
	it.textLock.RLock()
	for i, s := range it.text {
		{
			opts := &text.DrawOptions{}
			opts.GeoM.Translate(float64(img.Bounds().Dx())/2, (float64(i)*60)+200)
			opts.ColorScale.ScaleWithColor(color.White)
			opts.PrimaryAlign = text.AlignCenter
			text.Draw(img, s, it.fullFace, opts)
		}
		{
			opts := &text.DrawOptions{}
			opts.GeoM.Translate(float64(img.Bounds().Dx())/2, (float64(i)*60)+200)
			opts.ColorScale.ScaleWithColor(color.Black)
			opts.PrimaryAlign = text.AlignCenter
			text.Draw(img, s, it.outlineFace, opts)
		}
	}
	it.textLock.RUnlock()
}

func (it *IntroText) CreateText(text ...string) {
	it.textLock.Lock()
	it.text = append(it.text, text...)
	it.textLock.Unlock()
}

func (it *IntroText) AddText(text string) {
	it.textLock.Lock()
	it.text = append(it.text, text)
	it.textLock.Unlock()
}

func (it *IntroText) DeleteText() {
	it.textLock.Lock()
	it.text = []string{}
	it.textLock.Unlock()
}
