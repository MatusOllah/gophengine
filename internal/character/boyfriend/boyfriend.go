package boyfriend

import "github.com/MatusOllah/gophengine/internal/character"

type Boyfriend struct {
	*character.Character
	IsStunned bool
}

// New creates a new Boyfriend.
func New(x, y float64) *Boyfriend {
	return &Boyfriend{
		Character: character.New(x, y, "bf", true),
		IsStunned: false,
	}
}

func (bf *Boyfriend) Update(dt float64) error {
	return nil
}
