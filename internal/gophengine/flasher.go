package gophengine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type Flasher struct {
	color color.NRGBA
	alpha float32
	flash bool
	done  bool
	tween *gween.Tween
}

func NewFlasher(c color.NRGBA, dur float32) *Flasher {
	a := float32(c.A)

	return &Flasher{
		color: c,
		alpha: a,
		flash: false,
		done:  false,
		tween: gween.New(a, 0, dur, ease.Linear),
	}
}

func (f *Flasher) Draw(img *ebiten.Image) {
	if f.flash && !f.done {
		img.Fill(color.NRGBA{
			R: f.color.R,
			G: f.color.G,
			B: f.color.B,
			A: uint8(f.alpha),
		})
	}
}

func (f *Flasher) Update(dt float64) error {
	if f.done {
		f.flash = false
		f.done = false
		f.tween.Reset()
		return nil
	}

	if f.flash && !f.done {
		f.alpha, f.done = f.tween.Update(float32(dt))
	}

	return nil
}

func (f *Flasher) Flash() {
	f.flash = true
}
