package optionsui

import (
	"github.com/ebitenui/ebitenui/widget"
)

type pageContainer struct {
	widget     widget.PreferredSizeLocateableWidget
	titleLabel *widget.Label
	flipBook   *widget.FlipBook
}

// TODO: use AnchorLayout instead of RowLayout
func newPageContainer(res *uiResources) *pageContainer {
	c := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.TrackHover(false),
		),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
			widget.RowLayoutOpts.Spacing(5)),
		),
	)

	lbl := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			})),
		),
		widget.LabelOpts.Text("", res.fonts.titleFace, res.labelColor),
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
