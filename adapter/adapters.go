package adapter

type USBToTypeCAdapter struct {
	laptop TypeCPortLaptop
}

func (adapter *USBToTypeCAdapter) InsertIntoUSBPort(device Device) bool {
	// Do some conversion from USB to Type-C
	return !adapter.laptop.InsertIntoTypeCPort(device)
}

type TypeCToUSBAdapter struct {
	laptop USBPortLaptop
}

func (adapter *TypeCToUSBAdapter) InsertIntoTypeCPort(device Device) bool {
	// Do some conversion from Type-C to USB
	return !adapter.laptop.InsertIntoUSBPort(device)
}
