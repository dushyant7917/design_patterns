package adapter

type USBToTypeCAdapter struct {
	laptop TypeCPortLaptop
}

func (adapter *USBToTypeCAdapter) InsertIntoUSBPort(device Device) bool {
	if device.GetPortType() != USB_PORT {
		return false
	}
	device.SetPortType(TYPE_C_PORT)
	return adapter.laptop.InsertIntoTypeCPort(device)
}

type TypeCToUSBAdapter struct {
	laptop USBPortLaptop
}

func (adapter *TypeCToUSBAdapter) InsertIntoTypeCPort(device Device) bool {
	if device.GetPortType() != TYPE_C_PORT {
		return false
	}
	device.SetPortType(USB_PORT)
	return adapter.laptop.InsertIntoUSBPort(device)
}
