package gophengine

import (
	crand "crypto/rand"
	"encoding/binary"
	"io/fs"
	"log/slog"
	"math"
	"math/rand/v2"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/config"
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/MatusOllah/gophengine/internal/state"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Game struct {
	windowWidth          int
	windowHeight         int
	width                int
	height               int
	random               *rand.Rand
	optionsConfig        *config.Config
	progressConfig       *config.Config
	localizer            *i18n.Localizer
	conductor            *ge.Conductor
	sampleRate           beep.SampleRate
	audioMixer           *beep.Mixer
	audioResampleQuality int
	last                 time.Time
	dt                   float64
	curState             state.State
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

func NewGame(optionsPath, progressPath string) (*Game, error) {
	game := new(Game)
	game.windowWidth = 1280
	game.windowHeight = 720
	game.width = 1280
	game.height = 720

	game.last = time.Now()

	// Rand
	var seed1 uint64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed1); err != nil {
		return nil, err
	}
	var seed2 uint64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed2); err != nil {
		return nil, err
	}
	game.random = rand.New(rand.NewPCG(seed1, seed2))

	// Options config (config.gecfg)
	optionsConfig, err := config.New(optionsPath, true)
	if err != nil {
		return nil, err
	}
	game.optionsConfig = optionsConfig

	// Progress config (progress.gecfg)
	progressConfig, err := config.New(progressPath, false)
	if err != nil {
		return nil, err
	}
	game.progressConfig = progressConfig

	// initialize localizer
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if err := loadLocales(bundle); err != nil {
		return nil, err
	}

	locale := optionsConfig.MustGet("Locale").(string)
	slog.Info("using locale", "locale", locale)
	game.localizer = i18n.NewLocalizer(bundle, locale, "en")

	//Audio
	game.conductor = ge.NewConductor(100)
	game.sampleRate = beep.SampleRate(48000)
	game.audioMixer = &beep.Mixer{}
	game.audioResampleQuality = 4

	speaker.Init(game.sampleRate, game.sampleRate.N(time.Second/10))

	// State
	state, err := state.NewTitleState()
	if err != nil {
		return nil, err
	}
	game.curState = state

	return game, nil
}

func (game *Game) Update() error {
	game.dt = time.Since(game.last).Seconds()
	game.last = time.Now()

	if err := game.curState.Update(game.dt); err != nil {
		return err
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ge.G.OptionsConfig.Toggle("Fullscreen")
		ebiten.SetFullscreen(ge.G.OptionsConfig.MustGet("Fullscreen").(bool))
		slog.Info("toggled fullscreen", "Fullscreen", ge.G.OptionsConfig.MustGet("Fullscreen").(bool))
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.curState.Draw(screen)

	ebitenutil.DebugPrint(screen, ge.LocalizeTmpl("FPSCounter", map[string]interface{}{
		"FPS": ebiten.ActualFPS(),
		"TPS": ebiten.ActualTPS(),
	}))
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	return int(math.Ceil(float64(ge.G.Width) * scale)), int(math.Ceil(float64(ge.G.Height) * scale))
}

func (game *Game) InitEbiten() {
	ebiten.SetVsyncEnabled(false) // TODO: get vsync from config
	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetFullscreen(ge.G.OptionsConfig.MustGet("Fullscreen").(bool))
	slog.Info("creating window")
	ebiten.SetWindowSize(game.windowWidth, game.windowHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
}

func (game *Game) Start() error {
	speaker.Play(game.audioMixer)
	return ebiten.RunGame(game)
}

func (game *Game) Close() error {
	slog.Info("cleaning up")

	if err := game.optionsConfig.Flush(); err != nil {
		return err
	}

	if err := game.progressConfig.Flush(); err != nil {
		return err
	}

	if err := game.optionsConfig.Close(); err != nil {
		return err
	}

	if err := game.progressConfig.Close(); err != nil {
		return err
	}

	return nil
}
