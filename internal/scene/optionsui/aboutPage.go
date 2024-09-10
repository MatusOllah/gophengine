package optionsui

import (
	"runtime"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui/widget"
)

// NOTE: Creators: Matúš Ollah (me) & The Funkin' Crew
// Thinking if I could add something like "Made with ❤️ by Matúš Ollah & The Funkin' Crew" or something like that

func newAboutPage(ctx *context.Context, res *uiResources) *page {
	c := newPageContentContainer()

	//TODO: the actual about page (see options.md)

	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.LocalizeTmpl(ctx.Localizer, "AboutPageVersionText", map[string]interface{}{"Version": ctx.Version}),
			res.fonts.regularFace,
			res.labelColor,
		),
	))
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.LocalizeTmpl(ctx.Localizer, "AboutPageGoVersionText", map[string]interface{}{"GoVersion": runtime.Version()}),
			res.fonts.regularFace,
			res.labelColor,
		),
	))
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(
			i18nutil.LocalizeTmpl(ctx.Localizer, "AboutPageFNFVersionText", map[string]interface{}{"FNFVersion": ctx.FNFVersion}),
			res.fonts.regularFace,
			res.labelColor,
		),
	))

	return &page{
		name:    i18nutil.Localize(ctx.Localizer, "OptionsWindowAboutPage"),
		content: c,
	}
}
