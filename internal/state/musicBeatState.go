package state

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	ge "github.com/MatusOllah/gophengine/internal/gophengine"
)

type MusicBeatState struct {
	LastBeat float64
	LastStep float64

	CurBeat int
	CurStep int
}

func (s *MusicBeatState) Update(dt float64) error {
	oldStep := s.CurStep

	s.UpdateCurStep()
	s.UpdateBeat()

	if oldStep != s.CurStep && s.CurStep > 0 {
		s.StepHit()
	}

	return nil
}

func (s *MusicBeatState) Draw(screen *ebiten.Image) {
	return
}

func (s *MusicBeatState) UpdateBeat() {
	s.CurBeat = int(math.Floor(float64(s.CurStep / 4)))
}

func (s *MusicBeatState) UpdateCurStep() {
	lastChange := &ge.BPMChangeEvent{
		StepTime: 0,
		SongTime: 0,
		Bpm:      0,
	}

	for i := range ge.G.Conductor.BPMChangeMap {
		if ge.G.Conductor.SongPosition >= ge.G.Conductor.BPMChangeMap[i].SongTime {
			lastChange = &ge.G.Conductor.BPMChangeMap[i]
		}
	}

	s.CurStep = lastChange.StepTime + int(math.Floor((ge.G.Conductor.SongPosition-lastChange.SongTime)/ge.G.Conductor.StepCrochet))
}

func (s *MusicBeatState) StepHit() {
	if s.CurStep%4 == 0 {
		s.BeatHit()
	}
}

func (s *MusicBeatState) BeatHit() {
	return
}
