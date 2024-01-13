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
	"github.com/ncruces/zenity"
)

type Game struct {
	last         time.Time
	currentState state.State
}

func NewGame() (*Game, error) {
	state, err := state.NewTitleState()
	if err != nil {
		return nil, err
	}

	return &Game{
		last:         time.Now(),
		currentState: state,
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
	game.currentState.Draw(screen)

	ebitenutil.DebugPrint(screen, ge.LocalizeTmpl("FPSCounter", map[string]interface{}{
		"FPS": ebiten.ActualFPS(),
		"TPS": ebiten.ActualTPS(),
	}))
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ge.G.Width, ge.G.Height
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
	defer cleanUp()

	audio.NewContext(48000)

	fmt.Println()
	slog.Info(fmt.Sprintf("GophEngine version %s", ge.G.Version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	slog.Info(fmt.Sprintf("Friday Night Funkin' version %s", ge.G.FNFVersion))
	slog.Info("ahoj!")
	fmt.Println()

	if flagutil.MustGetBool(ge.G.FlagSet, "extract-assets") {
		if err := ge.ExtractAssets(); err != nil {
			slog.Error(err.Error())
			showError(err)
		}

		return
	}

	slog.Info("initializing game")
	beforeGameInit := time.Now()

	ebiten.SetVsyncEnabled(flagutil.MustGetBool(ge.G.FlagSet, "vsync"))
	ebiten.SetTPS(ebiten.SyncWithFPS)

	slog.Info("creating window")
	ebiten.SetWindowSize(ge.G.ScreenWidth, ge.G.ScreenHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
	if err := setIcon(); err != nil {
		slog.Error(err.Error())
		showError(err)
	}

	game, err := NewGame()
	if err != nil {
		slog.Error(err.Error())
		showError(err)
	}

	slog.Info("init game done", "time", time.Since(beforeGameInit))

	if flagutil.MustGetBool(ge.G.FlagSet, "just-init") {
		return
	}

	if err := ebiten.RunGame(game); err != nil {
		slog.Error(err.Error())
		showError(err)
	}

	cleanUp()
}

func cleanUp() {
	slog.Info("cleaning up")

	if err := ge.G.OptionsConfig.Flush(); err != nil {
		slog.Error(err.Error())
		showError(err)
	}

	if err := ge.G.ProgressConfig.Flush(); err != nil {
		slog.Error(err.Error())
		showError(err)
	}

	ge.G.OptionsConfig.Close()
	ge.G.ProgressConfig.Close()

	slog.Info("exiting")
	os.Exit(0)
}

func showError(err error) {
	if flagutil.MustGetBool(ge.G.FlagSet, "gui") {
		zenity.Error(err.Error())
	}
	os.Exit(1)
}
