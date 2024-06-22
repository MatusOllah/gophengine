package gophengine

import (
	"math"
)

// MusicBeat is basically MusicBeatState in vanilla FNF.
type MusicBeat struct {
	LastBeat float64
	LastStep float64

	CurBeat int
	CurStep int

	StepHitFunc func(int)
	BeatHitFunc func(int)

	c *Conductor
}

func NewMusicBeat(c *Conductor) *MusicBeat {
	return &MusicBeat{
		StepHitFunc: func(_ int) {},
		BeatHitFunc: func(_ int) {},
		c:           c,
	}
}

func (mb *MusicBeat) Update() {
	oldStep := mb.CurStep

	mb.UpdateCurStep()
	mb.UpdateBeat()

	if oldStep != mb.CurStep && mb.CurStep > 0 {
		mb.StepHit()
	}
}

func (mb *MusicBeat) UpdateBeat() {
	mb.CurBeat = int(math.Floor(float64(mb.CurStep) / 4))
}

func (mb *MusicBeat) UpdateCurStep() {
	lastChange := BPMChangeEvent{
		StepTime: 0,
		SongTime: 0,
		Bpm:      0,
	}

	for _, bcm := range mb.c.BPMChangeMap {
		if mb.c.SongPosition >= bcm.SongTime {
			lastChange = bcm
		}
	}

	mb.CurStep = lastChange.StepTime + int(math.Floor((mb.c.SongPosition-lastChange.SongTime)/mb.c.StepCrochet))
}

func (mb *MusicBeat) StepHit() {
	if mb.CurStep%4 == 0 {
		mb.BeatHit()
	}

	mb.StepHitFunc(mb.CurStep)
}

func (mb *MusicBeat) BeatHit() {
	mb.BeatHitFunc(mb.CurBeat)
}
