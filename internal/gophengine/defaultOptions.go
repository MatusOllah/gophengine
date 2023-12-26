package gophengine

import "github.com/MatusOllah/gophengine/internal/config"

func loadDefaultOptions(cfg *config.Config) {
	cfg.SetData(map[string]interface{}{
		"locale": "en",
	})
}
