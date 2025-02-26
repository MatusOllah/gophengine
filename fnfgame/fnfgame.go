package fnfgame

import (
	"fmt"
	"io/fs"
	"log/slog"
	"runtime"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/controls"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/internal/scene"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.design/x/clipboard"
)

type FNFGame struct {
	ctx *context.Context

	protanShader    *ebiten.Shader
	deuteranShader  *ebiten.Shader
	tritanShader    *ebiten.Shader
	grayscaleShader *ebiten.Shader

	colorblindImage *ebiten.Image
}

// New creates a new [FNFGame].
func New(ctx *context.Context) (*FNFGame, error) {
	g := new(FNFGame)
	g.ctx = ctx

	// State
	if err := g.ctx.SceneCtrl.SwitchScene(scene.NewTitleScene(ctx)); err != nil {
		return nil, fmt.Errorf("fnfgame New: error initializing TitleState: %w", err)
	}

	if runtime.GOARCH != "wasm" {
		if err := clipboard.Init(); err != nil {
			return nil, fmt.Errorf("fnfgame New: failed to initialize clipboard: %w", err)
		}
	}

	g.colorblindImage = ebiten.NewImage(engine.GameWidth, engine.GameHeight)

	return g, nil
}

func (g *FNFGame) Update() error {
	g.ctx.InputSystem.Update()

	if err := g.ctx.SceneCtrl.Update(); err != nil {
		return fmt.Errorf("fnfgame Update: error updating state: %w", err)
	}

	if g.ctx.InputHandler.ActionIsJustPressed(controls.ActionFullscreen) {
		g.ctx.OptionsConfig.Toggle("Fullscreen")
		ebiten.SetFullscreen(g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
		slog.Info("toggled fullscreen", "Fullscreen", g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
	}

	return nil
}

func (g *FNFGame) Draw(screen *ebiten.Image) {
	g.ctx.SceneCtrl.Draw(screen)

	if g.ctx.OptionsConfig.MustGet("Graphics.EnableFPSCounter").(bool) {
		ebitenutil.DebugPrint(screen, i18n.LT("FPSCounter", map[string]interface{}{
			"FPS": ebiten.ActualFPS(),
			"TPS": ebiten.ActualTPS(),
		}))
	}
}

func (g *FNFGame) loadShader(path string) (*ebiten.Shader, error) {
	b, err := fs.ReadFile(g.ctx.AssetsFS, path)
	if err != nil {
		return nil, fmt.Errorf("failed to read shader bytes: %w", err)
	}

	shader, err := ebiten.NewShader(b)
	if err != nil {
		return nil, fmt.Errorf("failed to compile shader: %w", err)
	}

	return shader, nil
}

func (g *FNFGame) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	// Colorblind filter
	g.colorblindImage.DrawImage(offscreen, nil)
	var err error
	switch engine.ColorblindFilter(g.ctx.OptionsConfig.MustGet("Graphics.ColorblindFilter").(int)) {
	case engine.ColorblindNone:
		break
	case engine.ColorblindProtanopia:
		if g.protanShader == nil {
			g.protanShader, err = g.loadShader("shaders/colorblind/protanopia.kage")
			if err != nil {
				panic(fmt.Errorf("failed to load protanopia shader: %w", err))
			}
		}
		offscreen.DrawRectShader(offscreen.Bounds().Dx(), offscreen.Bounds().Dy(), g.protanShader, &ebiten.DrawRectShaderOptions{
			Blend:  ebiten.BlendCopy,
			Images: [4]*ebiten.Image{g.colorblindImage},
		})
	case engine.ColorblindDeuteranopia:
		if g.deuteranShader == nil {
			g.deuteranShader, err = g.loadShader("shaders/colorblind/deuteranopia.kage")
			if err != nil {
				panic(fmt.Errorf("failed to load deuteranopia shader: %w", err))
			}
		}
		offscreen.DrawRectShader(offscreen.Bounds().Dx(), offscreen.Bounds().Dy(), g.deuteranShader, &ebiten.DrawRectShaderOptions{
			Blend:  ebiten.BlendCopy,
			Images: [4]*ebiten.Image{g.colorblindImage},
		})
	case engine.ColorblindTritanopia:
		if g.tritanShader == nil {
			g.tritanShader, err = g.loadShader("shaders/colorblind/tritanopia.kage")
			if err != nil {
				panic(fmt.Errorf("failed to load tritanopia shader: %w", err))
			}
		}
		offscreen.DrawRectShader(offscreen.Bounds().Dx(), offscreen.Bounds().Dy(), g.tritanShader, &ebiten.DrawRectShaderOptions{
			Blend:  ebiten.BlendCopy,
			Images: [4]*ebiten.Image{g.colorblindImage},
		})
	case engine.ColorblindGrayscale:
		if g.grayscaleShader == nil {
			g.grayscaleShader, err = g.loadShader("shaders/colorblind/grayscale.kage")
			if err != nil {
				panic(fmt.Errorf("failed to load grayscale shader: %w", err))
			}
		}
		offscreen.DrawRectShader(offscreen.Bounds().Dx(), offscreen.Bounds().Dy(), g.grayscaleShader, &ebiten.DrawRectShaderOptions{
			Blend:  ebiten.BlendCopy,
			Images: [4]*ebiten.Image{g.colorblindImage},
		})
	}
	g.colorblindImage.Clear()

	// Upscaling
	switch engine.Upscaling(g.ctx.OptionsConfig.MustGet("Graphics.UpscaleMethod").(int)) {
	case engine.UpscaleNearest:
		screen.DrawImage(offscreen, &ebiten.DrawImageOptions{
			Blend:  ebiten.BlendCopy,
			Filter: ebiten.FilterNearest,
			GeoM:   geoM,
		})
	case engine.UpscaleLinear:
		screen.DrawImage(offscreen, &ebiten.DrawImageOptions{
			Blend:  ebiten.BlendCopy,
			Filter: ebiten.FilterLinear,
			GeoM:   geoM,
		})
	}
}

func (g *FNFGame) Layout(_, _ int) (int, int) {
	return engine.GameWidth, engine.GameHeight
}

func (g *FNFGame) InitEbiten() {
	ebiten.SetVsyncEnabled(g.ctx.OptionsConfig.MustGet("Graphics.EnableVSync").(bool))
	ebiten.SetFullscreen(g.ctx.OptionsConfig.MustGet("Fullscreen").(bool))
	slog.Info("creating window")
	ebiten.SetWindowSize(g.ctx.WindowWidth, g.ctx.WindowHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}

func (g *FNFGame) initAudio() error {
	sr := beep.SampleRate(44100)
	bufSize := 4096
	if runtime.GOARCH == "wasm" {
		bufSize = 8194
	}

	slog.Info("initializing audio", "sampleRate", sr, "bufferSize", bufSize)
	if err := speaker.Init(sr, bufSize); err != nil {
		return fmt.Errorf("error initializing speaker: %w", err)
	}

	if g.ctx.OptionsConfig.MustGet("Audio.DownmixToMono").(bool) {
		speaker.Play(effects.Mono(g.ctx.AudioMixer))
	} else {
		speaker.Play(g.ctx.AudioMixer)
	}

	return nil
}

// Start starts the game.
func (g *FNFGame) Start() error {
	if err := g.initAudio(); err != nil {
		return fmt.Errorf("fnfgame Start: error initializing audio: %w", err)
	}
	if err := ebiten.RunGame(g); err != nil {
		return fmt.Errorf("fnfgame Start: error running game: %w", err)
	}
	return nil
}

func (g *FNFGame) StartWithOptions(opts *ebiten.RunGameOptions) error {
	if err := g.initAudio(); err != nil {
		return fmt.Errorf("fnfgame StartWithOptions: error initializing audio: %w", err)
	}
	if err := ebiten.RunGameWithOptions(g, opts); err != nil {
		return fmt.Errorf("fnfgame StartWithOptions: error running game: %w", err)
	}
	return nil
}

// Close cleans up resources and closes the game.
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
