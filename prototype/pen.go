package prototype

type Pen struct {
	color string
	len   int
}

func (pen *Pen) GetClone() Clonable {
	return &Pen{
		color: pen.color,
		len:   pen.len,
	}
}

func (pen *Pen) IsEqual(clone Clonable) bool {
	anotherPen := clone.(*Pen)
	return pen.color == anotherPen.color && pen.len == anotherPen.len
}
