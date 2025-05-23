package optionsui

import (
	"image"
	"image/color"
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/gui"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/pkg/context"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func MakeUI(ctx *context.Context, shouldExit *bool) (*ebitenui.UI, error) {
	ui := &ebitenui.UI{}

	cfg := ctx.OptionsConfig.Data()

	// The main window container
	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(gui.UIRes.BGImage),
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
			widget.GridLayoutOpts.Spacing(10, 0),
		)),
	)
	windowContainer.AddChild(mainContainer)

	// The pages
	pages := []any{
		newGameplayPage(ctx),
		newControlsPage(ctx),
		newAudioPage(ctx, cfg),
		newGraphicsPage(ctx, cfg),
		// TODO: Network tab; coming soon
		newMiscellaneousPage(ctx, cfg, ui),
		newModsPage(ctx),
		newAboutPage(ctx, ui),
	}

	pageContainer := newPageContainer()

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
			widget.ScrollContainerOpts.Image(gui.UIRes.PageListScrollContainerImage),
		),
		widget.ListOpts.HideHorizontalSlider(),
		widget.ListOpts.HideVerticalSlider(),
		widget.ListOpts.EntryFontFace(gui.UIRes.Fonts.PageListEntryFace),
		widget.ListOpts.EntryColor(gui.UIRes.PageListEntryColor),
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
	footerContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
		widget.ButtonOpts.Text(i18n.L("OK"), gui.UIRes.Fonts.FooterButtonFace, gui.UIRes.ButtonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 30,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[optionsui] clicked OK button")

			// save config
			slog.Info("[optionsui] saving config")
			ctx.OptionsConfig.SetData(cfg)

			// exit
			*shouldExit = true
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.ToolTip(widget.NewTextToolTip(
				i18n.L("OptionsRestartWarningTooltip"),
				gui.UIRes.Fonts.RegularFace,
				color.Black,
				gui.UIRes.TooltipBGImage,
			)),
		),
	))

	// Cancel button (exits)
	footerContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
		widget.ButtonOpts.Text(i18n.L("Cancel"), gui.UIRes.Fonts.FooterButtonFace, gui.UIRes.ButtonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 10,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[optionsui] clicked cancel button")
			*shouldExit = true
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.ToolTip(widget.NewTextToolTip(
				i18n.L("OptionsRestartWarningTooltip"),
				gui.UIRes.Fonts.RegularFace,
				color.Black,
				gui.UIRes.TooltipBGImage,
			)),
		),
	))

	// Apply button (saves config)
	footerContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
		widget.ButtonOpts.Text(i18n.L("Apply"), gui.UIRes.Fonts.FooterButtonFace, gui.UIRes.ButtonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 10,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[optionsui] clicked apply button")

			slog.Info("[optionsui] saving config")
			ctx.OptionsConfig.SetData(cfg)
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.ToolTip(widget.NewTextToolTip(
				i18n.L("OptionsRestartWarningTooltip"),
				gui.UIRes.Fonts.RegularFace,
				color.Black,
				gui.UIRes.TooltipBGImage,
			)),
		),
	))

	windowContainer.AddChild(footerContainer)

	// The title bar container
	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(gui.UIRes.TitleBarBGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18n.L("Options"), gui.UIRes.Fonts.TitleFace, gui.UIRes.LabelColor),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				}),
			),
		),
	))

	winwidth := 800
	winheight := 600

	// The window
	window := widget.NewWindow(
		widget.WindowOpts.Contents(windowContainer),
		widget.WindowOpts.TitleBar(titleBarContainer, 25),
		widget.WindowOpts.Draggable(),
		widget.WindowOpts.Modal(),
		widget.WindowOpts.CloseMode(widget.NONE),
		widget.WindowOpts.MinSize(winwidth, winheight),
		widget.WindowOpts.Resizeable(),
	)

	ui.Container = widget.NewContainer()

	// Spawn window
	x, y := window.Contents.PreferredSize()
	window.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(engine.GameWidth/2)-float64(winwidth)/2), int(float64(engine.GameHeight/2)-float64(winheight)/2))))
	ui.AddWindow(window)

	return ui, nil
}
