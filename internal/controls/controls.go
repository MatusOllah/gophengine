package controls

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/config"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionUp input.Action = iota
	ActionDown
	ActionLeft
	ActionRight
	ActionAccept
	ActionBack
	ActionPause
	ActionReset
	ActionFullscreen
)

func LoadKeymapFromConfig(cfg *config.Config) (input.Keymap, error) {
	slog.Info("loading keymap")

	keymap := input.Keymap{}

	getKeys := func(id string, action input.Action) error {
		_keys, err := cfg.Get(id)
		if err != nil {
			return err
		}

		var keys []input.Key
		for _, s := range _keys.([]string) {
			key, err := input.ParseKey(s)
			if err != nil {
				return err
			}
			keys = append(keys, key)
		}

		keymap[action] = keys

		return nil
	}

	if err := getKeys("Controls.Up", ActionUp); err != nil {
		return nil, err
	}
	if err := getKeys("Controls.Down", ActionDown); err != nil {
		return nil, err
	}
	if err := getKeys("Controls.Left", ActionLeft); err != nil {
		return nil, err
	}
	if err := getKeys("Controls.Right", ActionRight); err != nil {
		return nil, err
	}
	if err := getKeys("Controls.Accept", ActionAccept); err != nil {
		return nil, err
	}
	if err := getKeys("Controls.Back", ActionBack); err != nil {
		return nil, err
	}
	if err := getKeys("Controls.Pause", ActionPause); err != nil {
		return nil, err
	}
	if err := getKeys("Controls.Reset", ActionReset); err != nil {
		return nil, err
	}
	if err := getKeys("Controls.Fullscreen", ActionFullscreen); err != nil {
		return nil, err
	}

	slog.Info("loading keymap OK", "keymap", keymap)

	return keymap, nil
}
