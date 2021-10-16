package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUSBToTypeCAdapter(t *testing.T) {
	assert := assert.New(t)
	laptop := &MacLaptop{}  // Supports Type-C
	device1 := &Earphones{} // Type-C
	device2 := &Pendrive{}  // USB
	assert.Equal(laptop.InsertIntoTypeCPort(device1), true)
	assert.Equal(laptop.InsertIntoTypeCPort(device2), false)
	adapter := &USBToTypeCAdapter{laptop: laptop}
	assert.Equal(adapter.InsertIntoUSBPort(device2), true)
}

func TestTypeCToUSBAdapter(t *testing.T) {
	assert := assert.New(t)
	laptop := &DellLaptop{} // Supports USB
	device1 := &Earphones{} // Type-C
	device2 := &Pendrive{}  // USB
	assert.Equal(laptop.InsertIntoUSBPort(device1), false)
	adapter := &TypeCToUSBAdapter{laptop: laptop}
	assert.Equal(adapter.InsertIntoTypeCPort(device1), true)
	assert.Equal(laptop.InsertIntoUSBPort(device2), true)
}
