package scene

import (
	_ "image/png"

	ge "github.com/MatusOllah/gophengine"
	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/MatusOllah/gophengine/internal/scene/mainmenu"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jeandeaual/go-locale"
	"github.com/pkg/browser"
)

var mainMenuSceneInstance *MainMenuScene

type MainMenuScene struct {
	ctx        *context.Context
	menuItems  *mainmenu.MainMenuItemGroup
	bg         *ge.Sprite
	magenta    *ge.Sprite
	ui         *ebitenui.UI
	shouldExit bool
	bgOffsetY  int
}

var _ ge.State = (*MainMenuScene)(nil)

func NewMainMenuScene(ctx *context.Context) (*MainMenuScene, error) {
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

	menuItems := []*mainmenu.MainMenuItem{
		{
			Name:     "story mode",
			Sprite:   storyModeSprite,
			OnSelect: mainmenu.NopOnSelectFunc,
		},
		{
			Name:     "freeplay",
			Sprite:   freeplaySprite,
			OnSelect: mainmenu.NopOnSelectFunc,
		},
		{
			Name:   "donate",
			Sprite: donateSprite,
			OnSelect: func(i *mainmenu.MainMenuItem) error {
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
			Name:   "options",
			Sprite: optionsSprite,
			OnSelect: func(_ *mainmenu.MainMenuItem) error {
				optScene, err := NewOptionsScene(mainMenuSceneInstance.ctx)
				if err != nil {
					return err
				}
				mainMenuSceneInstance.ctx.StateController.SwitchState(optScene)
				return nil
			},
		},
	}

	scene := &MainMenuScene{
		ctx:        ctx,
		menuItems:  mainmenu.NewMainMenuItemGroup(ctx, menuItems, magenta),
		bg:         bg,
		magenta:    magenta,
		bgOffsetY:  0,
		shouldExit: false,
	}

	ui, err := mainmenu.MakeUI(ctx, &scene.shouldExit)
	if err != nil {
		return nil, err
	}
	scene.ui = ui

	mainMenuSceneInstance = scene

	return scene, nil
}

func (s *MainMenuScene) Draw(screen *ebiten.Image) {
	bgOpts := s.bg.DrawImageOptions()
	bgOpts.GeoM.Scale(1.1, 1.1)
	bgOpts.GeoM.Translate(0, float64(s.bgOffsetY))

	s.bg.DrawWithOptions(screen, bgOpts)
	s.magenta.DrawWithOptions(screen, bgOpts)

	s.menuItems.Draw(screen)

	ebitenutil.DebugPrintAt(screen, s.ctx.Version, 0, 700)

	s.ui.Draw(screen)
}

func (s *MainMenuScene) Update(dt float64) error {
	if s.shouldExit {
		return ebiten.Termination
	}

	s.ui.Update()

	return s.menuItems.Update(dt)
}
