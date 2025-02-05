package gui

import (
	"image/color"
	"io/fs"

	eui_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// UIRes is the main [UIResources] instance.
var UIRes *UIResources

func Init(fsys fs.FS) error {
	if UIRes != nil {
		return nil
	}
	res, err := NewUIResources(fsys)
	if err != nil {
		return err
	}
	UIRes = res
	return nil
}

type UIResources struct {
	Fonts *Fonts

	BGImage                      *eui_image.NineSlice
	TitleBarBGImage              *eui_image.NineSlice
	LabelColor                   *widget.LabelColor
	PageListScrollContainerImage *widget.ScrollContainerImage
	PageListEntryColor           *widget.ListEntryColor
	ButtonImage                  *widget.ButtonImage
	ButtonTextColor              *widget.ButtonTextColor
	DangerButtonImage            *widget.ButtonImage
	DangerButtonTextColor        *widget.ButtonTextColor
	TextAreaScrollContainerImage *widget.ScrollContainerImage
	ScrollSliderTrackImage       *widget.SliderTrackImage
	ScrollButtonImage            *widget.ButtonImage
	TooltipBGImage               *eui_image.NineSlice
	ListSliderTrackImage         *widget.SliderTrackImage
	ListScrollContainerImage     *widget.ScrollContainerImage
	ListEntryColor               *widget.ListEntryColor
	SliderTrackImage             *widget.SliderTrackImage
	SeparatorImage               *eui_image.NineSlice
	CheckboxGraphic              *widget.CheckboxGraphicImage
	CheckboxButtonImage          *widget.ButtonImage
	TextInputImage               *widget.TextInputImage
	TextInputColor               *widget.TextInputColor
}

// NewUIResources creates a new [UIResources].
func NewUIResources(fsys fs.FS) (*UIResources, error) {
	f, err := newFonts(fsys)
	if err != nil {
		return nil, err
	}

	checkImg, _, err := ebitenutil.NewImageFromFileSystem(fsys, "images/ui/checkmark.png")
	if err != nil {
		return nil, err
	}

	return &UIResources{
		Fonts: f,

		BGImage:         eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
		TitleBarBGImage: eui_image.NewNineSliceColor(color.NRGBA{0x0F, 0x0F, 0x0F, 0xFF}),
		LabelColor:      newLabelColorSimple(color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}),
		PageListScrollContainerImage: &widget.ScrollContainerImage{
			Idle:     eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
			Disabled: eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
			Mask:     eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
		},
		PageListEntryColor: &widget.ListEntryColor{
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

		// Button background color = Boyfriend's hair color (#50a5eb).
		// This choice adds a fun, thematic touch to the GUI and serves as a little easter egg
		// for fans of the game, aligning with the playful spirit of Friday Night Funkin'.
		// Also, it just looks good.
		ButtonImage: &widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{R: 0x50, G: 0xA5, B: 0xEB, A: 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{R: 0x29, G: 0x91, B: 0xEC, A: 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{R: 0x18, G: 0x7A, B: 0xCA, A: 0xFF}),
		},
		ButtonTextColor: newButtonTextColorSimple(color.NRGBA{0, 0, 0, 255}),

		DangerButtonImage: &widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x00, 0x00, 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x24, 0x24, 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x40, 0x40, 0xFF}),
		},
		DangerButtonTextColor: newButtonTextColorSimple(color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}),
		TextAreaScrollContainerImage: &widget.ScrollContainerImage{
			Idle:     eui_image.NewNineSliceColor(color.NRGBA{0x0E, 0x0E, 0x0E, 0xFF}),
			Disabled: eui_image.NewNineSliceColor(color.NRGBA{0x0E, 0x0E, 0x0E, 0xFF}),
			Mask:     eui_image.NewNineSliceColor(color.NRGBA{0x0E, 0x0E, 0x0E, 0xFF}),
		},
		ScrollSliderTrackImage: &widget.SliderTrackImage{
			Idle:  eui_image.NewNineSliceColor(color.NRGBA{0, 0, 0, 25}),
			Hover: eui_image.NewNineSliceColor(color.NRGBA{0, 0, 0, 25}),
		},
		ScrollButtonImage: &widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0x4E, 0x4E, 0x4E, 128}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0x5E, 0x5E, 0x5E, 128}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0x6E, 0x6E, 0x6E, 128}),
		},
		TooltipBGImage: eui_image.NewNineSliceColor(color.NRGBA{R: 0x6E, G: 0x6E, B: 0x6E, A: 0xFF}),
		ListSliderTrackImage: &widget.SliderTrackImage{
			Idle:  eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Hover: eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
		},
		ListScrollContainerImage: &widget.ScrollContainerImage{
			Idle:     eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Disabled: eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Mask:     eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
		},
		ListEntryColor: &widget.ListEntryColor{
			Selected:                   color.NRGBA{255, 255, 255, 255},                //Foreground color for the unfocused selected entry
			Unselected:                 color.NRGBA{255, 255, 255, 255},                //Foreground color for the unfocused unselected entry
			SelectedBackground:         color.NRGBA{R: 0x54, G: 0x98, B: 0xD0, A: 255}, //Background color for the unfocused selected entry
			SelectedFocusedBackground:  color.NRGBA{R: 0x58, G: 0x8B, B: 0xB5, A: 255}, //Background color for the focused selected entry
			FocusedBackground:          color.NRGBA{R: 170, G: 170, B: 180, A: 255},    //Background color for the focused unselected entry
			DisabledUnselected:         color.NRGBA{100, 100, 100, 255},                //Foreground color for the disabled unselected entry
			DisabledSelected:           color.NRGBA{100, 100, 100, 255},                //Foreground color for the disabled selected entry
			DisabledSelectedBackground: color.NRGBA{100, 100, 100, 255},                //Background color for the disabled selected entry
		},
		SliderTrackImage: &widget.SliderTrackImage{
			Idle:  eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
			Hover: eui_image.NewNineSliceColor(color.NRGBA{100, 100, 100, 255}),
		},
		SeparatorImage: eui_image.NewNineSliceColor(color.NRGBA{R: 0x3E, G: 0x3E, B: 0x3E, A: 0xFF}),
		CheckboxGraphic: &widget.CheckboxGraphicImage{
			Unchecked: &widget.ButtonImageImage{Idle: ebiten.NewImage(32, 32)},
			Checked:   &widget.ButtonImageImage{Idle: checkImg},
			Greyed:    &widget.ButtonImageImage{Idle: ebiten.NewImage(32, 32)},
		},
		CheckboxButtonImage: &widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{R: 75, G: 75, B: 75, A: 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{R: 50, G: 50, B: 50, A: 0xFF}),
		},
		TextInputImage: &widget.TextInputImage{
			Idle:     eui_image.NewNineSliceColor(color.NRGBA{0x0E, 0x0E, 0x0E, 0xFF}),
			Disabled: eui_image.NewNineSliceColor(color.NRGBA{0x0E, 0x0E, 0x0E, 0xFF}),
		},
		TextInputColor: &widget.TextInputColor{
			Idle:          color.NRGBA{254, 255, 255, 255},
			Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			Caret:         color.NRGBA{254, 255, 255, 255},
			DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
		},
	}, nil
}
