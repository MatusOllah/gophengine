package optionsui

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui/widget"
)

func newAudioPage(ctx *context.Context, res *uiResources) *page {
	c := newPageContentContainer()

	// Master volume
	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "MasterVolume")+"  ", res.fonts.regularFace, res.labelColor)),
		widget.NewSlider(
			widget.SliderOpts.Direction(widget.DirectionHorizontal),
			widget.SliderOpts.MinMax(0, 100),
			widget.SliderOpts.WidgetOpts(widget.WidgetOpts.MinSize(200, 10)),
			widget.SliderOpts.Images(res.sliderTrackImage, res.buttonImage),
			widget.SliderOpts.FixedHandleSize(10),
			widget.SliderOpts.TrackOffset(0),
			widget.SliderOpts.PageSizeFunc(func() int {
				return 1
			}),
			widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
				slog.Debug("[audioPage] dragging master volume slider", "current", args.Current, "dragging", args.Dragging)
				// TODO: update master volume
			}),
		),
	))

	return &page{
		name:    i18nutil.Localize(ctx.Localizer, "Audio"),
		content: c,
	}
}
