package gophengine

import "github.com/hajimehoshi/ebiten/v2"

// Scene represents a game scene.
type Scene interface {
	// Init initializes the scene. It's called once before the scene is displayed.
	Init() error

	// Close cleans up resources and closes the scene.
	Close() error

	// Update updates the scene by one tick; dt is the time since the last update (aka delta time).
	Update(dt float64) error

	// Draw renders the scene onto the screen for the current frame.
	Draw(screen *ebiten.Image)
}
