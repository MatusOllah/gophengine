package fnfgame

import (
	"fmt"
	"log/slog"
	"runtime"
	"time"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/internal/scene"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.design/x/clipboard"
)

type FNFGame struct {
	ctx  *context.Context
	last time.Time
	dt   float64
}

func New(ctx *context.Context) (*FNFGame, error) {
	g := new(FNFGame)
	g.ctx = ctx

	g.last = time.Now()

	// State
	if err := g.ctx.SceneCtrl.SwitchScene(scene.NewTitleScene(ctx)); err != nil {
		return nil, fmt.Errorf("fnfgame New: error initializing TitleState: %w", err)
	}

	if runtime.GOARCH != "wasm" {
		if err := clipboard.Init(); err != nil {
			return nil, fmt.Errorf("fnfgame New: failed to initialize clipboard: %w", err)
		}
	}

	return g, nil
}

func (g *FNFGame) Update() error {
	g.dt = time.Since(g.last).Seconds()
	g.last = time.Now()

	g.ctx.InputSystem.UpdateWithDelta(g.dt)

	if err := g.ctx.SceneCtrl.Update(g.dt); err != nil {
		return fmt.Errorf("error updating state: %w", err)
	}

	if g.ctx.InputHandler.ActionIsJustPressed(ge.ActionFullscreen) {
		g.ctx.OptionsConfig.Toggle("Fullscreen")
		ebiten.SetFullscreen(g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
		slog.Info("toggled fullscreen", "Fullscreen", g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
	}

	return nil
}

func (g *FNFGame) Draw(screen *ebiten.Image) {
	g.ctx.SceneCtrl.Draw(screen)

	ebitenutil.DebugPrint(screen, i18n.LT("FPSCounter", map[string]interface{}{
		"FPS": ebiten.ActualFPS(),
		"TPS": ebiten.ActualTPS(),
	}))
}

func (g *FNFGame) Layout(_, _ int) (int, int) {
	return g.ctx.Width, g.ctx.Height
}

func (g *FNFGame) InitEbiten() {
	ebiten.SetVsyncEnabled(false) // TODO: get vsync from config
	ebiten.SetFullscreen(g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
	slog.Info("creating window")
	ebiten.SetWindowSize(g.ctx.WindowWidth, g.ctx.WindowHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
}

func (g *FNFGame) initAudio() error {
	bufSize := 4096
	if runtime.GOARCH == "wasm" {
		bufSize = 8194
	}
	slog.Info("initializing audio", "sampleRate", g.ctx.SampleRate, "bufferSize", bufSize)
	if err := speaker.Init(g.ctx.SampleRate, bufSize); err != nil {
		return fmt.Errorf("fnfgame Start: error initializing audio: %w", err)
	}

	if g.ctx.OptionsConfig.MustGet("Audio.DownmixToMono").(bool) {
		speaker.Play(effects.Mono(g.ctx.AudioMixer))
	} else {
		speaker.Play(g.ctx.AudioMixer)
	}

	return nil
}

func (g *FNFGame) Start() error {
	if err := g.initAudio(); err != nil {
		return err
	}
	if err := ebiten.RunGame(g); err != nil {
		return fmt.Errorf("fnfgame Start: error running game: %w", err)
	}
	return nil
}

func (g *FNFGame) StartWithOptions(opts *ebiten.RunGameOptions) error {
	if err := g.initAudio(); err != nil {
		return err
	}
	if err := ebiten.RunGameWithOptions(g, opts); err != nil {
		return fmt.Errorf("fnfgame StartWithOptions: error running game: %w", err)
	}
	return nil
}

func (g *FNFGame) Close() error {
	slog.Info("cleaning up")

	if err := g.ctx.SceneCtrl.Close(); err != nil {
		return fmt.Errorf("fnfgame Close: error closing SceneCtrl: %w", err)
	}

	if err := g.ctx.OptionsConfig.Close(); err != nil {
		return fmt.Errorf("fnfgame Close: error closing OptionsConfig: %w", err)
	}

	if err := g.ctx.ProgressConfig.Close(); err != nil {
		return fmt.Errorf("fnfgame Close: error closing ProgressConfig: %w", err)
	}

	return nil
}
