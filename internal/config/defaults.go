package config

func LoadDefaultOptions(cfg *Config) {
	cfg.SetData(map[string]interface{}{
		"Locale": "en",
	})
}
