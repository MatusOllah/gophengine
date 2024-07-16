package mainmenu

import (
	_ "image/png"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jeandeaual/go-locale"
	"github.com/pkg/browser"
)

var instance *MainMenuState

type MainMenuState struct {
	ctx       *context.Context
	menuItems *mainMenuItemGroup
	bg        *ge.Sprite
	magenta   *ge.Sprite
	bgOffsetY int
}

var _ ge.State = (*MainMenuState)(nil)

func NewMainMenuState(ctx *context.Context) (*MainMenuState, error) {
	bgImg, _, err := ebitenutil.NewImageFromFileSystem(ctx.AssetsFS, "images/menuBG.png")
	if err != nil {
		return nil, err
	}
	bg := ge.NewSprite(0, 0)
	bg.Img = bgImg

	magentaImg, _, err := ebitenutil.NewImageFromFileSystem(ctx.AssetsFS, "images/menuBGMagenta.png")
	if err != nil {
		return nil, err
	}
	magenta := ge.NewSprite(0, 0)
	magenta.Img = magentaImg
	magenta.Visible = false

	storyModeSprite := ge.NewSprite(int(float64(ctx.Width/2)-615/2), 0) // Y coordinate handled by mainMenuItemGroup
	storyModeSprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/FNF_main_menu_assets/FNF_main_menu_assets.anim.hcl", "story mode")
	if err != nil {
		return nil, err
	}

	freeplaySprite := ge.NewSprite(int(float64(ctx.Width/2)-484/2), 0) // Y coordinate handled by mainMenuItemGroup
	freeplaySprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/FNF_main_menu_assets/FNF_main_menu_assets.anim.hcl", "freeplay")
	if err != nil {
		return nil, err
	}

	donateSprite := ge.NewSprite(int(float64(ctx.Width/2)-444/2), 0) // Y coordinate handled by mainMenuItemGroup
	donateSprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/FNF_main_menu_assets/FNF_main_menu_assets.anim.hcl", "donate")
	if err != nil {
		return nil, err
	}

	optionsSprite := ge.NewSprite(int(float64(ctx.Width/2)-487/2), 0) // Y coordinate handled by mainMenuItemGroup
	optionsSprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/FNF_main_menu_assets/FNF_main_menu_assets.anim.hcl", "options")
	if err != nil {
		return nil, err
	}

	menuItems := []*mainMenuItem{
		{
			Name:     "story mode",
			Sprite:   storyModeSprite,
			OnSelect: NopOnSelectFunc,
		},
		{
			Name:     "freeplay",
			Sprite:   freeplaySprite,
			OnSelect: NopOnSelectFunc,
		},
		{
			Name:   "donate",
			Sprite: donateSprite,
			OnSelect: func(i *mainMenuItem) error {
				l, err := locale.GetLocale()
				if err != nil {
					return err
				}

				if l == "sk" || l == "sk-SK" {
					return browser.OpenURL("https://github.com/MatusOllah/gophengine/blob/main/docs/README.sk.md#-darujte")
				} else {
					return browser.OpenURL("https://github.com/MatusOllah/gophengine/blob/main/README.md#-donate")
				}
			},
		},
		{
			Name:     "options",
			Sprite:   optionsSprite,
			OnSelect: NopOnSelectFunc,
		},
	}

	state := &MainMenuState{
		ctx:       ctx,
		menuItems: newMainMenuItemGroup(ctx, menuItems...),
		bg:        bg,
		magenta:   magenta,
		bgOffsetY: 0,
	}

	instance = state

	return state, nil
}

func (s *MainMenuState) Draw(screen *ebiten.Image) {
	bgOpts := s.bg.DrawImageOptions()
	bgOpts.GeoM.Scale(1.1, 1.1)
	bgOpts.GeoM.Translate(0, float64(s.bgOffsetY))

	s.bg.DrawWithOptions(screen, bgOpts)
	s.magenta.DrawWithOptions(screen, bgOpts)

	s.menuItems.Draw(screen)

	ebitenutil.DebugPrintAt(screen, s.ctx.Version, 0, 700)
}

func (s *MainMenuState) Update(dt float64) error {
	return s.menuItems.Update(dt)
}
