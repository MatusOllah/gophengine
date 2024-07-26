package fnfgame

import (
	"log/slog"
	"time"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/MatusOllah/gophengine/internal/state"
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
	state, err := state.NewTitleState(ctx)
	if err != nil {
		return nil, err
	}
	g.ctx.StateController.SwitchState(state)

	return g, nil
}

func (g *FNFGame) Update() error {
	g.dt = time.Since(g.last).Seconds()
	g.last = time.Now()

	g.ctx.InputSystem.UpdateWithDelta(g.dt)

	if err := g.ctx.StateController.Update(g.dt); err != nil {
		return err
	}

	if g.ctx.InputHandler.ActionIsJustPressed(ge.ActionFullscreen) {
		g.ctx.OptionsConfig.Toggle("Fullscreen")
		ebiten.SetFullscreen(g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
		slog.Info("toggled fullscreen", "Fullscreen", g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
	}

	return nil
}

func (g *FNFGame) Draw(screen *ebiten.Image) {
	g.ctx.StateController.Draw(screen)

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
		return err
	}
	speaker.Play(g.ctx.AudioMixer)
	return ebiten.RunGame(g)
}

func (g *FNFGame) StartWithOptions(opts *ebiten.RunGameOptions) error {
	if err := speaker.Init(g.ctx.SampleRate, g.ctx.SampleRate.N(time.Second/10)); err != nil {
		return err
	}
	speaker.Play(g.ctx.AudioMixer)
	return ebiten.RunGameWithOptions(g, opts)
}

func (g *FNFGame) Close() error {
	slog.Info("cleaning up")

	if err := g.ctx.OptionsConfig.Close(); err != nil {
		return err
	}

	if err := g.ctx.ProgressConfig.Close(); err != nil {
		return err
	}

	return nil
}
