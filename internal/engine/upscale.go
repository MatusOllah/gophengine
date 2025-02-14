package engine

type Upscaling int

const (
	UpscaleNearest Upscaling = iota
	UpscaleLinear
	UpscaleBicubic
	UpscaleFSR
)

func (u Upscaling) String() string {
	switch u {
	case UpscaleNearest:
		return "UpscaleNearest"
	case UpscaleLinear:
		return "UpscaleLinear"
	case UpscaleBicubic:
		return "UpscaleBicubic"
	case UpscaleFSR:
		return "UpscaleFSR"
	default:
		panic("invalid upscaling")
	}
}
