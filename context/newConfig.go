package context

import "io/fs"

type NewContextConfig struct {
	AssetsFS           fs.FS
	OptionsConfigPath  string
	ProgressConfigPath string
	Version            string

	// If empty, uses default locale (from config)
	Locale string
}
