package context

import "io/fs"

type NewContextConfig struct {
	AssetsFS           fs.FS
	OptionsConfigPath  string
	ProgressConfigPath string
}
