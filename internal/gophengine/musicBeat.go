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
}

func NewMusicBeat() *MusicBeat {
	return &MusicBeat{
		StepHitFunc: func(_ int) {},
		BeatHitFunc: func(_ int) {},
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

	for _, bcm := range G.Conductor.BPMChangeMap {
		if G.Conductor.SongPosition >= bcm.SongTime {
			lastChange = bcm
		}
	}

	mb.CurStep = lastChange.StepTime + int(math.Floor((G.Conductor.SongPosition-lastChange.SongTime)/G.Conductor.StepCrochet))
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
