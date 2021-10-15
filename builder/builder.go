package builder

type Builder interface {
	BuildHouse(*House)
}

func GetBuilder(builderType string) Builder {
	if builderType == "small" {
		return &SmallHouseBuilder{}
	} else if builderType == "big" {
		return &BigHouseBuilder{}
	} else {
		return nil
	}
}
