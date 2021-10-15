package prototype

type Notebook struct {
	numPages int
	isRuled  bool
}

func (notebook *Notebook) GetClone() Clonable {
	return &Notebook{
		numPages: notebook.numPages,
		isRuled:  notebook.isRuled,
	}
}

func (notebook *Notebook) IsEqual(clone Clonable) bool {
	anotherNotebook := clone.(*Notebook)
	return (notebook.numPages == anotherNotebook.numPages &&
		notebook.isRuled == anotherNotebook.isRuled)
}
