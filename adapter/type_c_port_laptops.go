package adapter

type TypeCPortLaptop interface {
	InsertIntoTypeCPort(Device) bool
}

type MacLaptop struct {
}

func (laptop *MacLaptop) InsertIntoTypeCPort(device Device) bool {
	return device.GetPortType() == TYPE_C_PORT
}
