package music

import (
	"io/fs"

	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/vorbis"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

func getFreakyMenuMeta(fsys fs.FS) (bpm int, tween *gween.Tween, err error) {
	path := "music/freakyMenu.meta.hcl"
	b, err := fs.ReadFile(fsys, path)
	if err != nil {
		return
	}

	var v struct {
		BPM   int `hcl:"bpm"`
		Tween struct {
			Begin    float32 `hcl:"begin"`
			End      float32 `hcl:"end"`
			Duration float32 `hcl:"duration"`
		} `hcl:"tween,block"`
	}
	if err = hclsimple.Decode(path, b, nil, &v); err != nil {
		return
	}

	return v.BPM, gween.New(v.Tween.Begin, v.Tween.End, v.Tween.Duration, ease.Linear), nil
}

// FreakyMenuMusic represents the background music.
type FreakyMenuMusic struct {
	streamer beep.StreamSeekCloser
	format   beep.Format
	volume   *effects.Volume
	tween    *gween.Tween
	bpm      int
}

// New creates a new [FreakyMenuMusic].
func New(fsys fs.FS) (*FreakyMenuMusic, error) {
	m := &FreakyMenuMusic{}

	file, err := fsys.Open("music/freakyMenu.ogg")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m.streamer, m.format, err = vorbis.Decode(file)
	if err != nil {
		return nil, err
	}

	m.volume = &effects.Volume{
		Streamer: audio.MustLoop2(m.streamer),
		Base:     2,
		Silent:   false,
	}

	m.bpm, m.tween, err = getFreakyMenuMeta(fsys)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Stream streams the wrapped Volume.
func (m *FreakyMenuMusic) Stream(samples [][2]float64) (n int, ok bool) {
	return m.volume.Stream(samples)
}

// Err propagates the wrapped Volume's errors.
func (m *FreakyMenuMusic) Err() error {
	return m.volume.Err()
}

func (m *FreakyMenuMusic) Update() {
	vol, _ := m.tween.Update(1 / float32(ebiten.TPS()))
	speaker.Lock()
	m.volume.Volume = float64(vol)
	speaker.Unlock()
}

// BPM returns the BPM.
func (m *FreakyMenuMusic) BPM() int {
	return m.bpm
}

func (m *FreakyMenuMusic) Position() int {
	return m.streamer.Position()
}

func (m *FreakyMenuMusic) Format() beep.Format {
	return m.format
}
