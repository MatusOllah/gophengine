//go:build !js

package dialog

import "github.com/ncruces/zenity"

func SelectFileOpen(title string, filename string, filters []FileFilter) (string, error) {
	var zenityFilters zenity.FileFilters
	for _, f := range filters {
		zenityFilters = append(zenityFilters, zenity.FileFilter{
			Name:     f.Name,
			Patterns: f.Patterns,
			CaseFold: f.CaseFold,
		})
	}

	return zenity.SelectFile(
		zenity.Title(title),
		zenity.Filename(filename),
		zenityFilters,
	)
}

func SelectFileSave(title string, filename string, filters []FileFilter) (string, error) {
	var zenityFilters zenity.FileFilters
	for _, f := range filters {
		zenityFilters = append(zenityFilters, zenity.FileFilter{
			Name:     f.Name,
			Patterns: f.Patterns,
			CaseFold: f.CaseFold,
		})
	}

	return zenity.SelectFileSave(
		zenity.Title(title),
		zenity.Filename(filename),
		zenityFilters,
		zenity.ConfirmOverwrite(),
	)
}
