package gophengine

import (
	"io/fs"

	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/vorbis"
)

func PlaySoundFromFS(fsys fs.FS, path string, vol float64) error {
	file, err := fsys.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	streamer, _, err := vorbis.Decode(file)
	if err != nil {
		return err
	}

	G.Mixer.Add(&effects.Volume{
		Streamer: streamer,
		Base:     2,
		Volume:   vol,
		Silent:   false,
	})

	return nil
}
