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

var ErrNotFound error = errors.New("key not found")

type Config struct {
	dataLock sync.RWMutex
	data     map[string]interface{}

	file    *os.File
	encoder *gob.Encoder
	decoder *gob.Decoder
}

// Register simply calls gob.Register. If you are encoding a non-primitive type that implements something you should use this.
func Register(value interface{}) {
	gob.Register(value)
}

// New creates / opens and decodes a new Config.
func New(path string) (*Config, error) {
	var cfg *Config

	// checks if file exists
	// skontroluje ci subor existuje
	_, ferr := os.Stat(path)
	if ferr == nil {
		_cfg, err := Open(path)
		if err != nil {
			return nil, err
		}

		cfg = _cfg
	} else if errors.Is(ferr, os.ErrNotExist) {
		_cfg, err := Create(path)
		if err != nil {
			return nil, err
		}

		cfg = _cfg
	}

	return cfg, nil
}

// Create creates a new Config.
func Create(path string) (*Config, error) {
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

// Open opens a existing Config.
func Open(path string) (*Config, error) {
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

// SetData overwrites the map.
func (cfg *Config) SetData(m map[string]interface{}) {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()

	cfg.data = m
}

// Append appends m to the map.
func (cfg *Config) Append(m map[string]interface{}) {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()

	maps.Copy(cfg.data, m)
}

// Get gets a value from the map.
func (cfg *Config) Get(key string) (interface{}, error) {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()

	value, ok := cfg.data[key]
	if !ok {
		return nil, ErrNotFound
	}

	return value, nil
}

// MustGet simply calls Get and returns nil if an error occured.
func (cfg *Config) MustGet(key string) interface{} {
	value, err := cfg.Get(key)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return value
}

// Set sets key to value.
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

// Flush gob encodes and writes data to the file.
func (cfg *Config) Flush() error {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()

	if err := cfg.encoder.Encode(cfg.data); err != nil {
		return err
	}

	return nil
}

// Close simply calls (*os.File).Close and closes the file.
func (cfg *Config) Close() error {
	return cfg.file.Close()
}
