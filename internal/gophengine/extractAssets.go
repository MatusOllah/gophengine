package gophengine

import (
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/MatusOllah/gophengine/assets"
)

func ExtractAssets() error {
	if err := os.Mkdir("assets", fs.ModePerm); err != nil {
		return err
	}

	err := fs.WalkDir(assets.FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			dirPath := filepath.Join("assets", path)

			slog.Info(fmt.Sprintf("creating directory %s", dirPath))
			if err := os.MkdirAll(dirPath, fs.ModePerm); err != nil {
				return err
			}

			return nil
		}

		dstPath := filepath.Join("assets", path)

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

	return nil
}
