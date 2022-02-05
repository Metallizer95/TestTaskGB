package domain

type Holders struct {
	Holders map[string]float64
}

func NewHolders() Holders {
	return Holders{
		make(map[string]float64),
	}
}
