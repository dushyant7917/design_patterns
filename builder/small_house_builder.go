package builder

type SmallHouseBuilder struct {
}

func (builder *SmallHouseBuilder) BuildHouse(house *House) {
	house.setWindowInfo("Wooden", 4)
	house.setDoorInfo("Wooden", 2)
}
