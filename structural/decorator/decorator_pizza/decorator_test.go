package decorator_pizza

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &VeggieMania{}

	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithTomatoAndCheese := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of Veggie Pizza with Cheese and Tomato is %d", pizzaWithTomatoAndCheese.getPrice())
}
