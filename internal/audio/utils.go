package audio

import "github.com/gopxl/beep/v2"

func MustLoop2(s beep.StreamSeeker, opts ...beep.LoopOption) beep.Streamer {
	l, err := beep.Loop2(s, opts...)
	if err != nil {
		panic(err)
	}

	return l
}
