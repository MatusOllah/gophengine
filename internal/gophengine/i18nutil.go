package gophengine

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Localize simply calls G.Localizer.MustLocalize .
func Localize(msgid string) string {
	return G.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: msgid,
	})
}

// Localize simply calls G.Localizer.MustLocalize with TemplateData.
func LocalizeTmpl(msgid string, tmplData map[string]interface{}) string {
	return G.Localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    msgid,
		TemplateData: tmplData,
	})
}
