package gui

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Fonts struct {
	TitleFace         text.Face
	FooterButtonFace  text.Face
	PageListEntryFace text.Face
	RegularFace       text.Face
	HeadingFace       text.Face
	MonospaceFace     text.Face
}

func newFonts(fsys fs.FS) (*Fonts, error) {
	notoRegular, err := loadFont(fsys, "fonts/NotoSans-Regular.ttf")
	if err != nil {
		return nil, err
	}

	notoBold, err := loadFont(fsys, "fonts/NotoSans-Bold.ttf")
	if err != nil {
		return nil, err
	}

	notoMono, err := loadFont(fsys, "fonts/NotoSansMono-Regular.ttf")
	if err != nil {
		return nil, err
	}

	notoEmoji, err := loadFont(fsys, "fonts/NotoEmoji-Regular.ttf")
	if err != nil {
		return nil, err
	}

	return &Fonts{
		TitleFace: &text.GoTextFace{
			Source: notoBold,
			Size:   24,
		},
		FooterButtonFace: &text.GoTextFace{
			Source: notoRegular,
			Size:   24,
		},
		PageListEntryFace: &text.GoTextFace{
			Source: notoRegular,
			Size:   24,
		},
		RegularFace: newMultiFaceSimple(16, notoRegular, notoEmoji),
		HeadingFace: &text.GoTextFace{
			Source: notoBold,
			Size:   18,
		},
		MonospaceFace: &text.GoTextFace{
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

func loadFont(fsys fs.FS, path string) (*text.GoTextFaceSource, error) {
	f, err := fsys.Open(path)
	if err != nil {
		return nil, err
	}

	return text.NewGoTextFaceSource(f)
}
