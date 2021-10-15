package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	assert := assert.New(t)

	house1 := &House{}
	builder1 := GetBuilder("small")
	builder1.BuildHouse(house1)
	assert.Equal(house1.GetWindowType(), "Wooden")
	assert.Equal(house1.GetNumWindows(), 4)
	assert.Equal(house1.GetDoorType(), "Wooden")
	assert.Equal(house1.GetNumDoors(), 2)
	assert.Equal(house1.GetNumSwimmingPools(), 0)
	assert.Equal(house1.GetNumGarages(), 0)

	house2 := &House{}
	builder2 := GetBuilder("big")
	builder2.BuildHouse(house2)
	assert.Equal(house2.GetWindowType(), "Steel")
	assert.Equal(house2.GetNumWindows(), 8)
	assert.Equal(house2.GetDoorType(), "Fiber")
	assert.Equal(house2.GetNumDoors(), 5)
	assert.Equal(house2.GetNumSwimmingPools(), 1)
	assert.Equal(house2.GetNumGarages(), 2)
}
