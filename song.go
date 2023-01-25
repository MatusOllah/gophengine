package main

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/ztrue/tracerr"
)

type Song struct {
	Song        string    `json:"song"`
	Notes       []Section `json:"notes"`
	Bpm         int       `json:"bpm"`
	NeedsVoices bool      `json:"needsVoices"`
	Speed       float64   `json:"speed"`
	Player1     string    `json:"player1"`
	Player2     string    `json:"player2"`
	ValidScore  bool      `json:"validScore"`
}

func LoadSongFromJSON(rawJson []byte) (*Song, error) {
	log.Info().Msg("loading song from JSON")

	var song struct {
		Song Song `json:"song"`
	}

	if err := json.Unmarshal(rawJson, &song); err != nil {
		return nil, tracerr.Wrap(err)
	}

	log.Info().Msgf("loaded song %s", song.Song.Song)

	song.Song.ValidScore = true

	return &song.Song, nil
}
