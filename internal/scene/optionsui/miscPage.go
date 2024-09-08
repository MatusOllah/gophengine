package optionsui

func newMiscellaneousPage() *page {
	c := newPageContentContainer()

	return &page{
		name:    "Miscellaneous",
		content: c,
	}
}
