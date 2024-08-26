package boyfriend

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/character"
)

// Self-explanatory.
type Boyfriend struct {
	*character.Character
	IsStunned bool
}

// New allocates and creates a new [Boyfriend].
func New(x, y float64) *Boyfriend {
	slog.Info("Beep!")

	return &Boyfriend{
		Character: character.New(x, y, "bf", true),
		IsStunned: false,
	}
}

func (bf *Boyfriend) Update(dt float64) error {
	// TODO: this
	return nil
}
