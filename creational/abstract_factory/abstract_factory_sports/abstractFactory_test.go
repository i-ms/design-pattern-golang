package abstract_factory_sports

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo : %s\n", s.getLogo())
	fmt.Printf("Size : %d\n", s.getSize())
}

func printShirtDetails(s IShoe) {
	fmt.Printf("Logo : %s\n", s.getLogo())
	fmt.Printf("Size : %d\n", s.getSize())
}
