package scene

import (
	_ "image/png"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/MatusOllah/gophengine/internal/browser"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/scene/mainmenu"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jeandeaual/go-locale"
)

var mainMenuSceneInstance *MainMenuScene

type MainMenuScene struct {
	ctx        *context.Context
	menuItems  *mainmenu.MainMenuItemGroup
	bg         *engine.Sprite
	magenta    *engine.Sprite
	ui         *ebitenui.UI
	shouldExit bool
	bgOffsetY  int // TODO: offset background when selecting menu items
}

var _ engine.Scene = (*MainMenuScene)(nil)

func NewMainMenuScene(ctx *context.Context) *MainMenuScene {
	return &MainMenuScene{ctx: ctx}
}

func (s *MainMenuScene) Init() error {
	s.bgOffsetY = 0
	s.shouldExit = false

	bgImg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/menuBG.png")
	if err != nil {
		return err
	}
	bg := engine.NewSprite(0, 0)
	bg.Img = bgImg
	s.bg = bg

	magentaImg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/menuBGMagenta.png")
	if err != nil {
		return err
	}
	magenta := engine.NewSprite(0, 0)
	magenta.Img = magentaImg
	magenta.Visible = false
	s.magenta = magenta

	storyModeSprite := engine.NewSprite(int(float64(s.ctx.Width/2)-615/2), 0) // Y coordinate handled by mainMenuItemGroup
	storyModeSprite.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/FNF_main_menu_assets/FNF_main_menu_assets.anim.hcl", "story mode")
	if err != nil {
		return err
	}

	freeplaySprite := engine.NewSprite(int(float64(s.ctx.Width/2)-484/2), 0) // Y coordinate handled by mainMenuItemGroup
	freeplaySprite.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/FNF_main_menu_assets/FNF_main_menu_assets.anim.hcl", "freeplay")
	if err != nil {
		return err
	}

	donateSprite := engine.NewSprite(int(float64(s.ctx.Width/2)-444/2), 0) // Y coordinate handled by mainMenuItemGroup
	donateSprite.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/FNF_main_menu_assets/FNF_main_menu_assets.anim.hcl", "donate")
	if err != nil {
		return err
	}

	optionsSprite := engine.NewSprite(int(float64(s.ctx.Width/2)-487/2), 0) // Y coordinate handled by mainMenuItemGroup
	optionsSprite.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/FNF_main_menu_assets/FNF_main_menu_assets.anim.hcl", "options")
	if err != nil {
		return err
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
				return mainMenuSceneInstance.ctx.SceneCtrl.SwitchScene(NewOptionsScene(mainMenuSceneInstance.ctx))
			},
		},
	}

	s.menuItems = mainmenu.NewMainMenuItemGroup(s.ctx, menuItems, magenta)

	ui, err := mainmenu.MakeUI(s.ctx, &s.shouldExit)
	if err != nil {
		return err
	}
	s.ui = ui

	mainMenuSceneInstance = s

	return nil
}

func (s *MainMenuScene) Close() error {
	return nil
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
