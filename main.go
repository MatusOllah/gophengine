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
	"strings"
	"time"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/flagutil"
	ge "github.com/MatusOllah/gophengine/internal/gophengine"
	"github.com/MatusOllah/gophengine/internal/state"
	"github.com/MatusOllah/slogcolor"
	"github.com/gopxl/beep/speaker"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ncruces/zenity"
	"golang.org/x/text/unicode/norm"
)

type Game struct {
	last         time.Time
	dt           float64
	currentState state.State
}

func NewGame() (*Game, error) {
	speaker.Init(ge.G.SampleRate, ge.G.SampleRate.N(time.Second/10))

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
	game.dt = time.Since(game.last).Seconds()
	game.last = time.Now()

	if err := game.currentState.Update(game.dt); err != nil {
		return err
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.currentState.Draw(screen)

	ebitenutil.DebugPrint(screen, norm.NFC.String(ge.LocalizeTmpl("FPSCounter", map[string]interface{}{
		"FPS": ebiten.ActualFPS(),
		"TPS": ebiten.ActualTPS(),
	})))
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

func getLogLevel() slog.Leveler {
	switch s := strings.ToLower(flagutil.MustGetString(ge.FlagSet, "log-level")); s {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		panic("unknown log level: " + s)
	}
}

func main() {
	if err := ge.InitFlags(); err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, &slogcolor.Options{
		Level:       getLogLevel(),
		TimeFormat:  time.DateTime,
		SrcFileMode: slogcolor.ShortFile,
	})))
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	if err := ge.InitGlobal(); err != nil {
		panic(err)
	}
	defer cleanUp()
	slog.Info("finished early initialization")

	slog.Info(fmt.Sprintf("GophEngine version %s", ge.G.Version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	slog.Info(fmt.Sprintf("Friday Night Funkin' version %s", ge.G.FNFVersion))
	slog.Info("ahoj!")

	if flagutil.MustGetBool(ge.FlagSet, "extract-assets") {
		if err := ge.ExtractAssets(); err != nil {
			slog.Error(err.Error())
			showError(err)
		}

		return
	}

	slog.Info("initializing game")
	beforeGameInit := time.Now()

	var dlg zenity.ProgressDialog
	if flagutil.MustGetBool(ge.FlagSet, "gui") {
		dlg, _ = zenity.Progress(zenity.Title(ge.Localize("InitGameDialogTitle")), zenity.Pulsate())
		dlg.Text(ge.Localize("InitGameDialogText"))
	}

	ebiten.SetVsyncEnabled(flagutil.MustGetBool(ge.FlagSet, "vsync"))
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

	if flagutil.MustGetBool(ge.FlagSet, "gui") {
		dlg.Complete()
		dlg.Close()
	}
	slog.Info("init game done", "time", time.Since(beforeGameInit))

	if flagutil.MustGetBool(ge.FlagSet, "just-init") {
		return
	}

	speaker.Play(ge.G.Mixer)
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
	if flagutil.MustGetBool(ge.FlagSet, "gui") {
		zenity.Error(err.Error())
	}
	os.Exit(1)
}
