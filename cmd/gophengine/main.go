package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
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
	"github.com/MatusOllah/gophengine/internal/fsutil"
	"github.com/MatusOllah/slogcolor"
	"github.com/MatusOllah/stripansi"
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
	switch s := strings.ToLower(logLevelFlag); s {
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
	tempDir := os.TempDir()
	if portableFlag {
		tempDir = "."
	}

	return filepath.Join(tempDir, "GophEngine", "logs", time.Now().Format("2006-01-02_15-04-05.log"))
}

// Actual main func here
func mainE() error {
	slog.Info(fmt.Sprintf("GophEngine version %s", version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	slog.Info(fmt.Sprintf("Friday Night Funkin' version %s", fnfVersion))
	slog.Info("ahoj!")

	if extractAssetsFlag {
		if err := fsutil.Extract(assets.FS, "assets"); err != nil {
			return err
		}

		return nil
	}

	// Window icon
	slog.Info("setting window icon")
	if err := setIcon(); err != nil {
		return err
	}

	// Context
	slog.Info("initializing context")

	cfg := &context.NewContextConfig{
		AssetsFS:           assets.FS,
		OptionsConfigPath:  configFlag,
		ProgressConfigPath: progressFlag,
		Version:            version,
		FNFVersion:         fnfVersion,
		Locale:             forceLocaleFlag,
	}
	if portableFlag {
		cfg.OptionsConfigPath = filepath.Join("GophEngine/options.gecfg")
		cfg.ProgressConfigPath = filepath.Join("GophEngine/progress.gecfg")
	}
	ctx, err := context.New(cfg)
	if err != nil {
		return err
	}

	// Game init
	slog.Info("initializing game")
	g, err := fnfgame.New(ctx)
	if err != nil {
		return err
	}

	// Ebiten init
	slog.Info("initializing ebitengine")
	g.InitEbiten()

	if justInitFlag {
		return nil
	}

	// Start
	slog.Info("starting game")
	if err := g.Start(); err != nil {
		return err
	}

	return g.Close()
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
	logfile, err := os.Create(logfilePath)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	// Logger (using slogcolor: https://github.com/MatusOllah/slogcolor)
	opts := slogcolor.DefaultOptions
	opts.Level = getLogLevel()
	opts.SrcFileLength = 32
	slog.SetDefault(slog.New(slogcolor.NewHandler(io.MultiWriter(os.Stderr, stripansi.NewWriter(logfile)), opts)))

	// moved main func to _main and handle error here
	// learned this from Melkey
	if err := mainE(); err != nil {
		slog.Error(err.Error())
		zenity.Error(err.Error())
		panic(err)
	}
}
