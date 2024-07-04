package gophengine

import (
	"fmt"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
)

type StateController struct {
	curState State
}

func NewStateController(state State) *StateController {
	return &StateController{state}
}

func (sc *StateController) Draw(screen *ebiten.Image) {
	sc.curState.Draw(screen)
}

func (sc *StateController) Update(dt float64) error {
	return sc.curState.Update(dt)
}

func (sc *StateController) SwitchState(newState State) {
	slog.Info("[StateController] switching state", "old", fmt.Sprintf("%T", sc.curState), "new", fmt.Sprintf("%T", newState))
	sc.curState = newState
}
