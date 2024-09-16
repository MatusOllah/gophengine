package optionsui

import (
	"image"
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func MakeUI(ctx *context.Context, shouldExit *bool) (*ebitenui.UI, error) {
	ui := &ebitenui.UI{}

	res, err := newUIResources(ctx)
	if err != nil {
		return nil, err
	}

	// The main window container
	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.bgImage),
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
		newGameplayPage(ctx),
		newControlsPage(ctx),
		newAudioPage(ctx),
		newGraphicsPage(ctx),
		// TODO: Network tab; coming soon
		newMiscellaneousPage(ctx, res),
		newModsPage(ctx),
		newAdvancedPage(ctx),
		newAboutPage(ctx, res, ui),
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
			widget.ScrollContainerOpts.Image(res.pageListScrollContainerImage),
		),
		widget.ListOpts.HideHorizontalSlider(),
		widget.ListOpts.HideVerticalSlider(),
		widget.ListOpts.EntryFontFace(res.fonts.pageListEntryFace),
		widget.ListOpts.EntryColor(res.pageListEntryColor),
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
		widget.ButtonOpts.Image(res.buttonImage),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowOKButton"), res.fonts.footerButtonFace, res.buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 30,
		}),
	))

	// Cancel button (exits)
	footerContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.buttonImage),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowCancelButton"), res.fonts.footerButtonFace, res.buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 10,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[optionsui] clicked cancel button")
			*shouldExit = true
		}),
	))

	// Apply button (saves config)
	// TODO: save config
	footerContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.buttonImage),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowApplyButton"), res.fonts.footerButtonFace, res.buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 10,
		}),
	))

	windowContainer.AddChild(footerContainer)

	// The title bar container
	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.titleBarBGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsWindowTitle"), res.fonts.titleFace, res.labelColor),
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

	ui.Container = widget.NewContainer()

	// Spawn window
	x, y := window.Contents.PreferredSize()
	window.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(ctx.Width/2)-640/2), int(float64(ctx.Height/2)-480/2))))
	ui.AddWindow(window)

	return ui, nil
}
