package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/fnfgame"
	"github.com/MatusOllah/gophengine/internal/flagutil"
	"github.com/MatusOllah/gophengine/internal/fsutil"
	"github.com/MatusOllah/slogcolor"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ncruces/zenity"
)

// setIcon sets the window icon.
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

// getLogLevel gets the log level from command-line flags.
func getLogLevel() slog.Leveler {
	switch s := strings.ToLower(flagutil.MustGetString(flagSet, "log-level")); s {
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

// getLogFilePath get the logfile path from the temporary directory and current time.
func getLogFilePath() string {
	return filepath.Join(os.TempDir(), "GophEngine", "logs", time.Now().Format("2006-01-02_15-04-05.log"))
}

func main() {
	// Flags
	if err := initFlags(); err != nil {
		panic(err)
	}

	// Logfile
	logfilePath := getLogFilePath()
	if err := os.MkdirAll(filepath.Dir(logfilePath), 0666); err != nil {
		panic(err)
	}
	logfile, err := os.OpenFile(logfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	// Logger (using slogcolor: https://github.com/MatusOllah/slogcolor)
	opts := slogcolor.DefaultOptions
	opts.Level = getLogLevel()
	opts.TimeFormat = time.DateTime
	opts.SrcFileMode = slogcolor.ShortFile
	slog.SetDefault(slog.New(slogcolor.NewHandler(io.MultiWriter(os.Stderr, NewStripANSIWriter(logfile)), opts)))

	slog.Info(fmt.Sprintf("GophEngine version %s", version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	slog.Info(fmt.Sprintf("Friday Night Funkin' version %s", fnfVersion))
	slog.Info("ahoj!")

	if flagutil.MustGetBool(flagSet, "extract-assets") {
		if err := fsutil.Extract(assets.FS, "assets", flagutil.MustGetBool(flagSet, "gui")); err != nil {
			slog.Error(err.Error())
			if flagutil.MustGetBool(flagSet, "gui") {
				zenity.Error(err.Error())
			}
			os.Exit(1)
		}

		return
	}

	// Window icon
	slog.Info("setting window icon")
	if err := setIcon(); err != nil {
		slog.Error(err.Error())
		if flagutil.MustGetBool(flagSet, "gui") {
			zenity.Error(err.Error())
		}
		os.Exit(1)
	}

	// Context
	slog.Info("initializing context")
	ctx, err := context.New(&context.NewContextConfig{
		AssetsFS:           assets.FS,
		OptionsConfigPath:  flagutil.MustGetString(flagSet, "config"),
		ProgressConfigPath: flagutil.MustGetString(flagSet, "progress"),
	})
	if err != nil {
		slog.Error(err.Error())
		if flagutil.MustGetBool(flagSet, "gui") {
			zenity.Error(err.Error())
		}
		os.Exit(1)
	}

	// Game init
	slog.Info("initializing game")
	g, err := fnfgame.New(ctx) // TODO: portable mode
	if err != nil {
		slog.Error(err.Error())
		if flagutil.MustGetBool(flagSet, "gui") {
			zenity.Error(err.Error())
		}
		os.Exit(1)
	}
	defer func() {
		if err := g.Close(); err != nil {
			slog.Error(err.Error())
			if flagutil.MustGetBool(flagSet, "gui") {
				zenity.Error(err.Error())
			}
			os.Exit(1)
		}
	}()

	// Ebiten init
	slog.Info("initializing ebitengine")
	g.InitEbiten()

	if flagutil.MustGetBool(flagSet, "just-init") {
		return
	}

	// Start
	slog.Info("starting game")
	if err := g.Start(); err != nil {
		slog.Error(err.Error())
		if flagutil.MustGetBool(flagSet, "gui") {
			zenity.Error(err.Error())
		}
		os.Exit(1)
	}
}
