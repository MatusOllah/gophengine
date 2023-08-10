package anim

type AnimController struct {
	anims map[string]*Animation
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
