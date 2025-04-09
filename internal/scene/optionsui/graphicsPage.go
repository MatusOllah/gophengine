package optionsui

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/gui"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/pkg/context"
	"github.com/ebitenui/ebitenui/widget"
)

func newGraphicsPage(ctx *context.Context, cfg map[string]any) *page {
	c := newPageContentContainer()

	// FPS counter
	c.AddChild(widget.NewLabeledCheckbox(
		widget.LabeledCheckboxOpts.LabelOpts(
			widget.LabelOpts.Text(i18n.L("EnableFPSCounter"), gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor),
		),
		widget.LabeledCheckboxOpts.CheckboxOpts(
			widget.CheckboxOpts.ButtonOpts(
				widget.ButtonOpts.Image(gui.UIRes.CheckboxButtonImage),
			),
			widget.CheckboxOpts.Image(gui.UIRes.CheckboxGraphic),
			widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
				slog.Info("[graphicsPage] clicked enable FPS counter checkbox", "state", args.State)
				cfg["Graphics.EnableFPSCounter"] = args.State == widget.WidgetChecked
			}),
			widget.CheckboxOpts.InitialState(func() widget.WidgetState {
				if cfg["Graphics.EnableFPSCounter"].(bool) {
					return widget.WidgetChecked
				}
				return widget.WidgetUnchecked
			}()),
		),
	))

	// VSync
	c.AddChild(widget.NewLabeledCheckbox(
		widget.LabeledCheckboxOpts.LabelOpts(
			widget.LabelOpts.Text(i18n.L("EnableVSync"), gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor),
		),
		widget.LabeledCheckboxOpts.CheckboxOpts(
			widget.CheckboxOpts.ButtonOpts(
				widget.ButtonOpts.Image(gui.UIRes.CheckboxButtonImage),
			),
			widget.CheckboxOpts.Image(gui.UIRes.CheckboxGraphic),
			widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
				slog.Info("[graphicsPage] clicked enable VSync checkbox", "state", args.State)
				cfg["Graphics.EnableVSync"] = args.State == widget.WidgetChecked
			}),
			widget.CheckboxOpts.InitialState(func() widget.WidgetState {
				if cfg["Graphics.EnableVSync"].(bool) {
					return widget.WidgetChecked
				}
				return widget.WidgetUnchecked
			}()),
		),
	))

	// Upscaling
	upscaleStringFunc := func(v any) string {
		return i18n.L(v.(engine.Upscaling).String())
	}

	upscaleMethods := []any{
		engine.UpscaleNearest,
		engine.UpscaleLinear,
	}

	upscaleComboBox := widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				widget.ComboButtonOpts.MaxContentHeight(200),
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
					widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
					widget.ButtonOpts.Text("", gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
					widget.ButtonOpts.WidgetOpts(
						widget.WidgetOpts.MinSize(150, 0),
					),
				),
			),
		),
		widget.ListComboButtonOpts.ListOpts(
			widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(150, 0))),
			widget.ListOpts.Entries(upscaleMethods),
			widget.ListOpts.ScrollContainerOpts(widget.ScrollContainerOpts.Image(gui.UIRes.ListScrollContainerImage)),
			widget.ListOpts.SliderOpts(
				widget.SliderOpts.Images(gui.UIRes.ListSliderTrackImage, gui.UIRes.ButtonImage),
				widget.SliderOpts.MinHandleSize(5),
				widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2)),
			),
			widget.ListOpts.EntryFontFace(gui.UIRes.Fonts.RegularFace),
			widget.ListOpts.EntryColor(gui.UIRes.ListEntryColor),
			widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
			widget.ListOpts.HideHorizontalSlider(),
		),
		widget.ListComboButtonOpts.EntryLabelFunc(upscaleStringFunc, upscaleStringFunc),
		widget.ListComboButtonOpts.EntrySelectedHandler(func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			slog.Info("[graphicsPage] selected upscaling method entry", "entry", args.Entry)
			cfg["Graphics.UpscaleMethod"] = int(args.Entry.(engine.Upscaling))
		}),
	)
	upscaleComboBox.SetSelectedEntry(engine.Upscaling(cfg["Graphics.UpscaleMethod"].(int)))

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18n.L("UpscaleMethod"), gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor)),
		upscaleComboBox,
	))

	// Colorblind filter
	colorblindStringFunc := func(v any) string {
		return i18n.L(v.(engine.ColorblindFilter).String())
	}

	colorblindFilters := []any{
		engine.ColorblindNone,
		engine.ColorblindProtanopia,
		engine.ColorblindDeuteranopia,
		engine.ColorblindTritanopia,
		engine.ColorblindGrayscale,
	}

	colorblindComboBox := widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				widget.ComboButtonOpts.MaxContentHeight(200),
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
					widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
					widget.ButtonOpts.Text("", gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
					widget.ButtonOpts.WidgetOpts(
						widget.WidgetOpts.MinSize(150, 0),
					),
				),
			),
		),
		widget.ListComboButtonOpts.ListOpts(
			widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(150, 0))),
			widget.ListOpts.Entries(colorblindFilters),
			widget.ListOpts.ScrollContainerOpts(widget.ScrollContainerOpts.Image(gui.UIRes.ListScrollContainerImage)),
			widget.ListOpts.SliderOpts(
				widget.SliderOpts.Images(gui.UIRes.ListSliderTrackImage, gui.UIRes.ButtonImage),
				widget.SliderOpts.MinHandleSize(5),
				widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2)),
			),
			widget.ListOpts.EntryFontFace(gui.UIRes.Fonts.RegularFace),
			widget.ListOpts.EntryColor(gui.UIRes.ListEntryColor),
			widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
			widget.ListOpts.HideHorizontalSlider(),
		),
		widget.ListComboButtonOpts.EntryLabelFunc(colorblindStringFunc, colorblindStringFunc),
		widget.ListComboButtonOpts.EntrySelectedHandler(func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			slog.Info("[graphicsPage] selected colorblind filter entry", "entry", args.Entry)
			cfg["Graphics.ColorblindFilter"] = int(args.Entry.(engine.ColorblindFilter))
		}),
	)
	colorblindComboBox.SetSelectedEntry(engine.ColorblindFilter(cfg["Graphics.ColorblindFilter"].(int)))

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18n.L("ColorblindFilter"), gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor)),
		colorblindComboBox,
	))

	return &page{
		name:    i18n.L("Graphics"),
		content: c,
	}
}
