package optionsui

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
)

func newGameplayPage(ctx *context.Context) *page {
	c := newPageContentContainer()

	return &page{
		name:    i18nutil.L(ctx.Localizer, "Gameplay"),
		content: c,
	}
}
