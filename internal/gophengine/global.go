package gophengine

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/MatusOllah/gophengine/internal/flagutil"
	"github.com/spf13/pflag"
	"github.com/vpxyz/xorshift/xorshift1024star"
)

type Global struct {
	Rand           *rand.Rand
	Version        string
	FNFVersion     string
	ScreenWidth    int
	ScreenHeight   int
	OptionsConfig  *config.Config
	ProgressConfig *config.Config
	Conductor      *Conductor
	FlagSet        *pflag.FlagSet
}

var G *Global

func InitGlobal() error {
	flagSet, err := initFlags()
	if err != nil {
		return err
	}

	optionsPath := flagutil.MustGetString(flagSet, "config")
	slog.Info(fmt.Sprintf("using config file %s", optionsPath))

	optionsConfig, err := config.New(optionsPath)
	if err != nil {
		return err
	}

	progressPath := flagutil.MustGetString(flagSet, "progress")
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
		OptionsConfig:  optionsConfig,
		ProgressConfig: progressConfig,
		Conductor:      NewConductor(100),
		FlagSet:        flagSet,
	}

	return nil
}
