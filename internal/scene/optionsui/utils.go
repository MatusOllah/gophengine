package optionsui

import (
	"image/color"

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
