package math

// Lerp calculates linear interpolation between a and b in proportion to t.
func Lerp(a, b, t float64) float64 {
	return a + t*(b-a)
}
