package config

import (
	"log/slog"

	"github.com/jeandeaual/go-locale"
	input "github.com/quasilyte/ebitengine-input"
)

func LoadDefaultOptions(cfg *Config) {
	slog.Info("[Config] loading defaults")

	userLocale, err := locale.GetLocale()
	if err != nil {
		panic(err)
	}
	slog.Info("got locale", "userLocale", userLocale)

	cfg.SetData(map[string]interface{}{
		"Locale":     userLocale,
		"Fullscreen": false,

		"Controls.Up":         []string{input.KeyUp.String(), input.KeyW.String()},
		"Controls.Down":       []string{input.KeyDown.String(), input.KeyS.String()},
		"Controls.Left":       []string{input.KeyLeft.String(), input.KeyA.String()},
		"Controls.Right":      []string{input.KeyRight.String(), input.KeyD.String()},
		"Controls.Accept":     []string{input.KeyEnter.String()},
		"Controls.Back":       []string{input.KeyEscape.String(), input.KeyBackspace.String()},
		"Controls.Pause":      []string{input.KeyEscape.String()},
		"Controls.Reset":      []string{input.KeyR.String()},
		"Controls.Fullscreen": []string{input.KeyF11.String()},
	})
}
