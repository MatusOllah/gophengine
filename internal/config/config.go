package config

import (
	"encoding/gob"
	"errors"
	"log/slog"
	"maps"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	dataLock sync.RWMutex
	data     map[string]interface{}

	file    *os.File
	encoder *gob.Encoder
	decoder *gob.Decoder
}

// Register simply calls gob.Register.
func Register(value interface{}) {
	gob.Register(value)
}

func New(path string) (*Config, error) {
	var cfg *Config

	// checks if file exists
	// skontroluje ci subor existuje
	_, ferr := os.Stat(path)
	if ferr == nil {
		_cfg, err := openConfig(path)
		if err != nil {
			return nil, err
		}

		cfg = _cfg
	} else if errors.Is(ferr, os.ErrNotExist) {
		_cfg, err := createConfig(path)
		if err != nil {
			return nil, err
		}

		cfg = _cfg
	}

	return cfg, nil
}

func createConfig(path string) (*Config, error) {
	cfg := new(Config)

	// creates file
	// vytvori subor
	os.Mkdir(filepath.Dir(path), os.ModePerm)

	file, _ := os.Create(path)

	cfg.data = make(map[string]interface{})
	cfg.file = file
	cfg.encoder = gob.NewEncoder(cfg.file)
	cfg.decoder = gob.NewDecoder(cfg.file)

	return cfg, nil
}

func openConfig(path string) (*Config, error) {
	cfg := new(Config)

	// opens file
	// otvori subor
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	cfg.file = file
	cfg.encoder = gob.NewEncoder(cfg.file)
	cfg.decoder = gob.NewDecoder(cfg.file)

	// decodes raw data
	// dekoduje surove data
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()

	if err := cfg.decoder.Decode(&cfg.data); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Data returns a copy of the map.
func (cfg *Config) Data() map[string]interface{} {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()

	data := cfg.data

	return data
}

// Append appends m to the map.
func (cfg *Config) Append(m map[string]interface{}) {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()

	maps.Copy(cfg.data, m)
}

func (cfg *Config) Get(key string) (interface{}, error) {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()

	value, ok := cfg.data[key]
	if !ok {
		return nil, errors.New("key not found")
	}

	return value, nil
}

func (cfg *Config) MustGet(key string) interface{} {
	value, err := cfg.Get(key)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return value
}

func (cfg *Config) Set(key string, value interface{}) {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()

	cfg.data[key] = value
}

// Delete deletes key from the map.
func (cfg *Config) Delete(key string) {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()

	delete(cfg.data, key)
}

// Wipe wipes (clears) the map.
func (cfg *Config) Wipe() {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()

	clear(cfg.data)
}

func (cfg *Config) Flush() error {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()

	if err := cfg.encoder.Encode(cfg.data); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) Close() error {
	if err := cfg.file.Close(); err != nil {
		return err
	}

	return nil
}
