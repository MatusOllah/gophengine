package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
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

}

func (s *MusicBeatState) UpdateBeat() {
	s.CurBeat = int(math.Floor(float64(s.CurStep / 4)))
}

func (s *MusicBeatState) UpdateCurStep() {
	lastChange := &BPMChangeEvent{
		StepTime: 0,
		SongTime: 0,
		Bpm:      0,
	}

	for i := range conductor.BpmChangeMap {
		if conductor.SongPosition >= conductor.BpmChangeMap[i].SongTime {
			lastChange = &conductor.BpmChangeMap[i]
		}
	}

	s.CurStep = lastChange.StepTime + int(math.Floor((conductor.SongPosition-lastChange.SongTime)/conductor.StepCrochet))
}

func (s *MusicBeatState) StepHit() {
	if s.CurStep%4 == 0 {
		s.BeatHit()
	}
}

func (s *MusicBeatState) BeatHit() {
	return
}
