package funkin

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/funkin/chart"
)

type Conductor struct {
	Bpm            int
	Crochet        float64
	StepCrochet    float64
	SongPosition   float64
	LastSongPos    float64
	Offset         float64
	SafeFrames     int
	SafeZoneOffset float64
	BPMChangeMap   []BPMChangeEvent
}

func NewConductor(bpm int) *Conductor {
	c := new(Conductor)

	c.Bpm = bpm
	c.Crochet = (60 / float64(bpm)) * 1000
	c.StepCrochet = c.Crochet / 4
	c.Offset = 0
	c.SafeFrames = 10
	c.SafeZoneOffset = (float64(c.SafeFrames) / 60) * 1000

	return c
}

func (c *Conductor) MapBPMChanges(song *chart.Song) {
	c.BPMChangeMap = []BPMChangeEvent{}

	var curBPM int = song.Bpm
	var totalSteps int = 0
	var totalPos float64 = 0
	for _, note := range song.Notes {
		if note.ChangeBPM && note.Bpm != curBPM {
			curBPM = note.Bpm
			c.BPMChangeMap = append(c.BPMChangeMap, BPMChangeEvent{
				StepTime: totalSteps,
				SongTime: totalPos,
				Bpm:      curBPM,
			})
		}

		var deltaSteps int = note.LengthInSteps
		totalSteps += deltaSteps
		totalPos += ((60 / float64(curBPM)) * 1000 / 4) * float64(deltaSteps)
	}

	slog.Info("new BPM map", "BPMChangeMap", c.BPMChangeMap)
}

func (c *Conductor) ChangeBPM(newBpm int) {
	c.Bpm = newBpm

	c.Crochet = (60 / float64(c.Bpm)) * 1000
	c.StepCrochet = c.Crochet / 4
}
