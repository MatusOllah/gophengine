package optionsui

import (
	"log/slog"
	"maps"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/dialog"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func newMiscellaneousPage(ctx *context.Context, res *uiResources, cfg map[string]interface{}, ui *ebitenui.UI) *page {
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
					widget.ButtonOpts.Image(res.buttonImage),
					widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
					widget.ButtonOpts.Text("", res.fonts.regularFace, res.buttonTextColor),
					widget.ButtonOpts.WidgetOpts(
						widget.WidgetOpts.MinSize(150, 0),
					),
				),
			),
		),
		widget.ListComboButtonOpts.ListOpts(
			widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(150, 0))),
			widget.ListOpts.Entries(l),
			widget.ListOpts.ScrollContainerOpts(widget.ScrollContainerOpts.Image(res.listScrollContainerImage)),
			widget.ListOpts.SliderOpts(
				widget.SliderOpts.Images(res.listSliderTrackImage, res.buttonImage),
				widget.SliderOpts.MinHandleSize(5),
				widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2)),
			),
			widget.ListOpts.EntryFontFace(res.fonts.regularFace),
			widget.ListOpts.EntryColor(res.listEntryColor),
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
			widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "Locale"), res.fonts.regularFace, res.labelColor),
		),
		comboBox,
	))

	// Separator
	c.AddChild(newSeparator(res, widget.RowLayoutData{Stretch: true}))

	// Options config
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsConfig"), res.fonts.headingFace, res.labelColor),
	))

	c.AddChild(newHorizontalContainer(
		widget.NewButton(
			widget.ButtonOpts.Image(res.buttonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Import"), res.fonts.regularFace, res.buttonTextColor),
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
			widget.ButtonOpts.Image(res.buttonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Export"), res.fonts.regularFace, res.buttonTextColor),
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
			widget.ButtonOpts.Image(res.dangerButtonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Wipe"), res.fonts.regularFace, res.dangerButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked options config wipe button")
				wipeOptionsConfig(ctx, res, ui)
				maps.Copy(cfg, ctx.OptionsConfig.Data())
			}),
		),
	))

	c.AddChild(widget.NewLabel(widget.LabelOpts.Text("", res.fonts.regularFace, res.labelColor)))

	// Progress config
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "ProgressConfig"), res.fonts.headingFace, res.labelColor),
	))

	c.AddChild(newHorizontalContainer(
		widget.NewButton(
			widget.ButtonOpts.Image(res.buttonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Import"), res.fonts.regularFace, res.buttonTextColor),
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
			widget.ButtonOpts.Image(res.buttonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Export"), res.fonts.regularFace, res.buttonTextColor),
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
			widget.ButtonOpts.Image(res.dangerButtonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Wipe"), res.fonts.regularFace, res.dangerButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked progress config wipe button")
				wipeProgressConfig(ctx, res, ui)
			}),
		),
	))

	return &page{
		name:    i18nutil.Localize(ctx.Localizer, "Miscellaneous"),
		content: c,
	}
}
