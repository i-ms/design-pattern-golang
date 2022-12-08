package decorator_pizza

type TomatoTopping struct {
	pizza Ipizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 15
}
