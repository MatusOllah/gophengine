package effects

import (
	"image/color"

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

func NewFlasher(width, height int, dur float32) *Flasher {
	white := ebiten.NewImage(width, height)
	white.Fill(color.White)

	return &Flasher{
		alpha:    1,
		flash:    false,
		done:     false,
		tween:    gween.New(1, 0, dur, ease.Linear),
		whiteImg: white,
	}
}

func (f *Flasher) Draw(img *ebiten.Image) {
	if f.flash && !f.done {
		cm := colorm.ColorM{}
		cm.Scale(255, 255, 255, float64(f.alpha))
		colorm.DrawImage(img, f.whiteImg, cm, &colorm.DrawImageOptions{})
	}
}

func (f *Flasher) Update() {
	if f.done {
		f.flash = false
		f.done = false
		f.tween.Reset()
		return
	}

	if f.flash && !f.done {
		f.alpha, f.done = f.tween.Update(1 / float32(ebiten.TPS()))
	}
}

func (f *Flasher) Flash() {
	f.tween.Reset()
	f.flash = true
}
