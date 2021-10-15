package builder

type House struct {
	windowType       string
	numWindows       int
	doorType         string
	numDoors         int
	numSwimmingPools int
	numGarages       int
}

func (house *House) setWindowInfo(windowType string, numWindows int) {
	house.windowType = windowType
	house.numWindows = numWindows
}

func (house *House) setDoorInfo(doorType string, numDoors int) {
	house.doorType = doorType
	house.numDoors = numDoors
}

func (house *House) setNumSwimmingPools(numSwimmingPools int) {
	house.numSwimmingPools = numSwimmingPools
}

func (house *House) setNumGarages(numGarages int) {
	house.numGarages = numGarages
}

func (house *House) GetWindowType() string {
	return house.windowType
}

func (house *House) GetNumWindows() int {
	return house.numWindows
}

func (house *House) GetDoorType() string {
	return house.doorType
}

func (house *House) GetNumDoors() int {
	return house.numDoors
}

func (house *House) GetNumSwimmingPools() int {
	return house.numSwimmingPools
}

func (house *House) GetNumGarages() int {
	return house.numGarages
}
