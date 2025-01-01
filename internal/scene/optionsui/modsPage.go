package optionsui

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18n"
)

func newModsPage(ctx *context.Context) *page {
	c := newPageContentContainer()

	return &page{
		name:    i18n.L("Mods"),
		content: c,
	}
}
