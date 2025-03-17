package scene

import (
	"fmt"
	"image/color"

	"github.com/MatusOllah/gophengine/context"
	"github.com/MatusOllah/gophengine/internal/engine"
	"github.com/MatusOllah/gophengine/internal/i18n"
	"github.com/MatusOllah/gophengine/internal/scene/title"
	"github.com/hajimehoshi/ebiten/v2"
)

type LoadingScene struct {
	ctx   *context.Context
	next  engine.Scene
	text  *title.IntroText
	errCh chan error
}

var _ engine.Scene = (*LoadingScene)(nil)

func NewLoadingScene(ctx *context.Context, next engine.Scene) *LoadingScene {
	return &LoadingScene{ctx: ctx, next: next, errCh: make(chan error, 1)}
}

func (s *LoadingScene) Init() error {
	var err error
	s.text, err = title.NewIntroText(s.ctx.AssetsFS)
	if err != nil {
		return err
	}

	s.text.CreateText("", i18n.L("Loading"))

	// Load the next scene asynchronously
	go func() {
		defer close(s.errCh)

		if err := s.ctx.SceneCtrl.CurScene().Close(); err != nil {
			s.errCh <- fmt.Errorf("LoadingScene: failed to close old scene: %w", err)
			return
		}
		if err := s.next.Init(); err != nil {
			s.errCh <- fmt.Errorf("LoadingScene: failed to initialize new scene: %w", err)
			return
		}
		s.ctx.SceneCtrl.SetScene(s.next)
	}()

	return nil
}

func (s *LoadingScene) Close() error {
	return nil
}

func (s *LoadingScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	s.text.Draw(screen)
}

func (s *LoadingScene) Update() error {
	select {
	case err := <-s.errCh:
		if err != nil {
			return err
		}
	default:
		// Continue with update routine
	}

	return nil
}
