package gophengine

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/rs/zerolog/log"
	"github.com/ztrue/tracerr"
)

func ExtractAssets() error {
	if err := os.Mkdir("assets", fs.ModePerm); err != nil {
		return tracerr.Wrap(err)
	}

	err := fs.WalkDir(assets.FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return tracerr.Wrap(err)
		}

		if d.IsDir() {
			dirPath := filepath.Join("assets", path)

			log.Info().Msgf("creating directory %s", dirPath)
			if err := os.MkdirAll(dirPath, fs.ModePerm); err != nil {
				return tracerr.Wrap(err)
			}

			return nil
		}

		dstPath := filepath.Join("assets", path)

		log.Info().Msgf("extracting %s => %s", path, dstPath)

		src, err := assets.FS.Open(path)
		if err != nil {
			return tracerr.Wrap(err)
		}
		defer src.Close()

		dst, err := os.Create(dstPath)
		if err != nil {
			return tracerr.Wrap(err)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return tracerr.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
