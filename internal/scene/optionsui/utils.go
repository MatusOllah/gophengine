package optionsui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
)

// newLabelColorSimple is short for &widget.LabelColor{clr, clr}.
func newLabelColorSimple(clr color.Color) *widget.LabelColor {
	return &widget.LabelColor{clr, clr}
}

// newButtonTextColorSimple is short for &widget.ButtonTextColor{clr, clr, clr, clr}.
func newButtonTextColorSimple(clr color.Color) *widget.ButtonTextColor {
	return &widget.ButtonTextColor{clr, clr, clr, clr}
}
