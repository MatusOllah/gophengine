package gophengine

import "github.com/rs/zerolog/log"

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
	c.Crochet = float64((60 / bpm) * 1000)
	c.StepCrochet = c.Crochet / 4
	c.Offset = 0
	c.SafeFrames = 10
	c.SafeZoneOffset = float64((c.SafeFrames / 60) * 1000)

	return c
}

func (c *Conductor) MapBPMChanges(song *Song) {
	c.BPMChangeMap = []BPMChangeEvent{}

	var curBPM int = song.Bpm
	var totalSteps int = 0
	var totalPos float64 = 0
	for i := range song.Notes {
		if song.Notes[i].ChangeBPM && song.Notes[i].Bpm != curBPM {
			curBPM = song.Notes[i].Bpm
			c.BPMChangeMap = append(c.BPMChangeMap, BPMChangeEvent{
				StepTime: totalSteps,
				SongTime: totalPos,
				Bpm:      curBPM,
			})
		}

		var deltaSteps int = song.Notes[i].LengthInSteps
		totalSteps += deltaSteps
		totalPos += float64(((60 / curBPM) * 1000 / 4) * deltaSteps)
	}

	log.Info().Msgf("new BPM map %v", c.BPMChangeMap)
}

func (c *Conductor) ChangeBPM(newBpm int) {
	c.Bpm = newBpm

	c.Crochet = float64((60 / c.Bpm) * 1000)
	c.StepCrochet = c.Crochet / 4
}
