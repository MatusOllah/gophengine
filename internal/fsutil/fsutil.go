package fsutil

import (
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

// NumFiles counts all the files in the filesystem and returns the number of files.
func NumFiles(fsys fs.FS) (int, error) {
	var num int
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		num++
		return nil
	})
	if err != nil {
		return 0, err
	}

	return num, nil
}

// Extract extracts the filesystem to dst.
func Extract(fsys fs.FS, dst string) error {

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
