package config

import (
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jeandeaual/go-locale"
)

func LoadDefaultOptions(cfg *Config) {
	slog.Info("[Config] loading defaults")

	userLocale, err := locale.GetLocale()
	if err != nil {
		panic(err)
	}
	slog.Info("got locale", "userLocale", userLocale)

	mustMarshalKey := func(key ebiten.Key) []byte {
		b, err := key.MarshalText()
		if err != nil {
			panic(err)
		}

		return b
	}

	cfg.SetData(map[string]interface{}{
		"Locale":     userLocale,
		"Fullscreen": false,

		"Controls.Up":         mustMarshalKey(ebiten.KeyUp),
		"Controls.Down":       mustMarshalKey(ebiten.KeyDown),
		"Controls.Left":       mustMarshalKey(ebiten.KeyLeft),
		"Controls.Right":      mustMarshalKey(ebiten.KeyRight),
		"Controls.Accept":     mustMarshalKey(ebiten.KeyEnter),
		"Controls.Back":       mustMarshalKey(ebiten.KeyEscape),
		"Controls.Pause":      mustMarshalKey(ebiten.KeyEscape),
		"Controls.Reset":      mustMarshalKey(ebiten.KeyR),
		"Controls.Fullscreen": mustMarshalKey(ebiten.KeyF11),
	})
}
