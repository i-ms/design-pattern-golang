package decorator_pizza

type CheeseTopping struct {
	pizza Ipizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 20
}
