package optionsui

func newAdvancedPage() *page {
	c := newPageContentContainer()

	return &page{
		name:    "Advanced",
		content: c,
	}
}
