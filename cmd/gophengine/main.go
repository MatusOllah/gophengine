// Copyright 2025 Matúš Ollah
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
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

	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/fnfgame"
	"github.com/MatusOllah/gophengine/internal/dialog"
	"github.com/MatusOllah/slogcolor"
	"github.com/MatusOllah/stripansi"
	"github.com/hajimehoshi/ebiten/v2"
	"gopkg.in/natefinch/lumberjack.v2"
)

// extractFS extracts the filesystem to dst.
func extractFS(fsys fs.FS, dst string) error {
	// create destination directory
	if err := os.Mkdir(dst, fs.ModePerm); err != nil {
		return err
	}

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// create directory
		if d.IsDir() {
			dirPath := filepath.Join(dst, path)

			slog.Info("creating directory", "path", dirPath)
			if err := os.MkdirAll(dirPath, fs.ModePerm); err != nil {
				return err
			}

			return nil
		}

		// create file
		dstPath := filepath.Join(dst, path)

		slog.Info("extracting", "src", path, "dst", dstPath)

		srcFile, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// setIcon sets the window icon.
func setIcon() error {
	if runtime.GOARCH == "wasm" {
		return nil
	}

	f, err := assets.FS.Open("icon.png")
	if err != nil {
		return fmt.Errorf("failed to open icon image file: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to decode icon image file: %w", err)
	}

	ebiten.SetWindowIcon([]image.Image{img})

	return nil
}

// getLogLevel gets the log level from command-line flags.
func getLogLevel() slog.Leveler {
	switch s := strings.ToLower(*logLevelFlag); s {
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

// Actual main func here
func mainE() error {
	slog.Info(fmt.Sprintf("GophEngine version %s", version))
	slog.Info(fmt.Sprintf("Go version %s", runtime.Version()))
	slog.Info(fmt.Sprintf("Friday Night Funkin' version %s", fnfVersion))
	slog.Info("ahoj!")

	if *extractAssetsFlag {
		if err := extractFS(assets.FS, "assets"); err != nil {
			return fmt.Errorf("failed to extract assets: %w", err)
		}

		return nil
	}

	// Window icon
	slog.Info("setting window icon")
	if err := setIcon(); err != nil {
		return fmt.Errorf("failed to set window icon: %w", err)
	}

	// Context
	slog.Info("initializing context")

	cfg := &context.NewContextConfig{
		AssetsFS:           assets.FS,
		OptionsConfigPath:  *configFlag,
		ProgressConfigPath: *progressFlag,
		Version:            version,
		FNFVersion:         fnfVersion,
		Locale:             *forceLocaleFlag,
	}
	if runtime.GOARCH != "wasm" {
		if *configFlag == "" {
			configDir, err := os.UserConfigDir()
			if err != nil {
				return fmt.Errorf("failed to fetch user config dir: %w", err)
			}

			cfg.OptionsConfigPath = filepath.Join(configDir, "GophEngine/options.gecfg")
		}
		if *progressFlag == "" {
			configDir, err := os.UserConfigDir()
			if err != nil {
				return fmt.Errorf("failed to fetch user config dir: %w", err)
			}

			cfg.ProgressConfigPath = filepath.Join(configDir, "GophEngine/progress.gecfg")
		}
		if *portableFlag {
			cfg.OptionsConfigPath = filepath.Join("GophEngine/options.gecfg")
			cfg.ProgressConfigPath = filepath.Join("GophEngine/progress.gecfg")
		}
	}
	ctx, err := context.New(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize context: %w", err)
	}

	// Game init
	slog.Info("initializing game")
	g, err := fnfgame.New(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize game: %w", err)
	}

	// Ebiten init
	slog.Info("initializing ebitengine")
	g.InitEbiten()

	if *justInitFlag {
		return nil
	}

	// Start
	slog.Info("starting game")
	if err := g.Start(); err != nil {
		return fmt.Errorf("failed to run game: %w", err)
	}

	return g.Close()
}

func main() {
	// Flags
	if err := initFlags(); err != nil {
		panic(err)
	}

	// Logger (using slogcolor: https://github.com/MatusOllah/slogcolor)
	opts := slogcolor.DefaultOptions
	opts.Level = getLogLevel()
	opts.SrcFileLength = 32

	if runtime.GOARCH != "wasm" {
		tempDir := os.TempDir()
		if *portableFlag {
			tempDir = "."
		}

		slog.SetDefault(slog.New(slogcolor.NewHandler(io.MultiWriter(os.Stderr, stripansi.NewWriter(&lumberjack.Logger{
			Filename:   filepath.Join(tempDir, "GophEngine", "logs", "game.log"),
			MaxSize:    500, // megabytes
			MaxBackups: 5,
			MaxAge:     28, // days
			Compress:   true,
		})), opts)))
	} else {
		slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, opts)))
	}

	// moved main func to mainE and handle error here
	// learned this from Melkey
	if err := mainE(); err != nil {
		slog.Error(err.Error())
		dialog.Error(err.Error())
		os.Exit(1)
	}
}
