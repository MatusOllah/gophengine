package engine

import (
	"fmt"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneController struct {
	curScene Scene
}

func NewSceneController(state Scene) *SceneController {
	return &SceneController{state}
}

func (sc *SceneController) Close() error {
	return sc.curScene.Close()
}

func (sc *SceneController) Draw(screen *ebiten.Image) {
	sc.curScene.Draw(screen)
}

func (sc *SceneController) Update() error {
	return sc.curScene.Update()
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
