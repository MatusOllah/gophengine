package optionsui

func newGraphicsPage() *page {
	c := newPageContentContainer()

	return &page{
		name:    "Graphics",
		content: c,
	}
}
