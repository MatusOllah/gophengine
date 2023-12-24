package character

type Character struct {
	AnimOffsets  map[string]interface{}
	IsPlayer     bool
	CurCharacter string
	HoldTimer    float64
}

func New(x, y float64, char string, isPlayer bool) *Character {
	return &Character{
		IsPlayer:     isPlayer,
		CurCharacter: char,
	}
}

func (char *Character) AddOffset(name string, x float64, y float64) {
	char.AnimOffsets[name] = []float64{x, y}
}
