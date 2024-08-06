package optionsui

import (
	"image"
	"image/color"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui"
	eui_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func MakeUI(ctx *context.Context) (*ebitenui.UI, error) {
	notoBold, _, err := loadFonts(ctx)
	if err != nil {
		return nil, err
	}

	// Title font
	titleFace := truetype.NewFace(notoBold, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(5)),
		)),
	)
	windowContainer.AddChild(makePageList(notoBold))

	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eui_image.NewNineSliceColor(color.NRGBA{0x0F, 0x0F, 0x0F, 0xFF})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowTitle"), titleFace, newLabelColorSimple(color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF})),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				}),
			),
		),
	))

	window := widget.NewWindow(
		widget.WindowOpts.Contents(windowContainer),
		widget.WindowOpts.TitleBar(titleBarContainer, 25),
		widget.WindowOpts.Draggable(),
		widget.WindowOpts.Modal(),
		widget.WindowOpts.CloseMode(widget.NONE),
		widget.WindowOpts.MinSize(640, 480),
		widget.WindowOpts.Resizeable(),
	)

	ui := &ebitenui.UI{Container: widget.NewContainer()}

	x, y := window.Contents.PreferredSize()
	window.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(ctx.Width/2)-640/2), int(float64(ctx.Height/2)-480/2))))
	ui.AddWindow(window)

	return ui, nil
}

func makePageList(theFont *truetype.Font) widget.PreferredSizeLocateableWidget {
	face := truetype.NewFace(theFont, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	return widget.NewList(
		widget.ListOpts.Entries([]any{"Audio", "Video", "Gameplay", "Controls", "Network", "Mods", "About"}),
		widget.ListOpts.EntryLabelFunc(func(a any) string {
			return a.(string)
		}),
		widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(150, 0),
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				StretchVertical:    true,
				Padding:            widget.NewInsetsSimple(10),
			}),
		)),
		widget.ListOpts.SliderOpts(
			widget.SliderOpts.Images(nil, nil),
			widget.SliderOpts.MinHandleSize(0),
			widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(0)),
		),
		widget.ListOpts.ScrollContainerOpts(
			widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
				Idle:     eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
				Disabled: eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
				Mask:     eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF}),
			}),
		),
		widget.ListOpts.HideHorizontalSlider(),
		widget.ListOpts.HideVerticalSlider(),
		widget.ListOpts.EntryFontFace(face),
		widget.ListOpts.EntryColor(&widget.ListEntryColor{
			Selected:                   color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Foreground color for the unfocused selected entry
			Unselected:                 color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Foreground color for the unfocused unselected entry
			SelectedBackground:         color.NRGBA{R: 0x2E, G: 0x2E, B: 0x2E, A: 0xFF}, // Background color for the unfocused selected entry
			SelectingBackground:        color.NRGBA{R: 0x3E, G: 0x3E, B: 0x3E, A: 0xFF}, // Background color for the unfocused being selected entry
			SelectingFocusedBackground: color.NRGBA{R: 0x4E, G: 0x4E, B: 0x4E, A: 0xFF}, // Background color for the focused being selected entry
			SelectedFocusedBackground:  color.NRGBA{R: 0x3E, G: 0x3E, B: 0x3E, A: 0xFF}, // Background color for the focused selected entry
			FocusedBackground:          color.NRGBA{R: 0x2E, G: 0x2E, B: 0x2E, A: 0xFF}, // Background color for the focused unselected entry
			DisabledUnselected:         color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Foreground color for the disabled unselected entry
			DisabledSelected:           color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, // Foreground color for the disabled selected entry
			DisabledSelectedBackground: color.NRGBA{R: 0x2E, G: 0x2E, B: 0x2E, A: 0xFF}, // Background color for the disabled selected entry
		}),
		widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
		widget.ListOpts.EntryTextPosition(widget.TextPositionStart, widget.TextPositionCenter),
	)
}
