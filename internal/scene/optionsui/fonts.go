package optionsui

import (
	"github.com/MatusOllah/gophengine/context"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type fonts struct {
	titleFace         text.Face
	footerButtonFace  text.Face
	pageListEntryFace text.Face
	regularFace       text.Face
	headingFace       text.Face
	monospaceFace     text.Face
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

	notoMono, err := loadFont(ctx, "fonts/NotoSansMono-Regular.ttf")
	if err != nil {
		return nil, err
	}

	notoEmoji, err := loadFont(ctx, "fonts/NotoEmoji-Regular.ttf")
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
		regularFace: newMultiFaceSimple(16, notoRegular, notoEmoji),
		headingFace: &text.GoTextFace{
			Source: notoBold,
			Size:   18,
		},
		monospaceFace: &text.GoTextFace{
			Source: notoMono,
			Size:   16,
		},
	}, nil
}

func newMultiFaceSimple(size float64, srcs ...*text.GoTextFaceSource) *text.MultiFace {
	var faces []text.Face
	for _, src := range srcs {
		faces = append(faces, &text.GoTextFace{
			Source: src,
			Size:   size,
		})
	}

	mf, err := text.NewMultiFace(faces...)
	if err != nil {
		panic(err)
	}

	return mf
}

func loadFont(ctx *context.Context, path string) (*text.GoTextFaceSource, error) {
	f, err := ctx.AssetsFS.Open(path)
	if err != nil {
		return nil, err
	}

	return text.NewGoTextFaceSource(f)
}
