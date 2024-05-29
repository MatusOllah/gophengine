package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/fs"
	"log/slog"
	"math"
	"os"
	"path/filepath"
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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ge.G.OptionsConfig.Toggle("Fullscreen")
		ebiten.SetFullscreen(ge.G.OptionsConfig.MustGet("Fullscreen").(bool))
		slog.Info("toggled fullscreen", "Fullscreen", ge.G.OptionsConfig.MustGet("Fullscreen").(bool))
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.currentState.Draw(screen)

	// kvôli niečomu diakritika nefunguje cez ten default Ebitengine font...
	// preto tu je to norm.NFC.String()
	ebitenutil.DebugPrint(screen, norm.NFC.String(ge.LocalizeTmpl("FPSCounter", map[string]interface{}{
		"FPS": ebiten.ActualFPS(),
		"TPS": ebiten.ActualTPS(),
	})))
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	return int(math.Ceil(float64(ge.G.Width) * scale)), int(math.Ceil(float64(ge.G.Height) * scale))
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
	case "":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		panic(fmt.Sprintf("invalid log level: \"%s\"; should be one of \"debug\", \"info\", \"warn\", \"error\"", s))
	}
}

func getLogFilePath() string {
	return filepath.Join(os.TempDir(), "GophEngine", "logs", time.Now().Format("2006-01-02_15-04-05.log"))
}

// main func here
func run() error {
	if err := ge.InitGlobal(); err != nil {
		return err
	}
	slog.Info("finished early initialization")

	slog.Info(fmt.Sprintf("GophEngine version %s", ge.G.Version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	slog.Info(fmt.Sprintf("Friday Night Funkin' version %s", ge.G.FNFVersion))
	slog.Info("ahoj!")

	if flagutil.MustGetBool(ge.FlagSet, "extract-assets") {
		if err := ge.ExtractAssets(); err != nil {
			return err
		}

		return nil
	}

	slog.Info("initializing game")
	beforeGameInit := time.Now()

	var dlg zenity.ProgressDialog
	if flagutil.MustGetBool(ge.FlagSet, "gui") {
		dlg, _ = zenity.Progress(zenity.Pulsate(), zenity.Title(ge.Localize("InitGameDialogTitle")), zenity.NoCancel())
		dlg.Text(ge.Localize("InitGameDialogText"))
	}

	ebiten.SetVsyncEnabled(flagutil.MustGetBool(ge.FlagSet, "vsync"))
	ebiten.SetTPS(ebiten.SyncWithFPS)

	ebiten.SetFullscreen(ge.G.OptionsConfig.MustGet("Fullscreen").(bool))

	slog.Info("creating window")
	ebiten.SetWindowSize(ge.G.ScreenWidth, ge.G.ScreenHeight)
	ebiten.SetWindowTitle("Friday Night Funkin': GophEngine")
	if err := setIcon(); err != nil {
		return err
	}

	game, err := NewGame()
	if err != nil {
		return err
	}

	if flagutil.MustGetBool(ge.FlagSet, "gui") {
		dlg.Complete()
		dlg.Close()
	}
	slog.Info("init game done", "time", time.Since(beforeGameInit))

	if flagutil.MustGetBool(ge.FlagSet, "just-init") {
		return nil
	}

	speaker.Play(ge.G.Mixer)
	if err := ebiten.RunGame(game); err != nil {
		return err
	}

	slog.Info("cleaning up")

	if err := ge.G.OptionsConfig.Flush(); err != nil {
		return err
	}

	if err := ge.G.ProgressConfig.Flush(); err != nil {
		return err
	}

	ge.G.OptionsConfig.Close()
	ge.G.ProgressConfig.Close()

	return nil
}

func main() {
	if err := ge.InitFlags(); err != nil {
		panic(err)
	}

	logfilePath := getLogFilePath()
	if err := os.MkdirAll(filepath.Dir(logfilePath), 0666); err != nil {
		panic(err)
	}
	logfile, err := os.OpenFile(logfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	opts := slogcolor.DefaultOptions
	opts.Level = getLogLevel()
	opts.TimeFormat = time.DateTime
	opts.SrcFileMode = slogcolor.ShortFile
	slog.SetDefault(slog.New(slogcolor.NewHandler(io.MultiWriter(os.Stderr, NewStripANSIWriter(logfile)), opts)))

	// moved main func to run(); learned this from Melkey
	if err := run(); err != nil {
		slog.Error(err.Error())
		if flagutil.MustGetBool(ge.FlagSet, "gui") {
			zenity.Error(err.Error())
		}
		os.Exit(1)
	}
}
