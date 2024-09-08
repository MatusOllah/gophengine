package optionsui

func newModsPage() *page {
	c := newPageContentContainer()

	return &page{
		name:    "Mods",
		content: c,
	}
}
