package character

// Boyfriend = goat
type Boyfriend struct {
	*Character
	IsStunned bool
	IsGoat    bool
}

func NewBoyfriend(x, y float64) *Boyfriend {
	return &Boyfriend{
		Character: NewCharacter(x, y, "bf", true),
		IsStunned: false,
		IsGoat:    true,
	}
}

func (bf *Boyfriend) Update(dt float64) error {
	if !bf.IsGoat {
		panic("(*Boyfriend).IsGoat must be true")
	}

	return nil
}
