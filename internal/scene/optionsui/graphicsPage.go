package optionsui

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/ebitenui/ebitenui/widget"
)

func newGraphicsPage(ctx *context.Context, res *uiResources, cfg map[string]interface{}) *page {
	c := newPageContentContainer()

	c.AddChild(widget.NewLabeledCheckbox(
		widget.LabeledCheckboxOpts.LabelOpts(
			widget.LabelOpts.Text(i18n.L("EnableFPSCounter"), res.fonts.regularFace, res.labelColor),
		),
		widget.LabeledCheckboxOpts.CheckboxOpts(
			widget.CheckboxOpts.ButtonOpts(
				widget.ButtonOpts.Image(res.checkboxButtonImage),
			),
			widget.CheckboxOpts.Image(res.checkboxGraphic),
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

	return &page{
		name:    i18n.L("Graphics"),
		content: c,
	}
}
