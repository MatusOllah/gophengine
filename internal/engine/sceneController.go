package engine

import (
	"fmt"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneController struct {
	curScene Scene
}

// NewSceneController creates a new [SceneController].
func NewSceneController(state Scene) *SceneController {
	return &SceneController{state}
}

// Close closes the current scene.
func (sc *SceneController) Close() error {
	return sc.curScene.Close()
}

// Draw draws the current scene onto the screen.
func (sc *SceneController) Draw(screen *ebiten.Image) {
	sc.curScene.Draw(screen)
}

// Update updates the current scene.
func (sc *SceneController) Update() error {
	return sc.curScene.Update()
}

// CurScene returns the current scene.
func (sc *SceneController) CurScene() Scene {
	return sc.curScene
}

// SwitchScene closes the old scene, initializes the new scene and switches to the new scene.
func (sc *SceneController) SwitchScene(newScene Scene) error {
	slog.Info("[SceneController] switching scene", "old", fmt.Sprintf("%T", sc.curScene), "new", fmt.Sprintf("%T", newScene))
	if err := newScene.Init(); err != nil {
		return err
	}
	oldScene := sc.curScene
	if oldScene != nil {
		if err := oldScene.Close(); err != nil {
			return err
		}
	}
	sc.curScene = newScene

	return nil
}

// SetScene just switches the scene.
func (sc *SceneController) SetScene(newScene Scene) {
	sc.curScene = newScene
}
