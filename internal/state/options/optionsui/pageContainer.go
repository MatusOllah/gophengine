package optionsui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type pageContainer struct {
	widget     widget.PreferredSizeLocateableWidget
	titleLabel *widget.Label
	flipBook   *widget.FlipBook
}

func newPageContainer(res *uiResources) *pageContainer {
	c := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.TrackHover(false),
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{ // TODO: fix layout
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				StretchVertical:    true,
				Padding:            widget.NewInsetsSimple(10),
			}),
		),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
			widget.RowLayoutOpts.Spacing(10)),
		),
	)

	titleFace := truetype.NewFace(res.notoBold, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	lbl := widget.NewLabel(
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			})),
		),
		widget.LabelOpts.Text("", titleFace, newLabelColorSimple(color.NRGBA{255, 255, 255, 255})),
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
