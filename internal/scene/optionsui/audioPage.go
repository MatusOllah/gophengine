package optionsui

import (
	"fmt"
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/dialog"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/vorbis"
)

func mapRange(value, inMin, inMax, outMin, outMax float64) float64 {
	return outMin + (value-inMin)*(outMax-outMin)/(inMax-inMin)
}

func newAudioPage(ctx *context.Context, res *uiResources, cfg map[string]interface{}) *page {
	c := newPageContentContainer()

	//TODO: make the volume sliders aligned

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

			slog.Debug("[audioPage] setting master volume", "volume", volume)
			speaker.Lock()
			ctx.AudioMixer.Master.SetVolume(volume)
			cfg["Audio.MasterVolume"] = volume
			speaker.Unlock()
		}),
	)

	masterVolumeSlider.Current = int(mapRange(cfg["Audio.MasterVolume"].(float64), -10, 0, 0, 100))
	masterVolumeValueLabel.Label = fmt.Sprint(masterVolumeSlider.Current) + "%"

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18n.L("MasterVolume")+"  ", res.fonts.regularFace, res.labelColor)),
		masterVolumeSlider,
		widget.NewLabel(widget.LabelOpts.Text(" ", res.fonts.regularFace, res.labelColor)),
		masterVolumeValueLabel,
	))

	// SFX volume
	sfxVolumeValueLabel := widget.NewLabel(widget.LabelOpts.Text("0%", res.fonts.regularFace, res.labelColor))

	sfxVolumeSlider := widget.NewSlider(
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
			slog.Debug("[audioPage] dragging sfx volume slider", "current", args.Current, "dragging", args.Dragging)
			sfxVolumeValueLabel.Label = fmt.Sprint(args.Current) + "%"

			volume := mapRange(float64(args.Current), 0, 100, -10, 0)

			slog.Debug("[audioPage] setting sfx volume", "volume", volume)
			speaker.Lock()
			ctx.AudioMixer.SFX.SetVolume(volume)
			cfg["Audio.SFXVolume"] = volume
			speaker.Unlock()
		}),
	)

	sfxVolumeSlider.Current = int(mapRange(cfg["Audio.SFXVolume"].(float64), -10, 0, 0, 100))
	sfxVolumeValueLabel.Label = fmt.Sprint(sfxVolumeSlider.Current) + "%"

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18n.L("SFXVolume")+"  ", res.fonts.regularFace, res.labelColor)),
		sfxVolumeSlider,
		widget.NewLabel(widget.LabelOpts.Text(" ", res.fonts.regularFace, res.labelColor)),
		sfxVolumeValueLabel,
	))

	// Music volume
	musicVolumeValueLabel := widget.NewLabel(widget.LabelOpts.Text("0%", res.fonts.regularFace, res.labelColor))

	musicVolumeSlider := widget.NewSlider(
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
			slog.Debug("[audioPage] dragging music volume slider", "current", args.Current, "dragging", args.Dragging)
			musicVolumeValueLabel.Label = fmt.Sprint(args.Current) + "%"

			volume := mapRange(float64(args.Current), 0, 100, -10, 0)

			slog.Debug("[audioPage] setting music volume", "volume", volume)
			speaker.Lock()
			ctx.AudioMixer.Music.SetVolume(volume)
			cfg["Audio.MusicVolume"] = volume
			speaker.Unlock()
		}),
	)

	musicVolumeSlider.Current = int(mapRange(cfg["Audio.MusicVolume"].(float64), -10, 0, 0, 100))
	musicVolumeValueLabel.Label = fmt.Sprint(musicVolumeSlider.Current) + "%"

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18n.L("MusicVolume")+"  ", res.fonts.regularFace, res.labelColor)),
		musicVolumeSlider,
		widget.NewLabel(widget.LabelOpts.Text(" ", res.fonts.regularFace, res.labelColor)),
		musicVolumeValueLabel,
	))

	// Instrumental track volume
	instVolumeValueLabel := widget.NewLabel(widget.LabelOpts.Text("0%", res.fonts.regularFace, res.labelColor))

	instVolumeSlider := widget.NewSlider(
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
			slog.Debug("[audioPage] dragging inst volume slider", "current", args.Current, "dragging", args.Dragging)
			instVolumeValueLabel.Label = fmt.Sprint(args.Current) + "%"

			volume := mapRange(float64(args.Current), 0, 100, -10, 0)

			slog.Debug("[audioPage] setting inst volume", "volume", volume)
			speaker.Lock()
			ctx.AudioMixer.Music_Instrumental.SetVolume(volume)
			cfg["Audio.InstVolume"] = volume
			speaker.Unlock()
		}),
	)

	instVolumeSlider.Current = int(mapRange(cfg["Audio.InstVolume"].(float64), -10, 0, 0, 100))
	instVolumeValueLabel.Label = fmt.Sprint(instVolumeSlider.Current) + "%"

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18n.L("InstVolume")+"  ", res.fonts.regularFace, res.labelColor)),
		instVolumeSlider,
		widget.NewLabel(widget.LabelOpts.Text(" ", res.fonts.regularFace, res.labelColor)),
		instVolumeValueLabel,
	))

	// Voices track volume
	voicesVolumeValueLabel := widget.NewLabel(widget.LabelOpts.Text("0%", res.fonts.regularFace, res.labelColor))

	voicesVolumeSlider := widget.NewSlider(
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
			slog.Debug("[audioPage] dragging voices volume slider", "current", args.Current, "dragging", args.Dragging)
			voicesVolumeValueLabel.Label = fmt.Sprint(args.Current) + "%"

			volume := mapRange(float64(args.Current), 0, 100, -10, 0)

			slog.Debug("[audioPage] setting voices volume", "volume", volume)
			speaker.Lock()
			ctx.AudioMixer.Music_Voices.SetVolume(volume)
			cfg["Audio.VoicesVolume"] = volume
			speaker.Unlock()
		}),
	)

	voicesVolumeSlider.Current = int(mapRange(cfg["Audio.VoicesVolume"].(float64), -10, 0, 0, 100))
	voicesVolumeValueLabel.Label = fmt.Sprint(voicesVolumeSlider.Current) + "%"

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(widget.LabelOpts.Text(i18n.L("VoicesVolume")+"  ", res.fonts.regularFace, res.labelColor)),
		voicesVolumeSlider,
		widget.NewLabel(widget.LabelOpts.Text(" ", res.fonts.regularFace, res.labelColor)),
		voicesVolumeValueLabel,
	))

	// Separator
	c.AddChild(newSeparator(res, widget.RowLayoutData{Stretch: true}))

	// Test button
	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.buttonImage),
		widget.ButtonOpts.Text(i18n.L("TestAudio"), res.fonts.regularFace, res.buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			testAudio(ctx)
		}),
	))

	return &page{
		name:    i18n.L("Audio"),
		content: c,
	}
}

func testAudio(ctx *context.Context) {
	slog.Info("[audioPage] testing audio")

	path := "sounds/test_beep.ogg"
	if ctx.Rand.Float64() < 0.05 { // 5% chance
		path = "sounds/bf_test_beep.ogg"
	}

	file, err := ctx.AssetsFS.Open(path)
	if err != nil {
		slog.Error("[audioPage] failed to test audio", "err", err)
		dialog.Error("failed to test audio: " + err.Error())
	}
	defer file.Close()

	streamer, format, err := vorbis.Decode(file)
	if err != nil {
		slog.Error("[audioPage] failed to test audio", "err", err)
		dialog.Error("failed to test audio: " + err.Error())
	}

	ctx.AudioMixer.Master.Add(&effects.Volume{
		Streamer: beep.Resample(ctx.AudioResampleQuality, format.SampleRate, ctx.SampleRate, streamer),
		Base:     2,
		Volume:   -1.5,
	})
}
