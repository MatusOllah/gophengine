package optionsui

import (
	"image"
	"image/color"
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui"
	eui_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func MakeUI(ctx *context.Context) (*ebitenui.UI, error) {
	res, err := newUIResources(ctx)
	if err != nil {
		return nil, err
	}

	// Title font
	titleFace := truetype.NewFace(res.notoBold, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	// The footer button font
	footerButtonFace := truetype.NewFace(res.notoRegular, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	// The main window container
	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true, false}),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(10)),
			widget.GridLayoutOpts.Spacing(0, 10),
		)),
	)

	// The main center container
	mainContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Stretch([]bool{false, true}, []bool{true}),
			widget.GridLayoutOpts.Spacing(20, 0),
		)),
	)
	windowContainer.AddChild(mainContainer)

	// The pages
	pages := []any{
		newTestPage(),
		newTestPage(),
		newTestPage(),
		newTestPage(),
	}

	pageContainer := newPageContainer(res)

	// The page select list
	mainContainer.AddChild(widget.NewList(
		widget.ListOpts.Entries(pages),
		widget.ListOpts.EntryLabelFunc(func(p any) string {
			return p.(*page).name
		}),
		widget.ListOpts.EntrySelectedHandler(func(args *widget.ListEntrySelectedEventArgs) {
			slog.Info("[optionsui] selected page", "PreviousEntry", args.PreviousEntry, "Entry", args.Entry)
			pageContainer.setPage(args.Entry.(*page))
		}),
		widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(150, 0),
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
		widget.ListOpts.EntryFontFace(titleFace),
		widget.ListOpts.EntryColor(&widget.ListEntryColor{
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
		}),
		widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
		widget.ListOpts.EntryTextPosition(widget.TextPositionStart, widget.TextPositionCenter),
	))
	mainContainer.AddChild(pageContainer.widget)

	// The footer container
	footerContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(5),
		)),
	)

	// OK button (saves config & exits)
	// TODO: save config & exit
	footerContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0xBF, 0xBF, 0xBF, 0xFF}), // "0xBF" :D
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0xC4, 0xC4, 0xC4, 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0xC8, 0xC8, 0xC8, 0xFF}),
		}),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowOKButton"), footerButtonFace, newButtonTextColorSimple(color.NRGBA{0, 0, 0, 255})),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 30,
		}),
	))

	// Cancel button (exits)
	footerContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0xBF, 0xBF, 0xBF, 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0xC4, 0xC4, 0xC4, 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0xC8, 0xC8, 0xC8, 0xFF}),
		}),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowCancelButton"), footerButtonFace, newButtonTextColorSimple(color.NRGBA{0, 0, 0, 255})),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 10,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[optionsui] clicked cancel button")
			// TODO: switch state to MainMenuState
		}),
	))

	// Apply button (saves config)
	// TODO: save config
	footerContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0xBF, 0xBF, 0xBF, 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0xC4, 0xC4, 0xC4, 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0xC8, 0xC8, 0xC8, 0xFF}),
		}),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowApplyButton"), footerButtonFace, newButtonTextColorSimple(color.NRGBA{0, 0, 0, 255})),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 10,
		}),
	))

	windowContainer.AddChild(footerContainer)

	// The title bar container
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

	// The window
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

	// Spawn window
	x, y := window.Contents.PreferredSize()
	window.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(ctx.Width/2)-640/2), int(float64(ctx.Height/2)-480/2))))
	ui.AddWindow(window)

	return ui, nil
}
