package main

import (
	"bytes"
	"fmt"
	"image"
	"io/fs"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/MatusOllah/gophengine/assets"
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/MatusOllah/gophengine/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/silbinarywolf/preferdiscretegpu"
	"github.com/ztrue/tracerr"
)

type Game struct {
	last         time.Time
	currentState state.State
}

func NewGame() (*Game, error) {
	state, err := state.NewTitleState()
	if err != nil {
		return nil, tracerr.Wrap(err)
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
		return tracerr.Wrap(err)
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.currentState.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"FPS: %v\nTPS: %v",
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
		return tracerr.Wrap(err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return tracerr.Wrap(err)
	}

	ebiten.SetWindowIcon([]image.Image{img})

	return nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).With().Caller().Logger()

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	if err := ge.InitGlobal(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	log.Info().Msgf("GophEngine version %s", ge.G.Version)
	log.Info().Msgf("Go version %s", runtime.Version())
	log.Info().Msgf("Friday Night Funkin' version %s", ge.G.FNFVersion)
	log.Info().Msg("ahoj!")

	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	log.Info().Msg("creating window")
	ebiten.SetWindowSize(ge.G.ScreenWidth, ge.G.ScreenHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
	if err := setIcon(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	game, err := NewGame()
	if err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := ebiten.RunGame(game); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := ge.G.ConfigSave.Flush(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := ge.G.ProgressSave.Flush(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := ge.G.ConfigSave.Close(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := ge.G.ProgressSave.Close(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	runtime.GC()
	os.Exit(0)
}
