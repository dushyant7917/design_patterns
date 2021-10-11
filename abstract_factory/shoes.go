package abstract_factory

type Shoe interface {
	GetSize() int
}

type NikeShoe struct {
	size int
}

func (shoe *NikeShoe) GetSize() int {
	return shoe.size
}

type AdidasShoe struct {
	size int
}

func (shoe *AdidasShoe) GetSize() int {
	return shoe.size
}
