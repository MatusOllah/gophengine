//go:build !js

package optionsui

import (
	"image"
	"log/slog"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/MatusOllah/gophengine/internal/dialog"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func exportOptionsConfig(ctx *context.Context) error {
	path, err := dialog.SelectFileSave(
		i18nutil.Localize(ctx.Localizer, "ExportOptionsConfig"),
		"options.gecfg",
		[]dialog.FileFilter{{i18nutil.Localize(ctx.Localizer, "GEConfigFile"), []string{"*.gecfg"}, false}},
	)
	if err != nil {
		return err
	}

	slog.Info("exporting options config", "path", path)

	cfg, err := config.New(path, false)
	if err != nil {
		return err
	}

	cfg.SetData(ctx.OptionsConfig.Data())

	if err := cfg.Close(); err != nil {
		return err
	}

	slog.Info("export OK")

	return nil
}

func importOptionsConfig(ctx *context.Context) error {
	path, err := dialog.SelectFileOpen(
		i18nutil.Localize(ctx.Localizer, "ImportOptionsConfig"),
		"options.gecfg",
		[]dialog.FileFilter{{i18nutil.Localize(ctx.Localizer, "GEConfigFile"), []string{"*.gecfg"}, false}},
	)
	if err != nil {
		return err
	}

	slog.Info("importing options config", "path", path)

	cfg, err := config.New(path, false)
	if err != nil {
		return err
	}

	m := cfg.Data()
	slog.Debug("got config", "m", m)
	ctx.OptionsConfig.Append(m)

	if err := cfg.Close(); err != nil {
		return err
	}

	slog.Info("import OK")

	return nil
}

func wipeOptionsConfig(ctx *context.Context, res *uiResources, ui *ebitenui.UI) {
	var confirmDialog *widget.Window

	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.bgImage),
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
		widget.ButtonOpts.Image(res.dangerButtonImage),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Wipe"), res.fonts.regularFace, res.dangerButtonTextColor),
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
		widget.ButtonOpts.Image(res.buttonImage),
		widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsCancelButton"), res.fonts.regularFace, res.buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			slog.Info("[wipeOptionsConfig] clicked cancel button")
			confirmDialog.Close()
		}),
	))

	container.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "WipeOptionsDialogText"), res.fonts.regularFace, res.labelColor),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.MaxWidth(360), // this is for word wrap, 400-(20*2)=360 px
		),
	))

	titleBarContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.titleBarBGImage),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	titleBarContainer.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "Wipe"), res.fonts.titleFace, res.labelColor),
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
	confirmDialog.SetLocation(image.Rect(0, 0, x, y).Add(image.Pt(int(float64(ctx.Width/2)-200), int(float64(ctx.Height/2)-100))))
	ui.AddWindow(confirmDialog)
}
