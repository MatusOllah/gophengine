package context

import (
	"fmt"
	"io/fs"
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/MatusOllah/gophengine/internal/controls"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/funkin"
	"github.com/MatusOllah/gophengine/internal/gui"
	"github.com/MatusOllah/gophengine/internal/i18n"
	input "github.com/quasilyte/ebitengine-input"
)

// Context holds global variables and shared game state.
type Context struct {
	AssetsFS       fs.FS
	SceneCtrl      *engine.SceneController
	InputSystem    input.System
	InputHandler   *input.Handler
	OptionsConfig  *config.Config
	ProgressConfig *config.Config
	Conductor      *funkin.Conductor
	AudioMixer     *audio.Mixer
	Version        string
	FNFVersion     string
	Weeks          []funkin.Week
}

// New creates a new [Context].
func New(cfg *NewContextConfig) (*Context, error) {
	ctx := &Context{}

	ctx.AssetsFS = cfg.AssetsFS

	ctx.SceneCtrl = engine.NewSceneController(nil)

	// Options config
	optionsConfig, err := config.New(cfg.OptionsConfigPath, true)
	if err != nil {
		return nil, fmt.Errorf("failed to load options config: %w", err)
	}
	ctx.OptionsConfig = optionsConfig

	// Progress config
	progressConfig, err := config.New(cfg.ProgressConfigPath, false)
	if err != nil {
		return nil, fmt.Errorf("failed to load progress config: %w", err)
	}
	ctx.ProgressConfig = progressConfig

	// Controls
	ctx.InputSystem.Init(input.SystemConfig{DevicesEnabled: input.AnyDevice})
	keymap, err := controls.LoadKeymapFromConfig(ctx.OptionsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to load keymap: %w", err)
	}
	ctx.InputHandler = ctx.InputSystem.NewHandler(0, keymap)

	// Localizer
	locale := cfg.Locale
	if locale == "" {
		locale = optionsConfig.MustGet("Locale").(string)
	}
	slog.Info("using locale", "locale", locale)
	if err := i18n.Init(cfg.AssetsFS, locale); err != nil {
		return nil, fmt.Errorf("failed to initialize i18n: %w", err)
	}

	// GUI
	if err := gui.Init(ctx.AssetsFS); err != nil {
		return nil, fmt.Errorf("failed to initialize gui resources: %w", err)
	}

	//Audio
	ctx.Conductor = funkin.NewConductor(100)
	ctx.AudioMixer = audio.NewMixer()
	ctx.AudioMixer.Master.SetVolume(ctx.OptionsConfig.MustGet("Audio.MasterVolume").(float64))
	ctx.AudioMixer.SFX.SetVolume(ctx.OptionsConfig.MustGet("Audio.SFXVolume").(float64))
	ctx.AudioMixer.Music.SetVolume(ctx.OptionsConfig.MustGet("Audio.MusicVolume").(float64))
	ctx.AudioMixer.Music_Instrumental.SetVolume(ctx.OptionsConfig.MustGet("Audio.InstVolume").(float64))
	ctx.AudioMixer.Music_Voices.SetVolume(ctx.OptionsConfig.MustGet("Audio.VoicesVolume").(float64))

	ctx.Version = cfg.Version
	ctx.FNFVersion = cfg.FNFVersion

	return ctx, nil
}
