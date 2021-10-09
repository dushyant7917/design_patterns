package decorator

type TomatoToppings struct {
	Pizza Pizza
}

func (tomatoToppings *TomatoToppings) GetPrice() int {
	return tomatoToppings.Pizza.GetPrice() + 10
}

type CheeseToppings struct {
	Pizza Pizza
}

func (cheeseToppings *CheeseToppings) GetPrice() int {
	return cheeseToppings.Pizza.GetPrice() + 20
}
