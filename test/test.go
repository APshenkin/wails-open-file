package test

type Fruit string

const (
	Banana Fruit = "Banana"
	Orange Fruit = "Orange"
)

var AllFruits = []Fruit{Banana, Orange}

func (w Fruit) TSName() string {
	switch w {
	case Banana:
		return "B"
	case Orange:
		return "O"
	default:
		return "???"
	}
}

type Color string

const (
	Green Color = "Green"
	Blue  Color = "Blue"
)

var AllColors = []struct {
	Value  Color
	TSName string
}{
	{Green, "G"},
	{Blue, "B"},
}
