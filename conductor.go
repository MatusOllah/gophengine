package main

import "github.com/rs/zerolog/log"

type Conductor struct {
	Bpm            int
	Crochet        float64
	StepCrochet    float64
	SongPosition   float64
	lastSongPos    float64
	Offset         float64
	SafeFrames     int
	SafeZoneOffset float64
	BpmChangeMap   []BPMChangeEvent
}

var conductor = &Conductor{
	Bpm: 100,
	//Crochet: float64((60 / bpm) * 1000),
	//StepCrochet: crochet / 4,
	Offset:     0,
	SafeFrames: 10,
	//SafeZoneOffset: float64((safeFrames / 60) * 1000),
}

func (c *Conductor) MapBPMChanges(song *Song) {
	c.BpmChangeMap = []BPMChangeEvent{}

	var curBPM int = song.Bpm
	var totalSteps int = 0
	var totalPos float64 = 0
	for i := range song.Notes {
		if song.Notes[i].ChangeBPM && song.Notes[i].Bpm != curBPM {
			curBPM = song.Notes[i].Bpm
			c.BpmChangeMap = append(c.BpmChangeMap, BPMChangeEvent{
				StepTime: totalSteps,
				SongTime: totalPos,
				Bpm:      curBPM,
			})
		}

		var deltaSteps int = song.Notes[i].LengthInSteps
		totalSteps += deltaSteps
		totalPos += float64(((60 / curBPM) * 1000 / 4) * deltaSteps)
	}

	log.Info().Msgf("new BPM map %v", c.BpmChangeMap)
}

func (c *Conductor) ChangeBPM(newBpm int) {
	c.Bpm = newBpm

	c.Crochet = float64((60 / c.Bpm) * 1000)
	c.StepCrochet = c.Crochet / 4
}
