package optionsui

import (
	"image/color"
	"io/fs"

	"github.com/MatusOllah/gophengine/context"
	eui_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
)

type uiResources struct {
	notoRegular *truetype.Font
	notoBold    *truetype.Font

	bgImage                      *eui_image.NineSlice
	titleBarBGImage              *eui_image.NineSlice
	labelColor                   *widget.LabelColor
	pageListScrollContainerImage *widget.ScrollContainerImage
	pageListEntryColor           *widget.ListEntryColor
	buttonImage                  *widget.ButtonImage
	buttonTextColor              *widget.ButtonTextColor
}

func newUIResources(ctx *context.Context) (*uiResources, error) {
	nb, nr, err := loadFonts(ctx)
	if err != nil {
		return nil, err
	}

	return &uiResources{
		notoRegular: nr,
		notoBold:    nb,

		bgImage:         eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
		titleBarBGImage: eui_image.NewNineSliceColor(color.NRGBA{0x0F, 0x0F, 0x0F, 0xFF}),
		labelColor:      newLabelColorSimple(color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}),
		pageListScrollContainerImage: &widget.ScrollContainerImage{
			Idle:     eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
			Disabled: eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
			Mask:     eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
		},
		pageListEntryColor: &widget.ListEntryColor{
			Selected:                   color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Foreground color for the unfocused selected entry
			Unselected:                 color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Foreground color for the unfocused unselected entry
			SelectedBackground:         color.NRGBA{R: 0x3E, G: 0x3E, B: 0x3E, A: 0xFF}, // Background color for the unfocused selected entry
			SelectingBackground:        color.NRGBA{R: 0x1E, G: 0x1E, B: 0x1E, A: 0xFF}, // Background color for the unfocused being selected entry
			SelectingFocusedBackground: color.NRGBA{R: 0x1E, G: 0x1E, B: 0x1E, A: 0xFF}, // Background color for the focused being selected entry
			SelectedFocusedBackground:  color.NRGBA{R: 0x3E, G: 0x3E, B: 0x3E, A: 0xFF}, // Background color for the focused selected entry
			FocusedBackground:          color.NRGBA{R: 0x2E, G: 0x2E, B: 0x2E, A: 0xFF}, // Background color for the focused unselected entry
			DisabledUnselected:         color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Foreground color for the disabled unselected entry
			DisabledSelected:           color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Foreground color for the disabled selected entry
			DisabledSelectedBackground: color.NRGBA{R: 0x2E, G: 0x2E, B: 0x2E, A: 0xFF}, // Background color for the disabled selected entry
		},
		buttonImage: &widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0xBF, 0xBF, 0xBF, 0xFF}), // "0xBF" :D
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0xCF, 0xCF, 0xCF, 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0xDF, 0xDF, 0xDF, 0xFF}),
		},
		buttonTextColor: newButtonTextColorSimple(color.NRGBA{0, 0, 0, 255}),
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