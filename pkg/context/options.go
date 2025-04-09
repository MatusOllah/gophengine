package context

import "io/fs"

type Options struct {
	AssetsFS           fs.FS
	OptionsConfigPath  string
	ProgressConfigPath string

	// If empty, uses default locale (from config)
	Locale string
}
