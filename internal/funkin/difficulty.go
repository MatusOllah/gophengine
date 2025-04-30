package funkin

type Difficulty int

const (
	DifficultyEasy Difficulty = iota
	DifficultyNormal
	DifficultyHard
)

// String returns a string representation of the difficulty.
func (d Difficulty) String() string {
	switch d {
	case DifficultyEasy:
		return "DifficultyEasy"
	case DifficultyNormal:
		return "DifficultyNormal"
	case DifficultyHard:
		return "DifficultyHard"
	default:
		panic("invalid difficulty")
	}
}
