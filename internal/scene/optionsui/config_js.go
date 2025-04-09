//go:build js

package optionsui

import (
	"bytes"
	"image"
	"log/slog"
	"syscall/js"

	"github.com/MatusOllah/gophengine/internal/dialog"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/gui"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/pkg/config"
	"github.com/MatusOllah/gophengine/pkg/context"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func download(buf *bytes.Buffer, mime string, path string) {
	global := js.Global()

	jsData := global.Get("Uint8Array").New(buf.Len())
	js.CopyBytesToJS(jsData, buf.Bytes())

	a := global.Get("document").Call("createElement", "a")
	blob := global.Get("Blob").New(
		[]any{jsData},
		map[string]any{"type": mime},
	)
	a.Set("href", global.Get("URL").Call("createObjectURL", blob))
	a.Set("download", path)
	a.Call("click")
}

func exportOptionsConfig(ctx *context.Context) error {
	slog.Info("exporting options config")

	cfg, err := config.New("", false)
	if err != nil {
		return err
	}

	cfg.SetData(ctx.OptionsConfig.Data())

	if err := cfg.Flush(); err != nil {
		return err
	}

	b := cfg.GobBytes()
	download(bytes.NewBuffer(b), "application/x-gob", "options.gecfg")

	slog.Info("export OK")

	return nil
}

func importOptionsConfig(_ *context.Context) error {
	slog.Warn("import config unsupported on js/wasm")
	dialog.Warning("import config unsupported on js/wasm")

	return nil
}

func wipeOptionsConfig(ctx *context.Context, ui *ebitenui.UI) {
	var confirmDialog *widget.Window

	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(gui.UIRes.BGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(20)),
		)),
	)

	container.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
		),
		widget.ButtonOpts.Image(gui.UIRes.DangerButtonImage),
		widget.ButtonOpts.Text(i18n.L("Wipe"), gui.UIRes.Fonts.RegularFace, gui.UIRes.DangerButtonTextColor),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[wipeOptionsConfig] clicked wipe button")

			ctx.OptionsConfig.Wipe()
			config.LoadDefaultOptions(ctx.OptionsConfig)

			confirmDialog.Close()
		}),
	))

	container.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
		),
		widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
		widget.ButtonOpts.Text(i18n.L("Cancel"), gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[wipeOptionsConfig] clicked cancel button")
			confirmDialog.Close()
		}),
	))

	container.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18n.L("WipeOptionsDialogText"), gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.MaxWidth(360), // this is for word wrap, 400-(20*2)=360 px
		),
	))

	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(gui.UIRes.TitleBarBGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18n.L("Wipe"), gui.UIRes.Fonts.TitleFace, gui.UIRes.LabelColor),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				}),
			),
		),
	))

	confirmDialog = widget.NewWindow(
		widget.WindowOpts.Contents(container),
		widget.WindowOpts.TitleBar(titleBarContainer, 25),
		widget.WindowOpts.Draggable(),
		widget.WindowOpts.CloseMode(widget.NONE),
		widget.WindowOpts.MinSize(400, 200),
	)

	// Spawn window
	x, y := confirmDialog.Contents.PreferredSize()
	confirmDialog.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(engine.GameWidth/2)-200), int(float64(engine.GameHeight/2)-100))))
	ui.AddWindow(confirmDialog)
}

func exportProgressConfig(ctx *context.Context) error {
	path, err := dialog.SelectFileSave(
		i18n.L("ExportProgressConfig"),
		"progress.gecfg",
		[]dialog.FileFilter{{Name: i18n.L("GEConfigFile"), Patterns: []string{"*.gecfg"}, CaseFold: false}},
	)
	if err != nil {
		return err
	}

	slog.Info("exporting progress config", "path", path)

	cfg, err := config.New("", false)
	if err != nil {
		return err
	}

	cfg.SetData(ctx.ProgressConfig.Data())

	if err := cfg.Flush(); err != nil {
		return err
	}

	b := cfg.GobBytes()
	download(bytes.NewBuffer(b), "application/x-gob", "options.gecfg")

	slog.Info("export OK")

	return nil
}

func importProgressConfig(_ *context.Context) error {
	slog.Warn("import config unsupported on js/wasm")
	dialog.Warning("import config unsupported on js/wasm")

	return nil
}

func wipeProgressConfig(ctx *context.Context, ui *ebitenui.UI) {
	var confirmDialog *widget.Window

	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(gui.UIRes.BGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(20)),
		)),
	)

	container.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
		),
		widget.ButtonOpts.Image(gui.UIRes.DangerButtonImage),
		widget.ButtonOpts.Text(i18n.L("Wipe"), gui.UIRes.Fonts.RegularFace, gui.UIRes.DangerButtonTextColor),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[wipeProgressConfig] clicked wipe button")

			ctx.ProgressConfig.Wipe()
			config.LoadDefaultOptions(ctx.ProgressConfig)

			confirmDialog.Close()
		}),
	))

	container.AddChild(widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
		),
		widget.ButtonOpts.Image(gui.UIRes.ButtonImage),
		widget.ButtonOpts.Text(i18n.L("Cancel"), gui.UIRes.Fonts.RegularFace, gui.UIRes.ButtonTextColor),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[wipeProgressConfig] clicked cancel button")
			confirmDialog.Close()
		}),
	))

	container.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18n.L("WipeProgressDialogText"), gui.UIRes.Fonts.RegularFace, gui.UIRes.LabelColor),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.MaxWidth(360), // this is for word wrap, 400-(20*2)=360 px
		),
	))

	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(gui.UIRes.TitleBarBGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18n.L("Wipe"), gui.UIRes.Fonts.TitleFace, gui.UIRes.LabelColor),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				}),
			),
		),
	))

	confirmDialog = widget.NewWindow(
		widget.WindowOpts.Contents(container),
		widget.WindowOpts.TitleBar(titleBarContainer, 25),
		widget.WindowOpts.Draggable(),
		widget.WindowOpts.CloseMode(widget.NONE),
		widget.WindowOpts.MinSize(400, 200),
	)

	// Spawn window
	x, y := confirmDialog.Contents.PreferredSize()
	confirmDialog.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(engine.GameWidth/2)-200), int(float64(engine.GameHeight/2)-100))))
	ui.AddWindow(confirmDialog)
}
