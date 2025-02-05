package gui

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
