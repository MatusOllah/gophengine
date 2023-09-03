package gophengine

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/MatusOllah/gophengine/internal/save"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/vpxyz/xorshift/xorshift1024star"
	"github.com/ztrue/tracerr"
)

type Global struct {
	Rand         *rand.Rand
	Version      string
	FNFVersion   string
	ScreenWidth  int
	ScreenHeight int
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

	rand := rand.New(xorshift1024star.NewSource(time.Now().UTC().UnixNano()))

	G = &Global{
		Rand:         rand,
		Version:      "1.0",
		FNFVersion:   "0.2.7.1",
		ScreenWidth:  1280,
		ScreenHeight: 720,
		AudioContext: audio.NewContext(48000),
		ConfigSave:   configSave,
		ProgressSave: progressSave,
		Conductor:    NewConductor(100),
	}

	return nil
}