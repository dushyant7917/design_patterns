package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	assert := assert.New(t)

	notebook := &Notebook{numPages: 120, isRuled: true}
	pen := &Pen{color: "Blue", len: 67}
	clonableItems := []Clonable{notebook, pen}
	clones := []Clonable{notebook.GetClone(), pen.GetClone()}
	for idx, clone := range clones {
		assert.Equal(clonableItems[idx].IsEqual(clone), true)
	}
}
