package fnfgame

import (
	"fmt"
	"log/slog"
	"time"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/MatusOllah/gophengine/internal/scene"
	"github.com/gopxl/beep/speaker"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

	ebitenutil.DebugPrint(screen, i18nutil.LocalizeTmpl(g.ctx.Localizer, "FPSCounter", map[string]interface{}{
		"FPS": ebiten.ActualFPS(),
		"TPS": ebiten.ActualTPS(),
	}))
}

func (g *FNFGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ctx.Width, g.ctx.Height
}

func (g *FNFGame) InitEbiten() {
	ebiten.SetVsyncEnabled(false) // TODO: get vsync from config
	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetFullscreen(g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
	slog.Info("creating window")
	ebiten.SetWindowSize(g.ctx.WindowWidth, g.ctx.WindowHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
}

func (g *FNFGame) Start() error {
	if err := speaker.Init(g.ctx.SampleRate, g.ctx.SampleRate.N(time.Second/10)); err != nil {
		return fmt.Errorf("fnfgame Start: error initializing speaker: %w", err)
	}
	speaker.Play(g.ctx.AudioMixer)
	if err := ebiten.RunGame(g); err != nil {
		return fmt.Errorf("fnfgame Start: error running game: %w", err)
	}
	return nil
}

func (g *FNFGame) StartWithOptions(opts *ebiten.RunGameOptions) error {
	if err := speaker.Init(g.ctx.SampleRate, g.ctx.SampleRate.N(time.Second/10)); err != nil {
		return fmt.Errorf("fnfgame StartWithOptions: error initializing speaker: %w", err)
	}
	speaker.Play(g.ctx.AudioMixer)
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
