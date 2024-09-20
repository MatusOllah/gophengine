package optionsui

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
)

func newGraphicsPage(ctx *context.Context) *page {
	c := newPageContentContainer()

	return &page{
		name:    i18nutil.Localize(ctx.Localizer, "OptionsGraphicsPage"),
		content: c,
	}
}
