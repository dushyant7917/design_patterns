package prototype

type Clonable interface {
	GetClone() Clonable
	IsEqual(Clonable) bool
}
