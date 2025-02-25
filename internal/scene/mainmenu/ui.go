package mainmenu

import (
	"image"
	"image/color"
	_ "image/png"
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/gui"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/ebitenui/ebitenui"
	eui_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func MakeUI(ctx *context.Context, shouldExit *bool) (*ebitenui.UI, error) {
	ui := &ebitenui.UI{}

	var exitDialog *widget.Window

	// exit dialog container
	exitDialogContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(gui.UIRes.BGImage),
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
		widget.ButtonOpts.Image(gui.UIRes.DangerButtonImage),
		widget.ButtonOpts.Text(i18n.L("Exit"), gui.UIRes.Fonts.RegularFace, gui.UIRes.DangerButtonTextColor),
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
		widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
		widget.ButtonOpts.Text(i18n.L("Stay"), gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
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
		widget.LabelOpts.Text(i18n.L("ExitDialogText"), gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor),
	))

	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(gui.UIRes.TitleBarBGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18n.L("Exit"), gui.UIRes.Fonts.TitleFace, gui.UIRes.LabelColor),
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

	exitImg, _, err := ebitenutil.NewImageFromFileSystem(ctx.AssetsFS, "images/ui/exit.png")
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
			Idle:         eui_image.NewNineSliceColor(color.Transparent),
			Hover:        eui_image.NewNineSliceColor(color.Transparent),
			Pressed:      eui_image.NewNineSliceColor(color.Transparent),
			PressedHover: eui_image.NewNineSliceColor(color.Transparent),
			Disabled:     eui_image.NewNineSliceColor(color.Transparent),
		}),

		widget.ButtonOpts.Graphic(exitImg),

		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("clicked exit button, spawning exit dialog")
			x, y := exitDialog.Contents.PreferredSize()
			exitDialog.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(engine.GameWidth/2)-200), int(float64(engine.GameHeight/2)-100))))
			ui.AddWindow(exitDialog)
		}),
	))

	ui.Container = root

	return ui, nil
}
