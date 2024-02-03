package gophengine

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type Flasher struct {
	shader *ebiten.Shader
	alpha  float32
	flash  bool
	done   bool
	tween  *gween.Tween
}

func NewFlasher(dur float32) (*Flasher, error) {
	slog.Info("compiling shaders")
	b, err := assets.FS.ReadFile("data/shaders/flash.kage")
	if err != nil {
		return nil, err
	}

	shader, err := ebiten.NewShader(b)
	if err != nil {
		return nil, err
	}

	return &Flasher{
		shader: shader,
		alpha:  1,
		flash:  false,
		done:   false,
		tween:  gween.New(1, 0, dur, ease.Linear),
	}, nil
}

func (f *Flasher) Draw(img *ebiten.Image) {
	if f.flash && !f.done {
		//slog.Info("drawing", "flash", f.flash, "done", f.done, "alpha", f.alpha)

		newImg := ebiten.NewImage(img.Bounds().Dx(), img.Bounds().Dy())

		op := &ebiten.DrawRectShaderOptions{}
		op.Images[0] = img
		op.Uniforms = map[string]any{
			"Opacity": f.alpha,
		}
		newImg.DrawRectShader(newImg.Bounds().Dx(), newImg.Bounds().Dy(), f.shader, op)

		img.DrawImage(newImg, &ebiten.DrawImageOptions{})
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
