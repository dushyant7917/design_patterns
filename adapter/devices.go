package adapter

type Device interface {
	SetPortType(string)
	GetPortType() string
}

type BaseDevice struct {
	portType string
}

type Pendrive struct {
	BaseDevice
}

func (pendrive *Pendrive) GetPortType() string {
	return pendrive.portType
}

func (pendrive *Pendrive) SetPortType(portType string) {
	pendrive.portType = portType
}

type Earphones struct {
	BaseDevice
}

func (earphones *Earphones) GetPortType() string {
	return earphones.portType
}

func (earphones *Earphones) SetPortType(portType string) {
	earphones.portType = portType
}
