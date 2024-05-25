package config

import (
	"log/slog"

	"github.com/jeandeaual/go-locale"
)

func LoadDefaultOptions(cfg *Config) {
	slog.Info("config: loading defaults")

	userLocale, err := locale.GetLocale()
	if err != nil {
		panic(err)
	}
	slog.Info("got locale", "userLocale", userLocale)

	cfg.SetData(map[string]interface{}{
		"Locale":     userLocale,
		"Fullscreen": false,
	})
}
