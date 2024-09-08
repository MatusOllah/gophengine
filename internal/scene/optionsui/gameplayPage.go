package optionsui

func newGameplayPage() *page {
	c := newPageContentContainer()

	return &page{
		name:    "Gameplay",
		content: c,
	}
}
