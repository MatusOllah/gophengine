package character

// najlepšia FNF postavička
type Boyfriend struct {
	*Character
	IsStunned bool
}

func NewBoyfriend(x, y float64) *Boyfriend {
	return &Boyfriend{
		Character: NewCharacter(x, y, "bf", true),
		IsStunned: false,
	}
}

func (bf *Boyfriend) Update(dt float64) error {
	return nil
}
