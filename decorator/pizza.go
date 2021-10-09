package decorator

type Pizza interface {
	GetPrice() int
}

type VegPizza struct {
}

func (pizza *VegPizza) GetPrice() int {
	return 100
}

type NonVegPizza struct {
}

func (pizza *NonVegPizza) GetPrice() int {
	return 150
}
