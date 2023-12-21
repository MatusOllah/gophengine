package gophengine

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"path/filepath"
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

func initFlags() (*pflag.FlagSet, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	flagSet := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	// help flag
	flagSet.BoolP("help", "h", false, "Shows this help message")

	flagSet.Bool("extract-assets", false, "Extract embedded assets")
	flagSet.String("config", filepath.Join(configDir, "GophEngine/config.gecfg"), "Path to config.gecfg config file")
	flagSet.String("progress", filepath.Join(configDir, "GophEngine/progress.gecfg"), "Path to progress.gecfg progress file")
	flagSet.Bool("vsync", false, "Enable VSync")

	if err := flagSet.Parse(os.Args[1:]); err != nil && err != pflag.ErrHelp {
		return nil, err
	}

	if flagutil.MustGetBool(flagSet, "help") {
		fmt.Printf("GophEngine is a Go implementation of Friday Night Funkin' with improvments.\n\n")
		fmt.Printf("Usage: %s [OPTIONS]\n\nOptions:\n", os.Args[0])
		fmt.Print(flagSet.FlagUsages())
		os.Exit(0)
	}
	return flagSet, nil
}
