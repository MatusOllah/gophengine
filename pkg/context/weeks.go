package context

import (
	"github.com/MatusOllah/gophengine/internal/funkin"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/goreg"
)

// initWeeks initializes the vanilla weeks.
func initWeeks() *goreg.OrderedRegistry[*funkin.Week] {
	weeks := goreg.NewOrderedRegistry[*funkin.Week]()
	weeks.Register("tutorial", &funkin.Week{Name: "", ID: "tutorial", Songs: []string{"tutorial"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "dad", Boyfriend: "bf", Girlfriend: "gf"}})
	weeks.Register("week1", &funkin.Week{Name: "", ID: "week1", Songs: []string{"bopeebo", "fresh", "dadbattle"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "dad", Boyfriend: "bf", Girlfriend: "gf"}})
	weeks.Register("week2", &funkin.Week{Name: "", ID: "week2", Songs: []string{"spookeez", "south"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "spooky", Boyfriend: "bf", Girlfriend: "gf"}})
	weeks.Register("week3", &funkin.Week{Name: "", ID: "week3", Songs: []string{"pico", "philly", "blammed"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "pico", Boyfriend: "bf", Girlfriend: "gf"}, Explicit: true})
	weeks.Register("week4", &funkin.Week{Name: "", ID: "week4", Songs: []string{"satin-panties", "high", "milf"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "mom", Boyfriend: "bf", Girlfriend: "gf"}})
	weeks.Register("week6", &funkin.Week{Name: "", ID: "week6", Songs: []string{"senpai", "roses", "thorns"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "senpai", Boyfriend: "bf", Girlfriend: "gf"}})

	for id, week := range weeks.Iter() {
		week.Name = i18n.L("Week_" + id)
	}

	return weeks
}
