package main

import (
	"bytes"
	"fmt"
	"image"
	"io/fs"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/flagutil"
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/MatusOllah/gophengine/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	last         time.Time
	currentState state.State
	shader       *ebiten.Shader
}

func NewGame() (*Game, error) {
	state, err := state.NewTitleState()
	if err != nil {
		return nil, err
	}

	slog.Info("Compiling shaders")
	shaderBytes, err := assets.FS.ReadFile("data/shaders/shader.kage")
	if err != nil {
		return nil, err
	}

	shader, err := ebiten.NewShader(shaderBytes)
	if err != nil {
		return nil, err
	}
	slog.Info("done compiling shaders")

	return &Game{
		last:         time.Now(),
		currentState: state,
		shader:       shader,
	}, nil
}

func (game *Game) Update() error {
	dt := time.Since(game.last).Seconds()
	game.last = time.Now()

	if err := game.currentState.Update(dt); err != nil {
		return err
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
	game.currentState.Draw(img)

	// shader
	op := &ebiten.DrawRectShaderOptions{}
	op.Images[0] = img
	screen.DrawRectShader(screen.Bounds().Dx(), screen.Bounds().Dy(), game.shader, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"FPS: %.2f\nTPS: %.2f",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
	))
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ge.G.ScreenWidth, ge.G.ScreenHeight
}

func setIcon() error {
	data, err := fs.ReadFile(assets.FS, "icon.png")
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	ebiten.SetWindowIcon([]image.Image{img})

	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	if err := ge.InitGlobal(); err != nil {
		panic(err)
	}

	audio.NewContext(48000)

	fmt.Println()
	slog.Info(fmt.Sprintf("GophEngine version %s", ge.G.Version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	slog.Info(fmt.Sprintf("Friday Night Funkin' version %s", ge.G.FNFVersion))
	slog.Info("ahoj!")
	fmt.Println()

	if flagutil.MustGetBool(ge.G.FlagSet, "extract-assets") {
		err := ge.ExtractAssets()
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	slog.Info("initializing game")
	beforeGameInit := time.Now()

	ebiten.SetVsyncEnabled(flagutil.MustGetBool(ge.G.FlagSet, "vsync"))
	ebiten.SetTPS(ebiten.SyncWithFPS)

	slog.Info("creating window")
	ebiten.SetWindowSize(ge.G.ScreenWidth, ge.G.ScreenHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
	if err := setIcon(); err != nil {
		panic(err)
	}

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	slog.Info("init game done", "time", time.Since(beforeGameInit))

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	cleanUp()
}

func cleanUp() {
	slog.Info("cleaning up")

	if err := ge.G.OptionsConfig.Flush(); err != nil {
		panic(err)
	}

	if err := ge.G.ProgressConfig.Flush(); err != nil {
		panic(err)
	}

	ge.G.OptionsConfig.Close()
	ge.G.ProgressConfig.Close()

	os.Exit(0)
}
