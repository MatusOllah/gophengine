package scene

import (
	_ "image/png"

	"github.com/MatusOllah/gophengine/internal/anim/animhcl"
	"github.com/MatusOllah/gophengine/internal/browser"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/scene/mainmenu"
	"github.com/MatusOllah/gophengine/pkg/context"
	"github.com/MatusOllah/gophengine/pkg/version"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jeandeaual/go-locale"
	"github.com/quasilyte/gmath"
)

type MainMenuScene struct {
	ctx        *context.Context
	menuItems  *mainmenu.MainMenuItemGroup
	bg         *engine.Sprite
	magenta    *engine.Sprite
	ui         *ebitenui.UI
	shouldExit bool
	bgOffsetY  float64
}

var _ engine.Scene = (*MainMenuScene)(nil)

func NewMainMenuScene(ctx *context.Context) *MainMenuScene {
	return &MainMenuScene{ctx: ctx}
}

func (s *MainMenuScene) Init() error {
	s.bgOffsetY = 0
	s.shouldExit = false

	bgImg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/ui/bg/menuBG.png")
	if err != nil {
		return err
	}
	bg := engine.NewSprite(0, 0)
	bg.Img = bgImg
	s.bg = bg

	magentaImg, _, err := ebitenutil.NewImageFromFileSystem(s.ctx.AssetsFS, "images/ui/bg/menuBGMagenta.png")
	if err != nil {
		return err
	}
	magenta := engine.NewSprite(0, 0)
	magenta.Img = magentaImg
	magenta.Visible = false
	s.magenta = magenta

	storyModeSprite := engine.NewSprite(0, 0) // Y coordinate handled by mainMenuItemGroup
	storyModeSprite.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/mainmenu/ui/anim.hcl", "story mode")
	if err != nil {
		return err
	}
	storyModeSprite.Position.X = (engine.GameWidth - storyModeSprite.AnimController.GetAnim("idle").Frames()[0].Bounds().Dx()) / 2

	freeplaySprite := engine.NewSprite(0, 0) // Y coordinate handled by mainMenuItemGroup
	freeplaySprite.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/mainmenu/ui/anim.hcl", "freeplay")
	if err != nil {
		return err
	}
	freeplaySprite.Position.X = (engine.GameWidth - freeplaySprite.AnimController.GetAnim("idle").Frames()[0].Bounds().Dx()) / 2

	donateSprite := engine.NewSprite(0, 0) // Y coordinate handled by mainMenuItemGroup
	donateSprite.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/mainmenu/ui/anim.hcl", "donate")
	if err != nil {
		return err
	}
	donateSprite.Position.X = (engine.GameWidth - donateSprite.AnimController.GetAnim("idle").Frames()[0].Bounds().Dx()) / 2

	optionsSprite := engine.NewSprite(0, 0) // Y coordinate handled by mainMenuItemGroup
	optionsSprite.AnimController, err = animhcl.LoadAnimsFromFS(s.ctx.AssetsFS, "images/mainmenu/ui/anim.hcl", "options")
	if err != nil {
		return err
	}
	optionsSprite.Position.X = (engine.GameWidth - optionsSprite.AnimController.GetAnim("idle").Frames()[0].Bounds().Dx()) / 2

	menuItems := []*mainmenu.MainMenuItem{
		{
			Name:   "story mode",
			Sprite: storyModeSprite,
			OnSelect: func(_ *mainmenu.MainMenuItem) error {
				return s.ctx.SceneCtrl.SwitchScene(NewStoryMenuScene(s.ctx))
			},
		},
		{
			Name:     "freeplay",
			Sprite:   freeplaySprite,
			OnSelect: mainmenu.NopOnSelectFunc,
		},
		{
			Name:   "donate",
			Sprite: donateSprite,
			OnSelect: func(_ *mainmenu.MainMenuItem) error {
				l, err := locale.GetLocale()
				if err != nil {
					return err
				}

				if l == "sk" || l == "sk-SK" {
					return browser.OpenURL("https://github.com/MatusOllah/gophengine/blob/main/docs/README.sk.md#%EF%B8%8F-darujte")
				} else {
					return browser.OpenURL("https://github.com/MatusOllah/gophengine/blob/main/README.md#%EF%B8%8F-donate")
				}
			},
		},
		{
			Name:   "options",
			Sprite: optionsSprite,
			OnSelect: func(_ *mainmenu.MainMenuItem) error {
				return s.ctx.SceneCtrl.SwitchScene(NewOptionsScene(s.ctx))
			},
		},
	}

	s.menuItems = mainmenu.NewMainMenuItemGroup(s.ctx, menuItems, magenta)

	ui, err := mainmenu.MakeUI(s.ctx, &s.shouldExit)
	if err != nil {
		return err
	}
	s.ui = ui

	return nil
}

func (s *MainMenuScene) Close() error {
	return nil
}

func (s *MainMenuScene) Draw(screen *ebiten.Image) {
	bgOpts := s.bg.DrawImageOptions()
	bgOpts.GeoM.Scale(1.1, 1.1)
	bgOpts.GeoM.Translate(0, s.bgOffsetY)

	s.bg.DrawWithOptions(screen, bgOpts)
	s.magenta.DrawWithOptions(screen, bgOpts)

	s.menuItems.Draw(screen)

	ebitenutil.DebugPrintAt(screen, version.Version, 0, 700)

	s.ui.Draw(screen)
}

func (s *MainMenuScene) Update() error {
	if s.shouldExit {
		return ebiten.Termination
	}

	s.bgOffsetY = gmath.Lerp(s.bgOffsetY, -float64(s.menuItems.CurSelectedY())/10, 0.06)

	s.ui.Update()

	return s.menuItems.Update()
}
