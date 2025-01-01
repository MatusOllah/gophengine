package optionsui

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18n"
)

func newControlsPage(ctx *context.Context) *page {
	c := newPageContentContainer()

	return &page{
		name:    i18n.L("Controls"),
		content: c,
	}
}
