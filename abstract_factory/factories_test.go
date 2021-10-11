package abstract_factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNikeFactory(t *testing.T) {
	assert := assert.New(t)

	factory := GetFactory("Nike")
	shirt := factory.CreateShirt(38)
	shoe := factory.CreateShoe(8)
	assert.Equal(shirt.GetSize(), 38)
	assert.Equal(shoe.GetSize(), 8)
}

func TestAdidasFactory(t *testing.T) {
	assert := assert.New(t)

	factory := GetFactory("Adidas")
	shirt := factory.CreateShirt(40)
	shoe := factory.CreateShoe(9)
	assert.Equal(shirt.GetSize(), 40)
	assert.Equal(shoe.GetSize(), 9)
}
