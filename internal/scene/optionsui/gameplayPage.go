package optionsui

import (
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/pkg/context"
)

func newGameplayPage(ctx *context.Context) *page {
	c := newPageContentContainer()

	return &page{
		name:    i18n.L("Gameplay"),
		content: c,
	}
}
