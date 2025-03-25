package storymenu

import (
	"image"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/funkin"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type MenuItem struct {
	TargetY float64
	Sprite  *engine.Sprite
	Bounds  image.Rectangle
}

func NewMenuItem(ctx *context.Context, x, y int, weekNum int, week *funkin.Week) (*MenuItem, error) {
	item := &MenuItem{TargetY: float64(weekNum), Sprite: engine.NewSprite(x, y)}

	var err error
	item.Sprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/campaign_menu_UI_assets/anim.hcl", week.ID)
	if err != nil {
		return nil, err
	}

	item.Sprite.AnimController.CurAnim().Pause()

	item.Bounds = item.Sprite.AnimController.CurAnim().Frames()[0].Bounds()

	return item, nil
}

func (item *MenuItem) Update() {
	if !item.Sprite.Visible {
		return
	}
	item.Sprite.AnimController.Update()
	item.Sprite.Position.Y = int(gmath.Lerp(float64(item.Sprite.Position.Y), (item.TargetY*120)+480, 0.17))
}

func (item *MenuItem) Draw(img *ebiten.Image) {
	if !item.Sprite.Visible {
		return
	}
	item.Sprite.AnimController.Draw(img, item.Sprite.Position)
}
