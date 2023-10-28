package gophengine

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
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

	var tmpSong struct {
		Song Song `json:"song"`
	}

	if err := json.Unmarshal(rawJson, &tmpSong); err != nil {
		return nil, err
	}

	log.Info().Msgf("loaded song %s", tmpSong.Song.Song)

	tmpSong.Song.ValidScore = true

	return &tmpSong.Song, nil
}
