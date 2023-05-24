package builder_building

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	normalBuilder, _ := getBuilder("normal")
	iglooBuilder, _ := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type : %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type : %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Floor Number : %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("Igloo House Door Type : %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type : %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Floor Number : %d\n", iglooHouse.floor)
}
