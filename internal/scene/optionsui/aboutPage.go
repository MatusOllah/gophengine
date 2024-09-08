package optionsui

func newAboutPage() *page {
	c := newPageContentContainer()

	return &page{
		name:    "About",
		content: c,
	}
}
