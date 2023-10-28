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
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/MatusOllah/gophengine/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jessevdk/go-flags"
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
	if _, err := flags.NewParser(&ge.Options, flags.HelpFlag|flags.IgnoreUnknown|flags.PassDoubleDash).Parse(); err != nil {
		panic(err)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	beforeInit := time.Now()

	if err := ge.InitGlobal(); err != nil {
		panic(err)
	}

	slog.Info(fmt.Sprintf("init took %v", time.Since(beforeInit)))

	fmt.Println()
	slog.Info(fmt.Sprintf("GophEngine version %s", ge.G.Version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	slog.Info(fmt.Sprintf("Friday Night Funkin' version %s", ge.G.FNFVersion))
	slog.Info("ahoj!")
	fmt.Println()

	if ge.Options.ExtractAssets {
		err := ge.ExtractAssets()
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
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

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	if err := ge.G.ConfigSave.Flush(); err != nil {
		panic(err)
	}

	if err := ge.G.ProgressSave.Flush(); err != nil {
		panic(err)
	}

	ge.G.ConfigSave.Close()
	ge.G.ProgressSave.Close()

	runtime.GC()
	os.Exit(0)
}
