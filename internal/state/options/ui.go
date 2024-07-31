package options

import (
	"image"
	"image/color"
	"io/fs"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui"
	eui_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func makeUI(ctx *context.Context) (*ebitenui.UI, error) {
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

	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eui_image.NewNineSliceColor(color.NRGBA{0x0F, 0x0F, 0x0F, 0xFF})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowTitle"), titleFace, newLabelColorSimple(color.NRGBA{255, 255, 255, 255})),
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
	)

	ui := &ebitenui.UI{Container: widget.NewContainer()}

	x, y := window.Contents.PreferredSize()
	window.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(ctx.Width/2)-640/2), int(float64(ctx.Height/2)-480/2))))
	ui.AddWindow(window)

	return ui, nil
}

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

func newLabelColorSimple(clr color.Color) *widget.LabelColor {
	return &widget.LabelColor{clr, clr}
}
