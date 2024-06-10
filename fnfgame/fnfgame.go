package fnfgame

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

type FNFGame struct {
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

func New(optionsPath, progressPath string) (*FNFGame, error) {
	g := new(FNFGame)
	g.windowWidth = 1280
	g.windowHeight = 720
	g.width = 1280
	g.height = 720

	g.last = time.Now()

	// Rand
	var seed1 uint64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed1); err != nil {
		return nil, err
	}
	var seed2 uint64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed2); err != nil {
		return nil, err
	}
	g.random = rand.New(rand.NewPCG(seed1, seed2))

	// Options config (config.gecfg)
	optionsConfig, err := config.New(optionsPath, true)
	if err != nil {
		return nil, err
	}
	g.optionsConfig = optionsConfig

	// Progress config (progress.gecfg)
	progressConfig, err := config.New(progressPath, false)
	if err != nil {
		return nil, err
	}
	g.progressConfig = progressConfig

	// initialize localizer
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if err := loadLocales(bundle); err != nil {
		return nil, err
	}

	locale := optionsConfig.MustGet("Locale").(string)
	slog.Info("using locale", "locale", locale)
	g.localizer = i18n.NewLocalizer(bundle, locale, "en")

	//Audio
	g.conductor = ge.NewConductor(100)
	g.sampleRate = beep.SampleRate(48000)
	g.audioMixer = &beep.Mixer{}
	g.audioResampleQuality = 4

	speaker.Init(g.sampleRate, g.sampleRate.N(time.Second/10))

	// State
	state, err := state.NewTitleState()
	if err != nil {
		return nil, err
	}
	g.curState = state

	return g, nil
}

func (g *FNFGame) Update() error {
	g.dt = time.Since(g.last).Seconds()
	g.last = time.Now()

	if err := g.curState.Update(g.dt); err != nil {
		return err
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ge.G.OptionsConfig.Toggle("Fullscreen")
		ebiten.SetFullscreen(ge.G.OptionsConfig.MustGet("Fullscreen").(bool))
		slog.Info("toggled fullscreen", "Fullscreen", ge.G.OptionsConfig.MustGet("Fullscreen").(bool))
	}

	return nil
}

func (g *FNFGame) Draw(screen *ebiten.Image) {
	g.curState.Draw(screen)

	ebitenutil.DebugPrint(screen, ge.LocalizeTmpl("FPSCounter", map[string]interface{}{
		"FPS": ebiten.ActualFPS(),
		"TPS": ebiten.ActualTPS(),
	}))
}

func (g *FNFGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	return int(math.Ceil(float64(ge.G.Width) * scale)), int(math.Ceil(float64(ge.G.Height) * scale))
}

func (g *FNFGame) InitEbiten() {
	ebiten.SetVsyncEnabled(false) // TODO: get vsync from config
	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetFullscreen(ge.G.OptionsConfig.MustGet("Fullscreen").(bool))
	slog.Info("creating window")
	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
}

func (g *FNFGame) Start() error {
	speaker.Play(g.audioMixer)
	return ebiten.RunGame(g)
}

func (g *FNFGame) Close() error {
	slog.Info("cleaning up")

	if err := g.optionsConfig.Flush(); err != nil {
		return err
	}

	if err := g.progressConfig.Flush(); err != nil {
		return err
	}

	if err := g.optionsConfig.Close(); err != nil {
		return err
	}

	if err := g.progressConfig.Close(); err != nil {
		return err
	}

	return nil
}
