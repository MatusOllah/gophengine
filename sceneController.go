package gophengine

import (
	"fmt"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneController struct {
	curScene Scene
}

func NewStateController(state Scene) *SceneController {
	return &SceneController{state}
}

func (sc *SceneController) Draw(screen *ebiten.Image) {
	sc.curScene.Draw(screen)
}

func (sc *SceneController) Update(dt float64) error {
	return sc.curScene.Update(dt)
}

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
