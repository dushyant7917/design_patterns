package adapter

type USBPortLaptop interface {
	InsertIntoUSBPort(Device) bool
}

type DellLaptop struct {
}

func (laptop *DellLaptop) InsertIntoUSBPort(device Device) bool {
	return device.GetPortType() == "USB"
}

type HPLaptop struct {
}

func (laptop *HPLaptop) InsertIntoUSBPort(device Device) bool {
	return device.GetPortType() == "USB"
}
