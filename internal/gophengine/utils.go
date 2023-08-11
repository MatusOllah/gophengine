package gophengine

import (
	"fmt"
	"image"
	_ "image/png"

	"github.com/MatusOllah/gophengine/assets"
	"github.com/qeesung/image2ascii/convert"
	"github.com/ztrue/tracerr"
)

func StringSliceContains(s []string, a string) bool {
	for _, value := range s {
		if value == a {
			return true
		}
	}

	return false
}

func PrintBoyfriend() error {
	opts := convert.DefaultOptions
	opts.FixedWidth = 50
	opts.FixedHeight = 25

	f, err := assets.FS.Open("images/characters/bf/BF HEY!!0025.png")
	if err != nil {
		return tracerr.Wrap(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return tracerr.Wrap(err)
	}

	fmt.Print(convert.NewImageConverter().Image2ASCIIString(img, &opts))

	return nil
}
