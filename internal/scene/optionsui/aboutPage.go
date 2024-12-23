package optionsui

import (
	"image"
	"image/color"
	"log/slog"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/pkg/browser"
	"golang.design/x/clipboard"
)

func newAboutPage(ctx *context.Context, res *uiResources, ui *ebitenui.UI) *page {
	c := newPageContentContainer()

	// The labels
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.LT(ctx.Localizer, "GEVersion", map[string]interface{}{"Version": ctx.Version}),
			res.fonts.regularFace,
			res.labelColor,
		),
	))
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.LT(ctx.Localizer, "GoVersion", map[string]interface{}{
				"GoVersion": runtime.Version(),
				"GOOS":      runtime.GOOS,
				"GOARCH":    runtime.GOARCH,
			}),
			res.fonts.regularFace,
			res.labelColor,
		),
	))
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.LT(ctx.Localizer, "FNFVersion", map[string]interface{}{"FNFVersion": ctx.FNFVersion}),
			res.fonts.regularFace,
			res.labelColor,
		),
	))

	c.AddChild(widget.NewLabel(widget.LabelOpts.Text("", res.fonts.regularFace, res.labelColor)))

	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.L(ctx.Localizer, "Credits"),
			res.fonts.regularFace,
			res.labelColor,
		),
	))
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.L(ctx.Localizer, "License"),
			res.fonts.regularFace,
			res.labelColor,
		),
	))
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.L(ctx.Localizer, "Creators"),
			res.fonts.regularFace,
			res.labelColor,
		),
	))

	// Separator
	c.AddChild(newSeparator(res, widget.RowLayoutData{Stretch: true}))

	// The buttons
	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.buttonImage),
		widget.ButtonOpts.Text(i18nutil.L(ctx.Localizer, "ShowBuildInfo"), res.fonts.regularFace, res.buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 10,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("clicked build info button")
			showBuildInfoWindow(ctx, res, ui)
		}),
	))
	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.buttonImage),
		widget.ButtonOpts.Text(i18nutil.L(ctx.Localizer, "GitHubButton"), res.fonts.regularFace, res.buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  10,
			Right: 10,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("clicked GitHub button, opening link")
			browser.OpenURL("https://github.com/MatusOllah/gophengine")
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.ToolTip(widget.NewTextToolTip(
				i18nutil.L(ctx.Localizer, "GitHubButtonTooltip"),
				res.fonts.regularFace,
				color.Black,
				res.tooltipBGImage,
			)),
		),
	))

	return &page{
		name:    i18nutil.L(ctx.Localizer, "About"),
		content: c,
	}
}

func showBuildInfoWindow(ctx *context.Context, res *uiResources, ui *ebitenui.UI) {
	var w *widget.Window

	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.bgImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(20)),
		)),
	)

	var textArea *widget.TextArea

	contextMenu := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Direction(widget.DirectionVertical))),
	)
	contextMenu.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.buttonImage),
		widget.ButtonOpts.Text(i18nutil.L(ctx.Localizer, "Copy"), res.fonts.regularFace, res.buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[showBuildInfoWindow] clicked copy build info button")
			if runtime.GOARCH != "wasm" {
				clipboard.Write(clipboard.FmtText, []byte(textArea.GetText()))
			} else {
				slog.Warn("cannot write to clipboard on wasm")
			}
		}),
	))

	textArea = widget.NewTextArea(
		widget.TextAreaOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  true,
				StretchVertical:    true,
			}),
			widget.WidgetOpts.ContextMenu(contextMenu),
		)),
		widget.TextAreaOpts.ControlWidgetSpacing(2),
		widget.TextAreaOpts.FontColor(color.NRGBA{0xFF, 0xFF, 0xFF, 0xFF}),
		widget.TextAreaOpts.FontFace(res.fonts.monospaceFace),
		widget.TextAreaOpts.Text(""),
		widget.TextAreaOpts.ShowVerticalScrollbar(),
		widget.TextAreaOpts.ShowHorizontalScrollbar(),
		widget.TextAreaOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.TextAreaOpts.ScrollContainerOpts(
			widget.ScrollContainerOpts.Image(res.textAreaScrollContainerImage),
		),
		widget.TextAreaOpts.SliderOpts(
			widget.SliderOpts.Images(res.scrollSliderTrackImage, res.scrollButtonImage),
		),
	)

	slog.Info("[showBuildInfoWindow] reading build info")
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		textArea.SetText("failed to read build info")
	} else {
		s := bi.String()
		s = strings.ReplaceAll(s, "\t", "    ") // replace tabs with 4 spaces because *widget.TextArea for some reason can't render tabs???
		textArea.SetText(s)
	}

	windowContainer.AddChild(textArea)

	// The title bar container
	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.titleBarBGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.L(ctx.Localizer, "BuildInfo"), res.fonts.titleFace, res.labelColor),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				}),
			),
		),
	))
	titleBarContainer.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ButtonOpts.Image(res.dangerButtonImage),
		widget.ButtonOpts.Text("X", res.fonts.monospaceFace, res.dangerButtonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   5,
			Right:  5,
			Top:    0,
			Bottom: 0,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[showBuildInfoWindow] clicked exit button, closing window")
			w.Close()
		}),
	))

	w = widget.NewWindow(
		widget.WindowOpts.Contents(windowContainer),
		widget.WindowOpts.TitleBar(titleBarContainer, 25),
		widget.WindowOpts.Draggable(),
		widget.WindowOpts.CloseMode(widget.NONE),
		widget.WindowOpts.MinSize(320, 240),
		widget.WindowOpts.MaxSize(640, 480),
		widget.WindowOpts.Resizeable(),
	)

	x, y := w.Contents.PreferredSize()
	w.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(ctx.Width/2)-320), int(float64(ctx.Height/2)-240))))
	ui.AddWindow(w)
}
