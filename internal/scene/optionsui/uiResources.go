package optionsui

import (
	"image/color"

	"github.com/MatusOllah/gophengine/context"
	eui_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type uiResources struct {
	fonts *fonts

	bgImage                      *eui_image.NineSlice
	titleBarBGImage              *eui_image.NineSlice
	labelColor                   *widget.LabelColor
	pageListScrollContainerImage *widget.ScrollContainerImage
	pageListEntryColor           *widget.ListEntryColor
	buttonImage                  *widget.ButtonImage
	buttonTextColor              *widget.ButtonTextColor
	exitButtonImage              *widget.ButtonImage
	exitButtonTextColor          *widget.ButtonTextColor
	textAreaScrollContainerImage *widget.ScrollContainerImage
	scrollSliderTrackImage       *widget.SliderTrackImage
	scrollButtonImage            *widget.ButtonImage
	tooltipBGImage               *eui_image.NineSlice
	listSliderTrackImage         *widget.SliderTrackImage
	listScrollContainerImage     *widget.ScrollContainerImage
	listEntryColor               *widget.ListEntryColor
}

func newUIResources(ctx *context.Context) (*uiResources, error) {
	f, err := newFonts(ctx)
	if err != nil {
		return nil, err
	}

	return &uiResources{
		fonts: f,

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

		// Button background color inspired by Boyfriend's hair color (#50a5eb).
		// This choice adds a fun, thematic touch to the GUI and serves as a little easter egg
		// for fans of the game, aligning with the playful spirit of Friday Night Funkin'.
		// Also, it just looks good.
		buttonImage: &widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{R: 0x50, G: 0xA5, B: 0xEB, A: 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{R: 0x29, G: 0x91, B: 0xEC, A: 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{R: 0x18, G: 0x7A, B: 0xCA, A: 0xFF}),
		},
		buttonTextColor: newButtonTextColorSimple(color.NRGBA{0, 0, 0, 255}),

		exitButtonImage: &widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x00, 0x00, 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x24, 0x24, 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x40, 0x40, 0xFF}),
		},
		exitButtonTextColor: newButtonTextColorSimple(color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}),
		textAreaScrollContainerImage: &widget.ScrollContainerImage{
			Idle:     eui_image.NewNineSliceColor(color.NRGBA{0x0E, 0x0E, 0x0E, 0xFF}),
			Disabled: eui_image.NewNineSliceColor(color.NRGBA{0x0E, 0x0E, 0x0E, 0xFF}),
			Mask:     eui_image.NewNineSliceColor(color.NRGBA{0x0E, 0x0E, 0x0E, 0xFF}),
		},
		scrollSliderTrackImage: &widget.SliderTrackImage{
			Idle:  eui_image.NewNineSliceColor(color.NRGBA{0, 0, 0, 25}),
			Hover: eui_image.NewNineSliceColor(color.NRGBA{0, 0, 0, 25}),
		},
		scrollButtonImage: &widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0x4E, 0x4E, 0x4E, 128}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0x5E, 0x5E, 0x5E, 128}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0x6E, 0x6E, 0x6E, 128}),
		},
		tooltipBGImage: eui_image.NewNineSliceColor(color.NRGBA{R: 0x6E, G: 0x6E, B: 0x6E, A: 0xFF}),
		listSliderTrackImage: &widget.SliderTrackImage{
			Idle:  eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Hover: eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
		},
		listScrollContainerImage: &widget.ScrollContainerImage{
			Idle:     eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Disabled: eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Mask:     eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
		},
		listEntryColor: &widget.ListEntryColor{
			Selected:                   color.NRGBA{255, 255, 255, 255},                //Foreground color for the unfocused selected entry
			Unselected:                 color.NRGBA{255, 255, 255, 255},                //Foreground color for the unfocused unselected entry
			SelectedBackground:         color.NRGBA{R: 0x54, G: 0x98, B: 0xD0, A: 255}, //Background color for the unfocused selected entry
			SelectedFocusedBackground:  color.NRGBA{R: 0x58, G: 0x8B, B: 0xB5, A: 255}, //Background color for the focused selected entry
			FocusedBackground:          color.NRGBA{R: 170, G: 170, B: 180, A: 255},    //Background color for the focused unselected entry
			DisabledUnselected:         color.NRGBA{100, 100, 100, 255},                //Foreground color for the disabled unselected entry
			DisabledSelected:           color.NRGBA{100, 100, 100, 255},                //Foreground color for the disabled selected entry
			DisabledSelectedBackground: color.NRGBA{100, 100, 100, 255},                //Background color for the disabled selected entry
		},
	}, nil
}
