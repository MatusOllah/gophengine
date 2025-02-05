package optionsui

import (
	"github.com/MatusOllah/gophengine/internal/gui"
	"github.com/ebitenui/ebitenui/widget"
)

type pageContainer struct {
	widget     widget.PreferredSizeLocateableWidget
	titleLabel *widget.Label
	flipBook   *widget.FlipBook
}

// TODO: use AnchorLayout instead of RowLayout
func newPageContainer() *pageContainer {
	c := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.TrackHover(false),
		),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
			widget.RowLayoutOpts.Spacing(5)),
		),
	)

	lbl := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			})),
		),
		widget.LabelOpts.Text("", gui.UIRes.Fonts.TitleFace, gui.UIRes.LabelColor),
	)
	c.AddChild(lbl)

	flipBook := widget.NewFlipBook(
		widget.FlipBookOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		}))),
	)
	c.AddChild(flipBook)

	return &pageContainer{
		widget:     c,
		titleLabel: lbl,
		flipBook:   flipBook,
	}
}

func (pc *pageContainer) setPage(page *page) {
	pc.titleLabel.Label = page.name
	pc.flipBook.SetPage(page.content)
	pc.flipBook.RequestRelayout()
}
