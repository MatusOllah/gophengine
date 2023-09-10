package save

import (
	"encoding/gob"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/ztrue/tracerr"
)

type Save struct {
	dataLock sync.RWMutex
	data     map[string]interface{}

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
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	save.file = file
	save.encoder = gob.NewEncoder(save.file)
	save.decoder = gob.NewDecoder(save.file)

	// decodes raw data
	// dekoduje surove data
	save.dataLock.RLock()
	defer save.dataLock.RUnlock()

	if err := save.decoder.Decode(&save.data); err != nil {
		return nil, tracerr.Wrap(err)
	}

	return save, nil
}

func (s *Save) Get(key string) (interface{}, error) {
	s.dataLock.RLock()
	defer s.dataLock.RUnlock()

	value, ok := s.data[key]
	if !ok {
		return nil, tracerr.New("key not found")
	}

	return value, nil
}

func (s *Save) MustGet(key string) interface{} {
	value, err := s.Get(key)
	if err != nil {
		panic(err)
	}

	return value
}

func (s *Save) Set(key string, value interface{}) {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()

	s.data[key] = value
}

func (s *Save) Delete(key string) {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()

	delete(s.data, key)
}

func (s *Save) Flush() error {
	s.dataLock.RLock()
	defer s.dataLock.RUnlock()

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
