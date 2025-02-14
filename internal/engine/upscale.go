package engine

type Upscaling int

const (
	UpscaleNearest Upscaling = iota
	UpscaleLinear
)

func (u Upscaling) String() string {
	switch u {
	case UpscaleNearest:
		return "UpscaleNearest"
	case UpscaleLinear:
		return "UpscaleLinear"
	default:
		panic("invalid upscaling")
	}
}
