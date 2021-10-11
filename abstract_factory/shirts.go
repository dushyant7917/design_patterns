package abstract_factory

type Shirt interface {
	GetSize() int
}

type NikeShirt struct {
	size int
}

func (shirt *NikeShirt) GetSize() int {
	return shirt.size
}

type AdidasShirt struct {
	size int
}

func (shirt *AdidasShirt) GetSize() int {
	return shirt.size
}
