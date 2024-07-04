package mainmenu

import (
	_ "image/png"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pkg/browser"
)

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

	storyModeSprite := ge.NewSprite(int(float64(ctx.GameWidth/2)-615/2), 0) // Y coordinate handled by mainMenuItemGroup
	storyModeSprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/FNF_main_menu_assets/story_mode.anim.hcl")
	if err != nil {
		return nil, err
	}

	freeplaySprite := ge.NewSprite(int(float64(ctx.GameWidth/2)-484/2), 0) // Y coordinate handled by mainMenuItemGroup
	freeplaySprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/FNF_main_menu_assets/freeplay.anim.hcl")
	if err != nil {
		return nil, err
	}

	donateSprite := ge.NewSprite(int(float64(ctx.GameWidth/2)-444/2), 0) // Y coordinate handled by mainMenuItemGroup
	donateSprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/FNF_main_menu_assets/donate.anim.hcl")
	if err != nil {
		return nil, err
	}

	optionsSprite := ge.NewSprite(int(float64(ctx.GameWidth/2)-487/2), 0) // Y coordinate handled by mainMenuItemGroup
	optionsSprite.AnimController, err = animhcl.LoadAnimsFromFS(ctx.AssetsFS, "images/FNF_main_menu_assets/options.anim.hcl")
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
				return browser.OpenURL("https://github.com/MatusOllah/gophengine/blob/main/README.md#-donate")
			},
		},
		{
			Name:     "options",
			Sprite:   optionsSprite,
			OnSelect: NopOnSelectFunc,
		},
	}

	return &MainMenuState{
		ctx:       ctx,
		menuItems: newMainMenuItemGroup(ctx, menuItems...),
		bg:        bg,
		magenta:   magenta,
		bgOffsetY: 0,
	}, nil
}

func (s *MainMenuState) Draw(screen *ebiten.Image) {
	bgOpts := s.bg.DrawImageOptions()
	bgOpts.GeoM.Scale(1.1, 1.1)
	bgOpts.GeoM.Translate(0, float64(s.bgOffsetY))

	s.bg.DrawWithOptions(screen, bgOpts)
	s.magenta.DrawWithOptions(screen, bgOpts)

	s.menuItems.Draw(screen)
}

func (s *MainMenuState) Update(dt float64) error {
	return s.menuItems.Update(dt)
}
