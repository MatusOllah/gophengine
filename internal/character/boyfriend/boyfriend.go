package boyfriend

import (
	"log/slog"

	"github.com/MatusOllah/gophengine/internal/character"
)

// Boyfriend represents Boyfriend, the main player character.
type Boyfriend struct {
	*character.Character
	IsStunned bool
}

// New creates a new [Boyfriend].
func New(x, y float64) *Boyfriend {
	slog.Info("Beep!")

	return &Boyfriend{
		Character: character.New(x, y, "bf", true),
		IsStunned: false,
	}
}

func (bf *Boyfriend) Update() error {
	// TODO: this
	return nil
}
