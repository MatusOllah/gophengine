package optionsui

import (
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/pkg/context"
)

func newControlsPage(ctx *context.Context) *page {
	c := newPageContentContainer()

	return &page{
		name:    i18n.L("Controls"),
		content: c,
	}
}
