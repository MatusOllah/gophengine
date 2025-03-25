package context

import (
	"github.com/MatusOllah/gophengine/internal/funkin"
	"github.com/MatusOllah/gophengine/internal/i18n"
)

// initWeeks initializes the vanilla weeks.
func initWeeks() []*funkin.Week {
	weeks := []*funkin.Week{
		{Name: "", ID: "tutorial", Songs: []string{"tutorial"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "dad", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "", ID: "week1", Songs: []string{"bopeebo", "fresh", "dadbattle"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "dad", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "", ID: "week2", Songs: []string{"spookeez", "south"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "spooky", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "", ID: "week3", Songs: []string{"pico", "philly", "blammed"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "pico", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "", ID: "week4", Songs: []string{"satin-panties", "high", "milf"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "mom", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "", ID: "week5", Songs: []string{"cocoa", "eggnog"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "parents-christmas", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "", ID: "week6", Songs: []string{"senpai", "roses", "thorns"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "senpai", Boyfriend: "bf", Girlfriend: "gf"}},
	}

	for _, week := range weeks {
		week.Name = i18n.L("Week_" + week.ID)
	}

	return weeks
}
