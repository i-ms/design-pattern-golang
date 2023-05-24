package factory_gun

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun : %s\n", g.getName())
	fmt.Printf("Power : %d\n", g.getPower())
}
