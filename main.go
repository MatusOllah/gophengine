package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/silbinarywolf/preferdiscretegpu"
	"github.com/ztrue/tracerr"
)

type Game struct {
	Last         time.Time
	CurrentState State
}

func NewGame() (*Game, error) {
	state, err := NewTitleState()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &Game{
		Last:         time.Now(),
		CurrentState: state,
	}, nil
}

func (game *Game) Update() error {
	dt := time.Since(game.Last).Seconds()
	game.Last = time.Now()

	if err := game.CurrentState.Update(dt); err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.CurrentState.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v\nTPS: %v", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
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

	if err := initGlobal(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	log.Info().Msgf("GophEngine version %s", g.Version)
	log.Info().Msgf("Go version %s", runtime.Version())
	log.Info().Msgf("Friday Night Funkin' version %s", g.FNFVersion)
	log.Info().Msg("ahoj!")

	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	log.Info().Msg("creating window")
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")

	game, err := NewGame()
	if err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := ebiten.RunGame(game); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := g.ConfigSave.Flush(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := g.ProgressSave.Flush(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := g.ConfigSave.Close(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}

	if err := g.ProgressSave.Close(); err != nil {
		tracerr.Print(err)
		os.Exit(1)
	}
}
