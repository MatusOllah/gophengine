package context

import (
	crand "crypto/rand"
	"encoding/binary"
	"io/fs"
	"log/slog"
	"math/rand/v2"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/internal/audio"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/gopxl/beep/v2"
	input "github.com/quasilyte/ebitengine-input"
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
	locale := cfg.Locale
	if locale == "" {
		locale = optionsConfig.MustGet("Locale").(string)
	}
	slog.Info("using locale", "locale", locale)
	if err := i18n.Init(cfg.AssetsFS, locale); err != nil {
		return nil, err
	}

	//Audio
	ctx.Conductor = ge.NewConductor(100)
	ctx.SampleRate = beep.SampleRate(44100)
	ctx.AudioMixer = audio.NewMixer()
	ctx.AudioMixer.Master.SetVolume(ctx.OptionsConfig.MustGet("Audio.MasterVolume").(float64))
	ctx.AudioMixer.SFX.SetVolume(ctx.OptionsConfig.MustGet("Audio.SFXVolume").(float64))
	ctx.AudioMixer.Music.SetVolume(ctx.OptionsConfig.MustGet("Audio.MusicVolume").(float64))
	ctx.AudioMixer.Music_Instrumental.SetVolume(ctx.OptionsConfig.MustGet("Audio.InstVolume").(float64))
	ctx.AudioMixer.Music_Voices.SetVolume(ctx.OptionsConfig.MustGet("Audio.VoicesVolume").(float64))
	ctx.AudioResampleQuality = 4

	ctx.Version = cfg.Version
	ctx.FNFVersion = cfg.FNFVersion

	return ctx, nil
}
