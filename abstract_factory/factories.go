package abstract_factory

type SportsProductFactory interface {
	CreateShoe(int) Shoe
	CreateShirt(int) Shirt
}

type NikeFactory struct {
}

func (factory *NikeFactory) CreateShoe(size int) Shoe {
	return &NikeShoe{size: size}
}

func (factory *NikeFactory) CreateShirt(size int) Shirt {
	return &NikeShirt{size: size}
}

type AdidasFactory struct {
}

func (factory *AdidasFactory) CreateShoe(size int) Shoe {
	return &AdidasShoe{size: size}
}

func (factory *AdidasFactory) CreateShirt(size int) Shirt {
	return &AdidasShirt{size: size}
}

func GetFactory(brandName string) SportsProductFactory {
	if brandName == "Nike" {
		return &NikeFactory{}
	} else if brandName == "Adidas" {
		return &AdidasFactory{}
	} else {
		return nil
	}
}
