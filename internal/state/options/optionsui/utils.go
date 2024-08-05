package optionsui

import (
	"image/color"
	"io/fs"

	"github.com/MatusOllah/gophengine/context"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
)

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

// loadPhantomFont loads the Phantom font from embedded FS.
func loadPhantomFont(ctx *context.Context) (*truetype.Font, error) {
	phantomBytes, err := fs.ReadFile(ctx.AssetsFS, "fonts/phantom-full.ttf")
	if err != nil {
		return nil, err
	}

	phantomFont, err := truetype.Parse(phantomBytes)
	if err != nil {
		return nil, err
	}

	return phantomFont, nil
}

// newLabelColorSimple is short for &widget.LabelColor{clr, clr}.
func newLabelColorSimple(clr color.Color) *widget.LabelColor {
	return &widget.LabelColor{clr, clr}
}
