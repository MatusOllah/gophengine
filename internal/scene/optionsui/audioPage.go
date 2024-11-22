package optionsui

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
)

func newAudioPage(ctx *context.Context) *page {
	c := newPageContentContainer()

	return &page{
		name:    i18nutil.Localize(ctx.Localizer, "Audio"),
		content: c,
	}
}
