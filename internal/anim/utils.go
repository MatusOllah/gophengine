package anim

import (
	"cmp"
	"errors"
	"fmt"
	_ "image/png"
	"io/fs"
	"path/filepath"
	"slices"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var ErrNotFound error = errors.New("image not found")

// GetImagesByPrefixFromFS is basically addByPrefix in HaxeFlixel.
// It gets images by prefix from the filesystem and returns them.
func GetImagesByPrefixFromFS(fsys fs.FS, path string, prefix string) ([]*ebiten.Image, error) {
	var finalFiles []string
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			name := d.Name()

			if name[:len(name)-8] == prefix {
				finalFiles = append(finalFiles, name)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(finalFiles, func(i, j int) bool {
		name1 := finalFiles[i]
		name2 := finalFiles[j]

		index1, _ := strconv.Atoi(name1[len(name1)-4:])
		index2, _ := strconv.Atoi(name2[len(name2)-4:])

		return index1 < index2
	})

	var images []*ebiten.Image
	for _, f := range finalFiles {
		img, _, err := ebitenutil.NewImageFromFileSystem(fsys, path+"/"+f)
		if err != nil {
			return nil, err
		}

		images = append(images, img)
	}

	return images, nil
}

// GetImagesByIndicesFromFS is basically addByIndices in HaxeFlixel.
// It gets images by indices from the filesystem and returns them.
func GetImagesByIndicesFromFS(fsys fs.FS, path string, prefix string, indices []int) ([]*ebiten.Image, error) {
	var files []string
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			name := d.Name()

			if name[:len(name)-8] == prefix {
				files = append(files, name)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	var fileIndices []int

	for _, index := range indices {
		wantedFile := prefix + fmt.Sprintf("%04d", index) + ".png"
		fileIndex, ok := slices.BinarySearchFunc(files, wantedFile, func(s1, s2 string) int {
			return cmp.Compare(filepath.Base(s1), filepath.Base(s2))
		})
		if !ok {
			return nil, ErrNotFound
		}
		fileIndices = append(fileIndices, fileIndex)
	}

	var images []*ebiten.Image
	for _, i := range fileIndices {
		img, _, err := ebitenutil.NewImageFromFileSystem(fsys, path+"/"+files[i])
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	return images, nil
}

func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
