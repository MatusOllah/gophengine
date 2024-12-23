package i18nutil

import "github.com/nicksnyder/go-i18n/v2/i18n"

func Localize(l *i18n.Localizer, msgid string) string {
	return l.MustLocalize(&i18n.LocalizeConfig{
		MessageID: msgid,
	})
}

func L(l *i18n.Localizer, msgid string) string {
	return Localize(l, msgid)
}

func LocalizeTmpl(l *i18n.Localizer, msgid string, tmplData map[string]interface{}) string {
	return l.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    msgid,
		TemplateData: tmplData,
	})
}

func LT(l *i18n.Localizer, msgid string, tmplData map[string]interface{}) string {
	return LocalizeTmpl(l, msgid, tmplData)
}
