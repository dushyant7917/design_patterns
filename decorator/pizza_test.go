package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPizza(t *testing.T) {
	assert := assert.New(t)

	vegPizza := &VegPizza{}
	cheeseToppingsPizza := &CheeseToppings{Pizza: vegPizza}
	tomatoToppingsPizza := &TomatoToppings{Pizza: cheeseToppingsPizza}
	assert.Equal(tomatoToppingsPizza.GetPrice(), 130)

	nonVegPizza := &NonVegPizza{}
	cheeseToppingsPizza = &CheeseToppings{Pizza: nonVegPizza}
	assert.Equal(cheeseToppingsPizza.GetPrice(), 170)
}
