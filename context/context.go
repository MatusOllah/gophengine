package context

import (
	crand "crypto/rand"
	"encoding/binary"
	"io/fs"
	"log/slog"
	"math/rand/v2"

	"github.com/BurntSushi/toml"
	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/gopxl/beep/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	input "github.com/quasilyte/ebitengine-input"
	"golang.org/x/text/language"
)

// Context holds global variables and shared game state.
type Context struct {
	WindowWidth          int
	WindowHeight         int
	Width                int
	Height               int
	AssetsFS             fs.FS
	SceneCtrl            *ge.SceneController
	InputSystem          input.System
	InputHandler         *input.Handler
	Rand                 *rand.Rand
	OptionsConfig        *config.Config
	ProgressConfig       *config.Config
	Localizer            *i18n.Localizer
	Conductor            *ge.Conductor
	SampleRate           beep.SampleRate
	AudioMixer           *audio.Mixer
	AudioResampleQuality int
	Version              string
	FNFVersion           string
}

func New(cfg *NewContextConfig) (*Context, error) {
	ctx := &Context{}
	ctx.WindowWidth = 1280
	ctx.WindowHeight = 720
	ctx.Width = 1280
	ctx.Height = 720

	ctx.AssetsFS = cfg.AssetsFS

	ctx.SceneCtrl = ge.NewStateController(nil)

	// Rand
	var seed1, seed2 uint64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed1); err != nil {
		return nil, err
	}
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed2); err != nil {
		return nil, err
	}
	ctx.Rand = rand.New(rand.NewPCG(seed1, seed2))

	// Options config (config.gecfg)
	optionsConfig, err := config.New(cfg.OptionsConfigPath, true)
	if err != nil {
		return nil, err
	}
	ctx.OptionsConfig = optionsConfig

	// Progress config (progress.gecfg)
	progressConfig, err := config.New(cfg.ProgressConfigPath, false)
	if err != nil {
		return nil, err
	}
	ctx.ProgressConfig = progressConfig

	// Controls
	ctx.InputSystem.Init(input.SystemConfig{input.AnyDevice})
	keymap, err := ge.LoadKeymapFromConfig(ctx.OptionsConfig)
	if err != nil {
		return nil, err
	}
	ctx.InputHandler = ctx.InputSystem.NewHandler(0, keymap)

	// Localizer
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if err := loadLocales(bundle, ctx.AssetsFS); err != nil {
		return nil, err
	}

	if cfg.Locale != "" {
		slog.Info("using locale", "locale", cfg.Locale)
		ctx.Localizer = i18n.NewLocalizer(bundle, cfg.Locale)
	} else {
		locale := optionsConfig.MustGet("Locale").(string)
		slog.Info("using locale", "locale", locale)
		ctx.Localizer = i18n.NewLocalizer(bundle, locale, "en")
	}

	//Audio
	ctx.Conductor = ge.NewConductor(100)
	ctx.SampleRate = beep.SampleRate(44100)
	ctx.AudioMixer = audio.NewMixer()
	ctx.AudioMixer.Master.SetVolume(ctx.OptionsConfig.MustGet("Audio.MasterVolume").(float64))
	ctx.AudioMixer.Music.SetVolume(ctx.OptionsConfig.MustGet("Audio.MusicVolume").(float64))
	ctx.AudioMixer.SFX.SetVolume(ctx.OptionsConfig.MustGet("Audio.SFXVolume").(float64))
	ctx.AudioResampleQuality = 4

	ctx.Version = cfg.Version
	ctx.FNFVersion = cfg.FNFVersion

	return ctx, nil
}

func loadLocales(bundle *i18n.Bundle, fsys fs.FS) error {
	files, err := fs.Glob(fsys, "data/locale/*.toml")
	if err != nil {
		return err
	}

	for _, file := range files {
		slog.Info("loading locale", "file", file)
		if _, err := bundle.LoadMessageFileFS(fsys, file); err != nil {
			return err
		}
	}

	return nil
}
