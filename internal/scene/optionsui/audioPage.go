package optionsui

import (
	"fmt"
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/gopxl/beep/v2/speaker"
)

func mapRange(value, inMin, inMax, outMin, outMax float64) float64 {
	return outMin + (value-inMin)*(outMax-outMin)/(inMax-inMin)
}

func newAudioPage(ctx *context.Context, res *uiResources, cfg map[string]interface{}) *page {
	c := newPageContentContainer()

	// Master volume
	masterVolumeValueLabel := widget.NewLabel(widget.LabelOpts.Text("0%", res.fonts.regularFace, res.labelColor))

	masterVolumeSlider := widget.NewSlider(
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
			masterVolumeValueLabel.Label = fmt.Sprint(args.Current) + "%"

			volume := mapRange(float64(args.Current), 0, 100, -10, 0)

			slog.Debug("[audioPage] setting volume", "volume", volume)
			speaker.Lock()
			ctx.AudioMasterVolume.Volume = volume
			cfg["Audio.MasterVolume"] = volume
			if volume == -10 {
				ctx.AudioMasterVolume.Silent = true
			} else {
				ctx.AudioMasterVolume.Silent = false
			}
			speaker.Unlock()
		}),
	)

	masterVolumeSlider.Current = int(mapRange(cfg["Audio.MasterVolume"].(float64), -10, 0, 0, 100))
	masterVolumeValueLabel.Label = fmt.Sprint(masterVolumeSlider.Current) + "%"

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "MasterVolume")+"  ", res.fonts.regularFace, res.labelColor)),
		masterVolumeSlider,
		widget.NewLabel(widget.LabelOpts.Text(" ", res.fonts.regularFace, res.labelColor)),
		masterVolumeValueLabel,
	))

	return &page{
		name:    i18nutil.Localize(ctx.Localizer, "Audio"),
		content: c,
	}
}
