package audio

import (
	"io/fs"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/vorbis"
)

func MustLoop2(s beep.StreamSeeker, opts ...beep.LoopOption) beep.Streamer {
	l, err := beep.Loop2(s, opts...)
	if err != nil {
		panic(err)
	}

	return l
}

func PlaySoundFromFS(fsys fs.FS, path string, vol float64, ch *MixerChannel) error {
	file, err := fsys.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	streamer, _, err := vorbis.Decode(file)
	if err != nil {
		return err
	}

	ch.Add(&effects.Volume{
		Streamer: streamer,
		Base:     2,
		Volume:   vol,
		Silent:   false,
	})

	return nil
}
