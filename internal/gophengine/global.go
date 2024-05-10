package gophengine

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"io/fs"
	"log/slog"
	"math/rand/v2"

	"github.com/BurntSushi/toml"
	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/MatusOllah/gophengine/internal/flagutil"
	"github.com/gopxl/beep"
	"github.com/nicksnyder/go-i18n/v2/i18n"
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
	Localizer      *i18n.Localizer
	SampleRate     beep.SampleRate
	Mixer          *beep.Mixer

	/*
		This specifies the quality of the resampling process. Higher quality implies worse performance. Values below 1 or above 64 are invalid and Resample will panic. Here's a table for deciding which quality to pick.

		quality | use case
		--------|---------
		1       | very high performance, on-the-fly resampling, low quality
		3-4     | good performance, on-the-fly resampling, good quality
		6       | higher CPU usage, usually not suitable for on-the-fly resampling, very good quality
		>6      | even higher CPU usage, for offline resampling, very good quality

		Sane quality values are usually below 16. Higher values will consume too much CPU, giving negligible quality improvements.
	*/
	ResampleQuality int
}

var G *Global

func InitGlobal() error {
	optionsPath := flagutil.MustGetString(FlagSet, "config")
	if flagutil.MustGetBool(FlagSet, "portable") {
		optionsPath = "GophEngine/config.gecfg"
	}
	slog.Info(fmt.Sprintf("using config file %s", optionsPath))

	optionsConfig, err := config.New(optionsPath, true)
	if err != nil {
		return err
	}

	if flagutil.MustGetBool(FlagSet, "config-load-defaults") {
		slog.Info("loading defaults")
		config.LoadDefaultOptions(optionsConfig)
	}

	progressPath := flagutil.MustGetString(FlagSet, "progress")
	if flagutil.MustGetBool(FlagSet, "portable") {
		progressPath = "GophEngine/progress.gecfg"
	}
	slog.Info(fmt.Sprintf("using progress file %s", progressPath))

	progressConfig, err := config.New(progressPath, false)
	if err != nil {
		return err
	}

	var seed int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		return err
	}
	rand := rand.New(xorshift1024star.NewSource(seed))

	// initialize localizer
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if err := loadLocales(bundle); err != nil {
		return err
	}

	locale := optionsConfig.MustGet("Locale").(string)
	slog.Info("using locale", "locale", locale)
	localizer := i18n.NewLocalizer(bundle, locale, "en")

	G = &Global{
		Rand:            rand,
		Version:         "1.0",
		FNFVersion:      "0.2.7.1",
		ScreenWidth:     1280,
		ScreenHeight:    720,
		Width:           1280,
		Height:          720,
		OptionsConfig:   optionsConfig,
		ProgressConfig:  progressConfig,
		Conductor:       NewConductor(100),
		Localizer:       localizer,
		SampleRate:      beep.SampleRate(48000),
		Mixer:           &beep.Mixer{},
		ResampleQuality: 4,
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
