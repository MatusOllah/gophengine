package funkin

// StoryMenuCharacters represents the characters shown in the story mode menu.
type StoryMenuCharacters struct {
	// Opponent is the opponent character ID.
	Opponent string

	// Boyfriend is the Boyfriend character ID i.e. the player character.
	Boyfriend string

	// Girlfriend is the Girlfriend character ID i.e. the spectator character.
	Girlfriend string
}

type StoryMenuCharacter struct {
	ID string
	// TODO: this
}
