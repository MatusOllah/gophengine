package engine

import (
	"io"

	"github.com/hajimehoshi/ebiten/v2"
)

// Scene represents a game scene.
type Scene interface {
	io.Closer

	// Init initializes the scene. It's called once before the scene is displayed.
	Init() error

	// Update updates the scene by one tick.
	Update() error

	// Draw renders the scene onto the screen for the current frame.
	Draw(screen *ebiten.Image)
}
