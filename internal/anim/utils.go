package anim

import (
	"io/fs"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GetImagesByPrefixFromFS(prefix string, fsys fs.FS, path string) ([]*ebiten.Image, error) {
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
