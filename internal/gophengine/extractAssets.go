package gophengine

import (
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/MatusOllah/gophengine/internal/flagutil"
	"github.com/MatusOllah/gophengine/internal/fsutil"
	"github.com/ncruces/zenity"
)

func ExtractAssets() error {
	isGUI := flagutil.MustGetBool(G.FlagSet, "gui")

	var value int = 0
	var dlg zenity.ProgressDialog
	if isGUI {
		slog.Info("counting files")
		numFiles, err := fsutil.NumFiles(assets.FS)
		if err != nil {
			return err
		}
		slog.Info("done", "numFiles", numFiles)

		_dlg, err := zenity.Progress(zenity.Title("Extracting assets"), zenity.MaxValue(numFiles))
		if err != nil {
			return err
		}
		defer _dlg.Close()
		dlg = _dlg

		dlg.Text("Extracting...")
	}

	if err := os.Mkdir("assets", fs.ModePerm); err != nil {
		return err
	}

	err := fs.WalkDir(assets.FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			dirPath := filepath.Join("assets", path)

			if isGUI {
				dlg.Text("Creating directory " + dirPath)
			}

			slog.Info(fmt.Sprintf("creating directory %s", dirPath))
			if err := os.MkdirAll(dirPath, fs.ModePerm); err != nil {
				return err
			}

			return nil
		}

		dstPath := filepath.Join("assets", path)

		if isGUI {
			value++
			dlg.Value(value)
			dlg.Text(fmt.Sprintf("extracting %s", path))
		}

		slog.Info(fmt.Sprintf("extracting %s => %s", path, dstPath))

		src, err := assets.FS.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	if isGUI {
		dlg.Complete()
	}

	return nil
}
