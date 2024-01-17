package anim

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimController struct {
	anims   map[string]*Animation
	curAnim string
}

func NewAnimController() *AnimController {
	return &AnimController{
		anims: make(map[string]*Animation),
	}
}

func (ac *AnimController) GetAnim(name string) *Animation {
	return ac.anims[name]
}

func (ac *AnimController) SetAnim(name string, anim *Animation) {
	ac.anims[name] = anim
}

func (ac *AnimController) UpdateWithDelta(dt float64) {
	if ac.curAnim == "" {
		return
	}
	ac.GetAnim(ac.curAnim).UpdateWithDelta(time.Duration(dt * float64(1e9)))
}

func (ac *AnimController) Draw(img *ebiten.Image, op *ebiten.DrawImageOptions) {
	if ac.curAnim == "" {
		return
	}
	ac.GetAnim(ac.curAnim).Draw(img, op)
}

func (ac *AnimController) Play(name string) {
	ac.curAnim = name
}
