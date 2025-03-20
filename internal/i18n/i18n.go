// Package i18n provides methods for localizing strings.
package i18n

import (
	"io/fs"
	"log/slog"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type LocalizeConfig = i18n.LocalizeConfig

// the global localizer instance
var localizer *i18n.Localizer

// Init loads message files from the file system and initializes the localizer.
func Init(fsys fs.FS, lang string) error {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if err := loadLocales(bundle, fsys); err != nil {
		return err
	}

	localizer = i18n.NewLocalizer(bundle, lang, "en") // fallback to English if something goes wrong

	return nil
}

// loadLocales loads message files from the file system into the bundle.
func loadLocales(bundle *i18n.Bundle, fsys fs.FS) error {
	files, err := fs.Glob(fsys, "data/i18n/*.toml")
	if err != nil {
		return err
	}

	for _, file := range files {
		slog.Info("loading locale", "file", file)
		if _, err := bundle.LoadMessageFileFS(fsys, file); err != nil {
			return err
		}
	}

	return nil
}

// L returns a localized string.
func L(msgid string) string {
	if localizer == nil {
		slog.Warn("localizer not initalized yet!")
		return ""
	}

	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: msgid,
	})
}

// LT returns a localized string with template data.
func LT(msgid string, tmplData map[string]any) string {
	if localizer == nil {
		slog.Warn("localizer not initalized yet!")
		return ""
	}

	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    msgid,
		TemplateData: tmplData,
	})
}

// LC returns a localizes string based on the LocalizeConfig object.
func LC(cfg *i18n.LocalizeConfig) string {
	if localizer == nil {
		slog.Warn("localizer not initalized yet!")
		return ""
	}

	return localizer.MustLocalize(cfg)
}
