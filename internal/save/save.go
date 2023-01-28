package save

import (
	"encoding/gob"
	"errors"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/ztrue/tracerr"
)

type Save struct {
	data    map[string]interface{}
	file    *os.File
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func NewSave(path string) (*Save, error) {
	var save *Save

	// checks if file exists
	// skontroluje ci subor existuje
	_, ferr := os.Stat(path)
	if ferr == nil {
		_save, err := openSave(path)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		save = _save
	} else if errors.Is(ferr, os.ErrNotExist) {
		_save, err := createSave(path)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		save = _save
	}

	return save, nil
}

func createSave(path string) (*Save, error) {
	save := new(Save)

	// creates file
	// vytvori subor
	log.Info().Msgf("creating file %s", path)

	os.Mkdir(filepath.Dir(path), os.ModePerm)

	file, _ := os.Create(path)

	save.data = make(map[string]interface{})
	save.file = file
	save.encoder = gob.NewEncoder(save.file)
	save.decoder = gob.NewDecoder(save.file)

	return save, nil
}

func openSave(path string) (*Save, error) {
	save := new(Save)

	// opens file
	// otvori subor
	log.Info().Msgf("opening file %s", path)
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	save.file = file
	save.encoder = gob.NewEncoder(save.file)
	save.decoder = gob.NewDecoder(save.file)

	// decodes raw data
	// dekoduje surove data
	log.Info().Msg("deserializing")
	if err := save.decoder.Decode(&save.data); err != nil {
		return nil, tracerr.Wrap(err)
	}

	return save, nil
}

func (s *Save) Read(key string) interface{} {
	log.Info().Msgf("reading from key %s", key)
	return s.data[key]
}

func (s *Save) Write(key string, value interface{}) {
	log.Info().Msgf("writing value %v to key %s", value, key)
	s.data[key] = value
}

func (s *Save) Delete(key string) {
	log.Info().Msgf("deleting key %s", key)
	delete(s.data, key)
}

func (s *Save) Flush() error {
	log.Info().Msg("serializing")
	if err := s.encoder.Encode(s.data); err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (s *Save) Close() error {
	if err := s.file.Close(); err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
