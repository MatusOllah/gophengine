package config

import "github.com/jeandeaual/go-locale"

func LoadDefaultOptions(cfg *Config) {
	userLocale, err := locale.GetLocale()
	if err != nil {
		panic(err)
	}

	cfg.SetData(map[string]interface{}{
		"Locale": userLocale,
	})
}
