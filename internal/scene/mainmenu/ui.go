package mainmenu

import (
	"image"
	"image/color"
	_ "image/png"
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui"
	eui_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func MakeUI(ctx *context.Context, shouldExit *bool) (*ebitenui.UI, error) {
	ui := &ebitenui.UI{}

	// Title font
	titleFontFile, err := ctx.AssetsFS.Open("fonts/NotoSans-Bold.ttf")
	if err != nil {
		return nil, err
	}

	titleFaceSource, err := text.NewGoTextFaceSource(titleFontFile)
	if err != nil {
		return nil, err
	}

	titleFace := &text.GoTextFace{
		Source: titleFaceSource,
		Size:   24,
	}

	// Regular font
	regularFontFile, err := ctx.AssetsFS.Open("fonts/NotoSans-Regular.ttf")
	if err != nil {
		return nil, err
	}

	regularFaceSource, err := text.NewGoTextFaceSource(regularFontFile)
	if err != nil {
		return nil, err
	}

	regularFace := &text.GoTextFace{
		Source: regularFaceSource,
		Size:   16,
	}

	var exitDialog *widget.Window

	// exit dialog container
	exitDialogContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eui_image.NewNineSliceColor(color.NRGBA{0x1E, 0x1E, 0x1E, 0xFF})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(20)),
		)),
	)

	// Red "Exit" button
	exitDialogContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
		),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x00, 0x00, 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x24, 0x24, 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0xFF, 0x40, 0x40, 0xFF}),
		}),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "ExitDialogExitButton"), regularFace, &widget.ButtonTextColor{
			color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
			color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
			color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
			color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF},
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   10,
			Right:  20,
			Top:    0,
			Bottom: 0,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("clicked exit button, exiting")
			*shouldExit = true
		}),
	))

	// Green "Stay" button
	exitDialogContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
		),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    eui_image.NewNineSliceColor(color.NRGBA{0x00, 0xFF, 0x00, 0xFF}),
			Hover:   eui_image.NewNineSliceColor(color.NRGBA{0x24, 0xFF, 0x24, 0xFF}),
			Pressed: eui_image.NewNineSliceColor(color.NRGBA{0x40, 0xFF, 0x40, 0xFF}),
		}),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "ExitDialogStayButton"), regularFace, &widget.ButtonTextColor{
			color.NRGBA{0, 0, 0, 0xFF},
			color.NRGBA{0, 0, 0, 0xFF},
			color.NRGBA{0, 0, 0, 0xFF},
			color.NRGBA{0, 0, 0, 0xFF},
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   10,
			Right:  10,
			Top:    0,
			Bottom: 0,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("clicked stay button")
			exitDialog.Close()
		}),
	))

	// The text
	exitDialogContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "ExitDialogText"), regularFace, &widget.LabelColor{color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}, color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}}),
	))

	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(eui_image.NewNineSliceColor(color.NRGBA{0x0F, 0x0F, 0x0F, 0xFF})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "ExitDialogTitle"), titleFace, &widget.LabelColor{color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}, color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}}),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				}),
			),
		),
	))

	exitDialog = widget.NewWindow(
		widget.WindowOpts.Contents(exitDialogContainer),
		widget.WindowOpts.TitleBar(titleBarContainer, 25),
		widget.WindowOpts.Draggable(),
		//widget.WindowOpts.Modal(),
		widget.WindowOpts.CloseMode(widget.NONE),
		widget.WindowOpts.MinSize(400, 200),
	)

	root := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	exitImg, _, err := ebitenutil.NewImageFromFileSystem(ctx.AssetsFS, "images/exit.png")
	if err != nil {
		return nil, err
	}

	root.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
		),

		widget.ButtonOpts.Image(&widget.ButtonImage{
			eui_image.NewNineSliceColor(color.Transparent),
			eui_image.NewNineSliceColor(color.Transparent),
			eui_image.NewNineSliceColor(color.Transparent),
			eui_image.NewNineSliceColor(color.Transparent),
			eui_image.NewNineSliceColor(color.Transparent),
		}),

		widget.ButtonOpts.Graphic(exitImg),

		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("clicked exit button, spawning exit dialog")
			x, y := exitDialog.Contents.PreferredSize()
			exitDialog.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(ctx.Width/2)-200), int(float64(ctx.Height/2)-100))))
			ui.AddWindow(exitDialog)
		}),
	))

	ui.Container = root

	return ui, nil
}
