package optionsui

import "github.com/ebitenui/ebitenui/widget"

func newTestPage() *page {
	return &page{
		name:    "Test",
		content: widget.NewContainer(),
	}
}
