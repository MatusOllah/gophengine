package audioutil

import (
	"io/fs"

	"github.com/MatusOllah/gophengine/context"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/vorbis"
)

func PlaySoundFromFS(ctx *context.Context, fsys fs.FS, path string, vol float64) error {
	file, err := fsys.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	streamer, format, err := vorbis.Decode(file)
	if err != nil {
		return err
	}

	ctx.AudioMixer.SFX.Add(&effects.Volume{
		Streamer: beep.Resample(ctx.AudioResampleQuality, format.SampleRate, ctx.SampleRate, streamer),
		Base:     2,
		Volume:   vol,
		Silent:   false,
	})

	return nil
}
