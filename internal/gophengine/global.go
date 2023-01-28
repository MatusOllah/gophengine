package gophengine

import (
	"os"
	"path/filepath"

	"github.com/MatusOllah/gophengine/internal/save"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/ztrue/tracerr"
)

type Global struct {
	Version      string
	FNFVersion   string
	ScreenWidth  int
	ScreenHeight int
	AssetsDir    string
	AudioContext *audio.Context
	ConfigSave   *save.Save
	ProgressSave *save.Save
	Conductor    *Conductor
}

var G *Global

func InitGlobal() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return tracerr.Wrap(err)
	}

	configSave, err := save.NewSave(filepath.Join(configDir, "GophEngine/config.gecfg"))
	if err != nil {
		return tracerr.Wrap(err)
	}

	progressSave, err := save.NewSave(filepath.Join(configDir, "GophEngine/progress.gecfg"))
	if err != nil {
		return tracerr.Wrap(err)
	}

	G = &Global{
		Version:      "1.0",
		FNFVersion:   "0.2.7.1",
		ScreenWidth:  1280,
		ScreenHeight: 720,
		AssetsDir:    "assets",
		AudioContext: audio.NewContext(48000),
		ConfigSave:   configSave,
		ProgressSave: progressSave,
		Conductor:    NewConductor(100),
	}

	return nil
}
