package context

import (
	"fmt"
	"io/fs"
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/MatusOllah/gophengine/internal/audio/music"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/MatusOllah/gophengine/internal/controls"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/funkin"
	"github.com/MatusOllah/gophengine/internal/gui"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/goreg"
	input "github.com/quasilyte/ebitengine-input"
)

// Context holds global variables and shared game state.
type Context struct {
	AssetsFS            fs.FS
	SceneCtrl           *engine.SceneController
	InputSystem         input.System
	InputHandler        *input.Handler
	OptionsConfig       *config.Config
	ProgressConfig      *config.Config
	Conductor           *funkin.Conductor
	AudioMixer          *audio.Mixer
	FreakyMenu          *music.FreakyMenuMusic
	Weeks               *goreg.OrderedRegistry[*funkin.Week]
	StoryMenuCharacters *goreg.StandardRegistry[*funkin.StoryMenuCharacter]
}

// New creates a new [Context].
func New(opts *Options) (*Context, error) {
	ctx := &Context{}

	ctx.AssetsFS = opts.AssetsFS

	ctx.SceneCtrl = engine.NewSceneController(nil)

	// Options config
	optionsConfig, err := config.New(opts.OptionsConfigPath, true)
	if err != nil {
		return nil, fmt.Errorf("failed to load options config: %w", err)
	}
	ctx.OptionsConfig = optionsConfig

	// Progress config
	progressConfig, err := config.New(opts.ProgressConfigPath, false)
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
	locale := opts.Locale
	if locale == "" {
		locale = optionsConfig.MustGet("Locale").(string)
	}
	slog.Info("using locale", "locale", locale)
	if err := i18n.Init(opts.AssetsFS, locale); err != nil {
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

	ctx.Weeks = initWeeks()

	return ctx, nil
}
