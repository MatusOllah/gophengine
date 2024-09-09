package optionsui

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type fonts struct {
	titleFace         text.Face
	footerButtonFace  text.Face
	pageListEntryFace text.Face
}

func newFonts(ctx *context.Context) (*fonts, error) {
	notoRegular, err := loadFont(ctx, "fonts/NotoSans-Regular.ttf")
	if err != nil {
		return nil, err
	}

	notoBold, err := loadFont(ctx, "fonts/NotoSans-Bold.ttf")
	if err != nil {
		return nil, err
	}

	return &fonts{
		titleFace: &text.GoTextFace{
			Source: notoBold,
			Size:   24,
		},
		footerButtonFace: &text.GoTextFace{
			Source: notoRegular,
			Size:   24,
		},
		pageListEntryFace: &text.GoTextFace{
			Source: notoRegular,
			Size:   24,
		},
	}, nil
}

func loadFont(ctx *context.Context, path string) (*text.GoTextFaceSource, error) {
	f, err := ctx.AssetsFS.Open(path)
	if err != nil {
		return nil, err
	}

	return text.NewGoTextFaceSource(f)
}
