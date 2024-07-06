package context

import (
	crand "crypto/rand"
	"encoding/binary"
	"io/fs"
	"log/slog"
	"math/rand/v2"

	"github.com/BurntSushi/toml"
	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/gopxl/beep"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Context holds global variables and shared game state.
type Context struct {
	WindowWidth          int
	WindowHeight         int
	Width                int
	Height               int
	AssetsFS             fs.FS
	StateController      *ge.StateController
	Controls             *ge.Controls
	Rand                 *rand.Rand
	OptionsConfig        *config.Config
	ProgressConfig       *config.Config
	Localizer            *i18n.Localizer
	Conductor            *ge.Conductor
	SampleRate           beep.SampleRate
	AudioMixer           *beep.Mixer
	AudioResampleQuality int
	Version              string
}

func New(cfg *NewContextConfig) (*Context, error) {
	ctx := &Context{}
	ctx.WindowWidth = 1280
	ctx.WindowHeight = 720
	ctx.Width = 1280
	ctx.Height = 720

	ctx.AssetsFS = cfg.AssetsFS

	ctx.StateController = ge.NewStateController(nil)

	// Rand
	var seed1 uint64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed1); err != nil {
		return nil, err
	}
	var seed2 uint64
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
	ctl, err := ge.GetControlsFromConfig(ctx.OptionsConfig)
	if err != nil {
		return nil, err
	}
	ctx.Controls = ctl

	// Localizer
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if err := loadLocales(bundle, ctx.AssetsFS); err != nil {
		return nil, err
	}

	locale := optionsConfig.MustGet("Locale").(string)
	slog.Info("using locale", "locale", locale)
	ctx.Localizer = i18n.NewLocalizer(bundle, locale, "en")

	//Audio
	ctx.Conductor = ge.NewConductor(100)
	ctx.SampleRate = beep.SampleRate(44100)
	ctx.AudioMixer = &beep.Mixer{}
	ctx.AudioResampleQuality = 4

	ctx.Version = cfg.Version

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
