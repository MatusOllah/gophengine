package optionsui

import (
	"log/slog"
	"maps"

	"github.com/MatusOllah/gophengine/internal/dialog"
	"github.com/MatusOllah/gophengine/internal/gui"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/pkg/context"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func newMiscellaneousPage(ctx *context.Context, cfg map[string]any, ui *ebitenui.UI) *page {
	c := newPageContentContainer()

	// Locale
	l, curLoc, err := getLocales(ctx)
	if err != nil {
		slog.Error("[locale] failed to get locales", "err", err)
	}

	locStringFunc := func(e any) string {
		return e.(*locale).String()
	}

	comboBox := widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				widget.ComboButtonOpts.MaxContentHeight(200),
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
					widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
					widget.ButtonOpts.Text("", gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
					widget.ButtonOpts.WidgetOpts(
						widget.WidgetOpts.MinSize(150, 0),
					),
				),
			),
		),
		widget.ListComboButtonOpts.ListOpts(
			widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(150, 0))),
			widget.ListOpts.Entries(l),
			widget.ListOpts.ScrollContainerOpts(widget.ScrollContainerOpts.Image(gui.UIRes.ListScrollContainerImage)),
			widget.ListOpts.SliderOpts(
				widget.SliderOpts.Images(gui.UIRes.ListSliderTrackImage, gui.UIRes.ButtonImage),
				widget.SliderOpts.MinHandleSize(5),
				widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2)),
			),
			widget.ListOpts.EntryFontFace(gui.UIRes.Fonts.RegularFace),
			widget.ListOpts.EntryColor(gui.UIRes.ListEntryColor),
			widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
			widget.ListOpts.HideHorizontalSlider(),
		),
		widget.ListComboButtonOpts.EntryLabelFunc(locStringFunc, locStringFunc),
		widget.ListComboButtonOpts.EntrySelectedHandler(func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			slog.Info("[locale] selected entry", "entry", args.Entry)
			cfg["Locale"] = args.Entry.(*locale).locale
		}),
	)

	comboBox.SetSelectedEntry(curLoc)

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(
			widget.LabelOpts.Text(i18n.L("Locale"), gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor),
		),
		comboBox,
	))

	// Separator
	c.AddChild(newSeparator(widget.RowLayoutData{Stretch: true}))

	// Options config
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18n.L("OptionsConfig"), gui.UIRes.Fonts.HeadingFace, gui.UIRes.LabelColor),
	))

	c.AddChild(newHorizontalContainer(
		widget.NewButton(
			widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
			widget.ButtonOpts.Text(i18n.L("Import"), gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked options config import button")
				if err := importOptionsConfig(ctx); err != nil {
					slog.Error("failed to import options config", "err", err)
					dialog.Error("failed to import options config: " + err.Error())
				}
				maps.Copy(cfg, ctx.OptionsConfig.Data()) // update the temporary map so that the changes don't reset when the user clicks "Apply"
			}),
		),
		widget.NewButton(
			widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
			widget.ButtonOpts.Text(i18n.L("Export"), gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked options config export button")
				if err := exportOptionsConfig(ctx); err != nil {
					slog.Error("failed to export options config", "err", err)
					dialog.Error("failed to export options config: " + err.Error())
				}
			}),
		),
		widget.NewButton(
			widget.ButtonOpts.Image(gui.UIRes.DangerButtonImage),
			widget.ButtonOpts.Text(i18n.L("Wipe"), gui.UIRes.Fonts.RegularFace, gui.UIRes.DangerButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked options config wipe button")
				wipeOptionsConfig(ctx, ui)
				maps.Copy(cfg, ctx.OptionsConfig.Data())
			}),
		),
	))

	c.AddChild(widget.NewLabel(widget.LabelOpts.Text("", gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor)))

	// Progress config
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18n.L("ProgressConfig"), gui.UIRes.Fonts.HeadingFace, gui.UIRes.LabelColor),
	))

	c.AddChild(newHorizontalContainer(
		widget.NewButton(
			widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
			widget.ButtonOpts.Text(i18n.L("Import"), gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked progress config import button")
				if err := importProgressConfig(ctx); err != nil {
					slog.Error("failed to import progress config", "err", err)
					dialog.Error("failed to import progress config: " + err.Error())
				}
			}),
		),
		widget.NewButton(
			widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
			widget.ButtonOpts.Text(i18n.L("Export"), gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked progress config export button")
				if err := exportProgressConfig(ctx); err != nil {
					slog.Error("failed to export progress config", "err", err)
					dialog.Error("failed to export progress config: " + err.Error())
				}
			}),
		),
		widget.NewButton(
			widget.ButtonOpts.Image(gui.UIRes.DangerButtonImage),
			widget.ButtonOpts.Text(i18n.L("Wipe"), gui.UIRes.Fonts.RegularFace, gui.UIRes.DangerButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked progress config wipe button")
				wipeProgressConfig(ctx, ui)
			}),
		),
	))

	return &page{
		name:    i18n.L("Miscellaneous"),
		content: c,
	}
}
