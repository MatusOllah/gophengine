package optionsui

import (
	"io/fs"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/pkg/context"
)

type locale struct {
	name   string
	locale string
}

func (l *locale) String() string {
	return l.name + " (" + l.locale + ")"
}

func getLocales(ctx *context.Context) (locales []any, cur *locale, err error) {
	// Get all locales
	paths, err := fs.Glob(ctx.AssetsFS, "data/i18n/*.toml")
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

		loc := strings.ReplaceAll(strings.ReplaceAll(path, "data/i18n/", ""), ".toml", "")

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
	loc, err := ctx.OptionsConfig.Get("Locale")
	if err != nil {
		return nil, err
	}

	return &locale{
		name:   i18n.L("_Name"),
		locale: loc.(string),
	}, nil
}
