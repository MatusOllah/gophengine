package config

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"maps"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var ErrKeyNotExist error = errors.New("key does not exist")

// MapFunc represents a mapping function. It takes in a key and value and returns a new value.
type MapFunc func(string, any) any

// Self-explanatory.
type Config struct {
	data     map[string]any
	dataLock sync.RWMutex

	file    *os.File
	buf     *bytes.Buffer
	encoder *gob.Encoder
	decoder *gob.Decoder
}

// Register simply calls gob.Register.
// If you are encoding a non-primitive type (like a struct or map) that implements something you should use this.
// At least that's how I understand it.
func Register(value any) {
	gob.Register(value)
}

// New creates / opens and decodes a new Config.
func New(path string, loadDefaults bool) (*Config, error) {
	if fileExists(path) {
		cfg, err := Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open config: %w", err)
		}

		return cfg, nil
	} else {
		slog.Warn("config not found, creating new config", "path", path)
		cfg, err := Create(path, loadDefaults)
		if err != nil {
			return nil, fmt.Errorf("failed to create config: %w", err)
		}

		return cfg, nil
	}
}

// Create creates a new Config.
func Create(path string, loadDefaults bool) (*Config, error) {
	cfg := new(Config)

	if runtime.GOARCH == "wasm" {
		cfg.file = nil
	} else {
		os.Mkdir(filepath.Dir(path), os.ModePerm)

		file, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		cfg.file = file
	}

	cfg.data = make(map[string]any)
	cfg.buf = new(bytes.Buffer)
	cfg.encoder = gob.NewEncoder(cfg.buf)
	cfg.decoder = gob.NewDecoder(cfg.buf)

	if loadDefaults {
		LoadDefaultOptions(cfg)
	}

	return cfg, nil
}

// Open opens a existing Config.
func Open(path string) (*Config, error) {
	cfg := new(Config)

	if runtime.GOARCH == "wasm" {
		cfg.file = nil
	} else {
		file, err := os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}

		cfg.file = file
	}

	cfg.buf = new(bytes.Buffer)
	cfg.encoder = gob.NewEncoder(cfg.buf)
	cfg.decoder = gob.NewDecoder(cfg.buf)

	_, err := io.Copy(cfg.buf, cfg.file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file to buffer: %w", err)
	}

	cfg.dataLock.Lock()
	if err := cfg.decoder.Decode(&cfg.data); err != nil {
		return nil, fmt.Errorf("failed to decode gob: %w", err)
	}
	cfg.dataLock.Unlock()

	return cfg, nil
}

// Data returns a copy of the map.
func (cfg *Config) Data() map[string]any {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()
	return maps.Clone(cfg.data)
}

// SetData overwrites the map.
func (cfg *Config) SetData(m map[string]any) {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()
	clear(cfg.data)
	maps.Copy(cfg.data, m)
}

// Append appends m to the map.
func (cfg *Config) Append(m map[string]any) {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()
	maps.Copy(cfg.data, m)
}

// Get gets a value from the map. If there is an error, it will be of type *KeyError.
func (cfg *Config) Get(key string) (any, error) {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()
	value, ok := cfg.data[key]
	if !ok {
		return nil, &KeyError{Op: "get", Key: key, Err: ErrKeyNotExist}
	}
	return value, nil
}

// MustGet simply calls Get and returns nil if an error occurred.
func (cfg *Config) MustGet(key string) any {
	value, err := cfg.Get(key)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return value
}

// GetWithFallback gets a value from the map and returns the given fallback if not found.
func (cfg *Config) GetWithFallback(key string, fallback any) any {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()
	value, ok := cfg.data[key]
	if !ok {
		slog.Warn("key not found", "key", key, "fallback", fallback)
		return fallback
	}
	return value
}

// Set sets key to value.
func (cfg *Config) Set(key string, value any) {
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

// Toggle toggles a bool value.
func (cfg *Config) Toggle(key string) {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()
	cfg.data[key] = !cfg.data[key].(bool)
}

// Exists checks if key exists.
func (cfg *Config) Exists(key string) bool {
	cfg.dataLock.RLock()
	defer cfg.dataLock.RUnlock()
	_, ok := cfg.data[key]
	return ok
}

// Map iterates over the map and applies the [MapFunc] to every item.
func (cfg *Config) Map(fn MapFunc) {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()
	for k, v := range cfg.data {
		cfg.data[k] = fn(k, v)
	}
}

// Wipe wipes (clears) the map.
func (cfg *Config) Wipe() {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()
	clear(cfg.data)
}

// GobBytes returns the actual gob-encoded data.
func (cfg *Config) GobBytes() []byte {
	return cfg.buf.Bytes()
}

// Flush gob encodes and writes data to the file.
func (cfg *Config) Flush() error {
	cfg.dataLock.Lock()
	defer cfg.dataLock.Unlock()

	if err := cfg.encoder.Encode(cfg.data); err != nil {
		return fmt.Errorf("failed to encode gob: %w", err)
	}

	if cfg.file == nil {
		return nil
	}

	if err := cfg.file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	if _, err := cfg.file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to seek file: %w", err)
	}

	_, err := io.Copy(cfg.file, cfg.buf)
	if err != nil {
		return fmt.Errorf("failed to write to file from buffer: %w", err)
	}

	if err := cfg.file.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}

	cfg.buf.Reset()

	return nil
}

// Close flushes and closes the file.
func (cfg *Config) Close() error {
	if err := cfg.Flush(); err != nil {
		return err
	}

	if cfg.file == nil {
		return nil
	}

	return cfg.file.Close()
}
