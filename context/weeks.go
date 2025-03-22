package context

import "github.com/MatusOllah/gophengine/internal/funkin"

// initWeeks initializes the vanilla weeks.
func initWeeks() []funkin.Week {
	return []funkin.Week{
		{Name: "", ID: "tutorial", Songs: []string{"tutorial"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "dad", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "Daddy Dearest", ID: "week1", Songs: []string{"bopeebo", "fresh", "dadbattle"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "dad", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "Spooky Month", ID: "week2", Songs: []string{"spookeez", "south"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "spooky", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "Pico", ID: "week3", Songs: []string{"pico", "philly", "blammed"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "pico", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "Mommy Mearest", ID: "week4", Songs: []string{"satin-panties", "high", "milf"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "mom", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "Happy and Merry", ID: "week5", Songs: []string{"cocoa", "eggnog"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "parents-christmas", Boyfriend: "bf", Girlfriend: "gf"}},
		{Name: "Dating simulator ft. moawling", ID: "week6", Songs: []string{"senpai", "roses", "thorns"}, MenuCharacters: funkin.StoryMenuCharacters{Opponent: "senpai", Boyfriend: "bf", Girlfriend: "gf"}},
	}
}
