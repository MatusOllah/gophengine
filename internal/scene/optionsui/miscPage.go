package optionsui

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui/widget"
)

func newMiscellaneousPage(ctx *context.Context, res *uiResources) *page {
	c := newPageContentContainer()

	//TODO: enumerate all locales (from `data/locale/*`, embedded FS), get the display names and show them in the list
	//TODO: maybe I could also format the strings? Something like "<display name of the locale> (<BCP 47 code>)", for example "Slovenčina (sk)"

	// Locale
	c.AddChild(newHorizontalContainer(
		widget.NewLabel(
			widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "Locale"), res.fonts.regularFace, res.labelColor),
		),
		widget.NewListComboButton(
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
				widget.ListOpts.Entries(nil), //TODO: this
				widget.ListOpts.ScrollContainerOpts(widget.ScrollContainerOpts.Image(res.listScrollContainerImage)),
				widget.ListOpts.SliderOpts(
					widget.SliderOpts.Images(res.listSliderTrackImage, res.buttonImage),
					widget.SliderOpts.MinHandleSize(5),
					widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2)),
				),
				widget.ListOpts.EntryFontFace(res.fonts.regularFace),
				widget.ListOpts.EntryColor(res.listEntryColor),
				widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
			),
			widget.ListComboButtonOpts.EntryLabelFunc(
				func(e interface{}) string {
					return "TODO" //TODO: this
				},
				func(e any) string {
					return "TODO" //TODO: this
				},
			),
			widget.ListComboButtonOpts.EntrySelectedHandler(func(args *widget.ListComboButtonEntrySelectedEventArgs) {
				slog.Info("[locale] selected entry", "entry", args.Entry)
			}),
		),
	))

	return &page{
		name:    i18nutil.Localize(ctx.Localizer, "OptionsWindowMiscellaneousPage"),
		content: c,
	}
}
