package gophengine

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

	var tmpSong struct {
		song Song `json:"song"`
	}

	if err := json.Unmarshal(rawJson, &tmpSong); err != nil {
		return nil, tracerr.Wrap(err)
	}

	log.Info().Msgf("loaded song %s", tmpSong.song.Song)

	tmpSong.song.ValidScore = true

	return &tmpSong.song, nil
}
