package i18nutil

import "github.com/nicksnyder/go-i18n/v2/i18n"

func Localize(l *i18n.Localizer, msgid string) string {
	return l.MustLocalize(&i18n.LocalizeConfig{
		MessageID: msgid,
	})
}

func LocalizeTmpl(l *i18n.Localizer, msgid string, tmplData map[string]interface{}) string {
	return l.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    msgid,
		TemplateData: tmplData,
	})
}
