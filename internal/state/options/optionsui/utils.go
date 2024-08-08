package optionsui

import (
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
)

// newLabelColorSimple is short for &widget.LabelColor{clr, clr}.
func newLabelColorSimple(clr color.Color) *widget.LabelColor {
	return &widget.LabelColor{clr, clr}
}
