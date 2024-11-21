package anim

import (
	"image"

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

func (ac *AnimController) Update() {
	if ac.curAnim == "" {
		return
	}
	ac.GetAnim(ac.curAnim).Update()
}

func (ac *AnimController) Draw(img *ebiten.Image, pt image.Point) {
	if ac.curAnim == "" {
		return
	}
	ac.GetAnim(ac.curAnim).Draw(img, pt)
}

func (ac *AnimController) Play(name string) {
	ac.curAnim = name
}
