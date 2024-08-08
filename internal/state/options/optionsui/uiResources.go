package optionsui

import (
	"io/fs"

	"github.com/MatusOllah/gophengine/context"
	"github.com/golang/freetype/truetype"
)

type uiResources struct {
	notoRegular *truetype.Font
	notoBold    *truetype.Font
}

func newUIResources(ctx *context.Context) (*uiResources, error) {
	nr, nb, err := loadFonts(ctx)
	if err != nil {
		return nil, err
	}

	return &uiResources{
		notoRegular: nr,
		notoBold:    nb,
	}, nil
}

// loadFonts loads the Noto Sans Bold and Noto Sans Regular fonts from embedded FS.
func loadFonts(ctx *context.Context) (*truetype.Font, *truetype.Font, error) {
	// Bold font
	boldBytes, err := fs.ReadFile(ctx.AssetsFS, "fonts/NotoSans-Bold.ttf")
	if err != nil {
		return nil, nil, err
	}

	boldFont, err := truetype.Parse(boldBytes)
	if err != nil {
		return nil, nil, err
	}

	// Regular font
	regularBytes, err := fs.ReadFile(ctx.AssetsFS, "fonts/NotoSans-Regular.ttf")
	if err != nil {
		return nil, nil, err
	}

	regularFont, err := truetype.Parse(regularBytes)
	if err != nil {
		return nil, nil, err
	}

	return boldFont, regularFont, nil
}
