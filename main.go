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
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	beforeInit := time.Now()

	if err := ge.InitGlobal(); err != nil {
		panic(err)
	}

	log.Info().Msgf("init took %v", time.Since(beforeInit))

	fmt.Println()
	log.Info().Msgf("GophEngine version %s", ge.G.Version)
	log.Info().Msgf("Go version %s", runtime.Version())
	log.Info().Msgf("Friday Night Funkin' version %s", ge.G.FNFVersion)
	log.Info().Msg("ahoj!")
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

	log.Info().Msg("creating window")
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
