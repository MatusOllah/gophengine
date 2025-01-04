package optionsui

import (
	"image/color"
	"math"

	"github.com/ebitenui/ebitenui/widget"
)

// newLabelColorSimple is short for &widget.LabelColor{clr, clr}.
func newLabelColorSimple(clr color.Color) *widget.LabelColor {
	return &widget.LabelColor{Idle: clr, Disabled: clr}
}

// newButtonTextColorSimple is short for &widget.ButtonTextColor{clr, clr, clr, clr}.
func newButtonTextColorSimple(clr color.Color) *widget.ButtonTextColor {
	return &widget.ButtonTextColor{Idle: clr, Disabled: clr, Hover: clr, Pressed: clr}
}

// newHorizontalContainer creates a Container with horizontal row layout.
func newHorizontalContainer(w ...widget.PreferredSizeLocateableWidget) *widget.Container {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(5),
		)),
	)
	c.AddChild(w...)

	return c
}

func newSeparator(res *uiResources, ld interface{}) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    10,
				Bottom: 10,
			}))),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(ld)),
	)

	c.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch:   true,
			MaxHeight: 2,
		})),
		widget.GraphicOpts.ImageNineSlice(res.separatorImage),
	))

	return c
}

func newScrollContainer(res *uiResources, c widget.PreferredSizeLocateableWidget) *widget.Container {
	root := widget.NewContainer(
		// the container will use an grid layout to layout its ScrollableContainer and Slider
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Spacing(2, 0),
			widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{true}),
		)),
	)

	scrollContainer := widget.NewScrollContainer(
		widget.ScrollContainerOpts.Content(c),
		widget.ScrollContainerOpts.StretchContentWidth(),
		widget.ScrollContainerOpts.Image(res.pageListScrollContainerImage),
	)
	root.AddChild(scrollContainer)

	//Create a function to return the page size used by the slider
	pageSizeFunc := func() int {
		return int(math.Round(float64(scrollContainer.ViewRect().Dy()) / float64(c.GetWidget().Rect.Dy()) * 1000))
	}
	//Create a vertical Slider bar to control the ScrollableContainer
	vSlider := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionVertical),
		widget.SliderOpts.MinMax(0, 1000),
		widget.SliderOpts.PageSizeFunc(pageSizeFunc),
		//On change update scroll location based on the Slider's value
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			scrollContainer.ScrollTop = float64(args.Slider.Current) / 1000
		}),
		widget.SliderOpts.Images(res.scrollSliderTrackImage, res.scrollButtonImage),
	)
	//Set the slider's position if the scrollContainer is scrolled by other means than the slider
	scrollContainer.GetWidget().ScrolledEvent.AddHandler(func(args interface{}) {
		a := args.(*widget.WidgetScrolledEventArgs)
		p := pageSizeFunc() / 3
		if p < 1 {
			p = 1
		}
		vSlider.Current -= int(math.Round(a.Y * float64(p)))
	})

	//Add the slider to the second slot in the root container
	root.AddChild(vSlider)

	return root
}
