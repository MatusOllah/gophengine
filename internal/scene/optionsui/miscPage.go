package optionsui

import (
	"io/fs"
	"log/slog"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/config"
	"github.com/MatusOllah/gophengine/internal/i18nutil"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/ncruces/zenity"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type locale struct {
	name   string
	locale string
}

func (l *locale) String() string {
	return l.name + " (" + l.locale + ")"
}

func newMiscellaneousPage(ctx *context.Context, res *uiResources, cfg map[string]interface{}) *page {
	c := newPageContentContainer()

	// Locale
	l, curLoc, err := getLocales(ctx)
	if err != nil {
		slog.Error("[locale] failed to get locales", "err", err)
	}

	locStringFunc := func(e any) string {
		return e.(*locale).String()
	}

	comboBox := widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				widget.ComboButtonOpts.MaxContentHeight(200),
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(res.buttonImage),
					widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
					widget.ButtonOpts.Text("", res.fonts.regularFace, res.buttonTextColor),
					widget.ButtonOpts.WidgetOpts(
						widget.WidgetOpts.MinSize(150, 0),
					),
				),
			),
		),
		widget.ListComboButtonOpts.ListOpts(
			widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.MinSize(150, 0))),
			widget.ListOpts.Entries(l),
			widget.ListOpts.ScrollContainerOpts(widget.ScrollContainerOpts.Image(res.listScrollContainerImage)),
			widget.ListOpts.SliderOpts(
				widget.SliderOpts.Images(res.listSliderTrackImage, res.buttonImage),
				widget.SliderOpts.MinHandleSize(5),
				widget.SliderOpts.TrackPadding(widget.NewInsetsSimple(2)),
			),
			widget.ListOpts.EntryFontFace(res.fonts.regularFace),
			widget.ListOpts.EntryColor(res.listEntryColor),
			widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
			widget.ListOpts.HideHorizontalSlider(),
		),
		widget.ListComboButtonOpts.EntryLabelFunc(locStringFunc, locStringFunc),
		widget.ListComboButtonOpts.EntrySelectedHandler(func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			slog.Info("[locale] selected entry", "entry", args.Entry)
			cfg["Locale"] = args.Entry.(*locale).locale
		}),
	)

	comboBox.SetSelectedEntry(curLoc)

	c.AddChild(newHorizontalContainer(
		widget.NewLabel(
			widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "Locale"), res.fonts.regularFace, res.labelColor),
		),
		comboBox,
	))

	// Separator
	c.AddChild(newSeparator(res, widget.RowLayoutData{Stretch: true}))

	// Options config
	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "OptionsConfig"), res.fonts.headingFace, res.labelColor),
	))

	c.AddChild(newHorizontalContainer(
		widget.NewButton(
			widget.ButtonOpts.Image(res.buttonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Import"), res.fonts.regularFace, res.buttonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked options config import button")
				//TODO: this
			}),
		),
		widget.NewButton(
			widget.ButtonOpts.Image(res.buttonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Export"), res.fonts.regularFace, res.buttonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked options config export button")
				if err := exportOptionsConfig(ctx); err != nil {
					slog.Error("failed to export options config", "err", err)
					zenity.Error("failed to export options config: "+err.Error(), zenity.Title("GophEngine error"))
				}
			}),
		),
		widget.NewButton(
			widget.ButtonOpts.Image(res.dangerButtonImage),
			widget.ButtonOpts.Text(i18nutil.Localize(ctx.Localizer, "Wipe"), res.fonts.regularFace, res.dangerButtonTextColor),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				slog.Info("[miscPage] clicked options config wipe button")
				//TODO: this
			}),
		),
	))

	/*
		c.AddChild(widget.NewLabel(widget.LabelOpts.Text("", res.fonts.regularFace, res.labelColor)))

		// Progress config
		c.AddChild(widget.NewLabel(
			widget.LabelOpts.Text(i18nutil.Localize(ctx.Localizer, "ProgressConfig"), res.fonts.headingFace, res.labelColor),
		))
	*/

	return &page{
		name:    i18nutil.Localize(ctx.Localizer, "OptionsMiscellaneousPage"),
		content: c,
	}
}

func getLocales(ctx *context.Context) (locales []any, cur *locale, err error) {
	// Get all locales
	paths, err := fs.Glob(ctx.AssetsFS, "data/locale/*.toml")
	if err != nil {
		return nil, nil, err
	}

	for _, path := range paths {
		b, err := fs.ReadFile(ctx.AssetsFS, path)
		if err != nil {
			return nil, nil, err
		}

		var v struct {
			Name string `toml:"_Name"`
		}
		if err := toml.Unmarshal(b, &v); err != nil {
			return nil, nil, err
		}

		loc := strings.ReplaceAll(strings.ReplaceAll(path, "data/locale/", ""), ".toml", "")

		locales = append(locales, &locale{
			name:   v.Name,
			locale: loc,
		})
	}

	// Get current locale
	cur, err = getCurLocale(ctx)
	if err != nil {
		return nil, nil, err
	}

	return
}

func getCurLocale(ctx *context.Context) (*locale, error) {
	name, err := ctx.Localizer.Localize(&i18n.LocalizeConfig{MessageID: "_Name"})
	if err != nil {
		return nil, err
	}

	loc, err := ctx.OptionsConfig.Get("Locale")
	if err != nil {
		return nil, err
	}

	return &locale{
		name:   name,
		locale: loc.(string),
	}, nil
}

func exportOptionsConfig(ctx *context.Context) error {
	path, _ := zenity.SelectFileSave(
		zenity.Title("Export options config"),
		zenity.Filename("options.gecfg"),
		zenity.ConfirmOverwrite(),
		zenity.FileFilters{
			{"GophEngine Configuration File", []string{"*.gecfg"}, false},
			//{"Gob-encoded File (must be a map[string]any!)", []string{"*.gob"}, false},
		},
	)
	slog.Info("export options config", "path", path)

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
