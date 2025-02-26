package engine

type ColorblindFilter int

const (
	ColorblindNone ColorblindFilter = iota
	ColorblindProtanopia
	ColorblindDeuteranopia
	ColorblindTritanopia
	ColorblindGrayscale
)

func (c ColorblindFilter) String() string {
	switch c {
	case ColorblindNone:
		return "ColorblindNone"
	case ColorblindProtanopia:
		return "ColorblindProtanopia"
	case ColorblindDeuteranopia:
		return "ColorblindDeuteranopia"
	case ColorblindTritanopia:
		return "ColorblindTritanopia"
	case ColorblindGrayscale:
		return "ColorblindGrayscale"
	default:
		panic("invalid colorblind filter")
	}
}
