package context

import "github.com/MatusOllah/gophengine/internal/funkin"

func initWeeks() []funkin.Week {
	return []funkin.Week{
		{Name: "", ID: "tutorial", Songs: []string{"tutorial"}, MenuCharacters: []string{"dad", "bf", "gf"}},
		{Name: "Daddy Dearest", ID: "week1", Songs: []string{"bopeebo", "fresh", "dadbattle"}, MenuCharacters: []string{"dad", "bf", "gf"}},
		{Name: "Spooky Month", ID: "week2", Songs: []string{"spookeez", "south"}, MenuCharacters: []string{"spooky", "bf", "gf"}},
		{Name: "Pico", ID: "week3", Songs: []string{"pico", "philly", "blammed"}, MenuCharacters: []string{"pico", "bf", "gf"}},
		{Name: "Mommy Mearest", ID: "week4", Songs: []string{"satin-panties", "high", "milf"}, MenuCharacters: []string{"mom", "bf", "gf"}},
		{Name: "Happy and Merry", ID: "week5", Songs: []string{"cocoa", "eggnog"}, MenuCharacters: []string{"parents-christmas", "bf", "gf"}},
		{Name: "Dating simulator ft. moawling", ID: "week6", Songs: []string{"senpai", "roses", "thorns"}, MenuCharacters: []string{"senpai", "bf", "gf"}},
	}
}
