package optionsui

func newControlsPage() *page {
	c := newPageContentContainer()

	return &page{
		name:    "Controls",
		content: c,
	}
}
