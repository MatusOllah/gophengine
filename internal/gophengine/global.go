package gophengine

import (
	"fmt"
	"io/fs"
	"log/slog"
	"math/rand"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/MatusOllah/gophengine/internal/flagutil"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/pflag"
	"github.com/vpxyz/xorshift/xorshift1024star"
	"golang.org/x/text/language"
)

type Global struct {
	Rand           *rand.Rand
	Version        string
	FNFVersion     string
	ScreenWidth    int
	ScreenHeight   int
	Width          int
	Height         int
	OptionsConfig  *config.Config
	ProgressConfig *config.Config
	Conductor      *Conductor
	FlagSet        *pflag.FlagSet
	Localizer      *i18n.Localizer
}

var G *Global

func InitGlobal() error {
	flagSet, err := initFlags()
	if err != nil {
		return err
	}

	optionsPath := flagutil.MustGetString(flagSet, "config")
	slog.Info(fmt.Sprintf("using config file %s", optionsPath))

	optionsConfig, err := config.New(optionsPath, true)
	if err != nil {
		return err
	}

	if err := overrideConfigValues(optionsConfig, flagSet); err != nil {
		return err
	}

	if flagutil.MustGetBool(flagSet, "config-load-defaults") {
		slog.Info("loading defaults")
		config.LoadDefaultOptions(optionsConfig)
	}

	progressPath := flagutil.MustGetString(flagSet, "progress")
	slog.Info(fmt.Sprintf("using progress file %s", progressPath))

	progressConfig, err := config.New(progressPath, false)
	if err != nil {
		return err
	}

	rand := rand.New(xorshift1024star.NewSource(time.Now().UTC().UnixNano()))

	// initialize localizer
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if err := loadLocales(bundle); err != nil {
		return err
	}

	locale := optionsConfig.MustGet("locale").(string)
	slog.Info("using locale", "locale", locale)
	localizer := i18n.NewLocalizer(bundle, locale)

	G = &Global{
		Rand:           rand,
		Version:        "1.0",
		FNFVersion:     "0.2.7.1",
		ScreenWidth:    1280,
		ScreenHeight:   720,
		Width:          1280,
		Height:         702,
		OptionsConfig:  optionsConfig,
		ProgressConfig: progressConfig,
		Conductor:      NewConductor(100),
		FlagSet:        flagSet,
		Localizer:      localizer,
	}

	return nil
}

func loadLocales(bundle *i18n.Bundle) error {
	files, err := fs.Glob(assets.FS, "data/locale/*.toml")
	if err != nil {
		return err
	}

	for _, file := range files {
		slog.Info("loading locale", "file", file)
		if _, err := bundle.LoadMessageFileFS(assets.FS, file); err != nil {
			return err
		}
	}

	return nil
}

func overrideConfigValues(cfg *config.Config, flagSet *pflag.FlagSet) error {
	// string
	ss, err := flagSet.GetStringSlice("config-string")
	if err != nil {
		return err
	}
	for _, s := range ss {
		slice := strings.Split(s, ":")
		if len(slice) != 2 {
			return fmt.Errorf("invalid config override syntax: %s", s)
		}
		key := slice[0]
		value := slice[1]

		cfg.Set(key, value)
	}

	return nil
}
