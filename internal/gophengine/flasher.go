package gophengine

import (
	"image/color"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type Flasher struct {
	alpha    float32
	flash    bool
	done     bool
	tween    *gween.Tween
	whiteImg *ebiten.Image
}

func NewFlasher(dur float32) (*Flasher, error) {
	white := ebiten.NewImage(G.Width, G.Height)
	white.Fill(color.White)

	return &Flasher{
		alpha:    1,
		flash:    false,
		done:     false,
		tween:    gween.New(1, 0, dur, ease.Linear),
		whiteImg: white,
	}, nil
}

func (f *Flasher) Draw(img *ebiten.Image) {
	if f.flash && !f.done {
		slog.Debug("drawing", "flash", f.flash, "done", f.done, "alpha", f.alpha)

		cm := colorm.ColorM{}
		cm.Scale(255, 255, 255, float64(f.alpha))
		colorm.DrawImage(img, f.whiteImg, cm, &colorm.DrawImageOptions{})
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
	f.tween.Reset()
	f.flash = true
}
