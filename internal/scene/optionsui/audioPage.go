package optionsui

func newAudioPage() *page {
	c := newPageContentContainer()

	return &page{
		name:    "Audio",
		content: c,
	}
}
