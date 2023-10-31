package gophengine

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/vpxyz/xorshift/xorshift1024star"
)

type Global struct {
	Rand           *rand.Rand
	Version        string
	FNFVersion     string
	ScreenWidth    int
	ScreenHeight   int
	AudioContext   *audio.Context
	OptionsConfig  *config.Config
	ProgressConfig *config.Config
	Conductor      *Conductor
}

var G *Global

func InitGlobal() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	optionsPath := filepath.Join(configDir, "GophEngine/config.gecfg")
	if Options.Config != "" {
		optionsPath = Options.Config
	}
	slog.Info(fmt.Sprintf("using config file %s", optionsPath))

	optionsConfig, err := config.New(optionsPath)
	if err != nil {
		return err
	}

	progressPath := filepath.Join(configDir, "GophEngine/progress.gecfg")
	if Options.Progress != "" {
		progressPath = Options.Progress
	}
	slog.Info(fmt.Sprintf("using progress file %s", progressPath))

	progressConfig, err := config.New(progressPath)
	if err != nil {
		return err
	}

	rand := rand.New(xorshift1024star.NewSource(time.Now().UTC().UnixNano()))

	G = &Global{
		Rand:           rand,
		Version:        "1.0",
		FNFVersion:     "0.2.7.1",
		ScreenWidth:    1280,
		ScreenHeight:   720,
		AudioContext:   audio.NewContext(48000),
		OptionsConfig:  optionsConfig,
		ProgressConfig: progressConfig,
		Conductor:      NewConductor(100),
	}

	return nil
}
