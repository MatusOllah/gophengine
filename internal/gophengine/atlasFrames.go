package gophengine

import "image"

type AtlasFrames struct {
	Frames map[string]image.Rectangle
}

func NewTextureAtlas() *AtlasFrames {
	return &AtlasFrames{
		Frames: make(map[string]image.Rectangle),
	}
}

func NewFromTextureAtlas(ta *TextureAtlas) *AtlasFrames {
	af := NewTextureAtlas()

	for _, st := range ta.SubTextures {
		name := st.Name

		rect := image.Rect(int(st.X), int(st.Y), st.Width, st.Height)

		af.AddAtlasFrame(name, rect)
	}

	return af
}

func (af *AtlasFrames) AddAtlasFrame(name string, frame image.Rectangle) {
	af.Frames[name] = frame
}
