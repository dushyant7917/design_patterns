package adapter

type Device interface {
	GetPortType() string
}

type Pendrive struct {
}

func (pendrive *Pendrive) GetPortType() string {
	return "USB"
}

type Earphones struct {
}

func (earphones *Earphones) GetPortType() string {
	return "TYPE_C"
}
