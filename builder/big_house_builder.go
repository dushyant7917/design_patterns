package builder

type BigHouseBuilder struct {
}

func (builder *BigHouseBuilder) BuildHouse(house *House) {
	house.setWindowInfo("Steel", 8)
	house.setDoorInfo("Fiber", 5)
	house.setNumSwimmingPools(1)
	house.setNumGarages(2)
}
