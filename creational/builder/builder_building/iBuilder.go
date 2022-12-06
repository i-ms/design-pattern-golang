package builder_building

import "fmt"

type IBuilder interface {
	setWindowType()
	setDoorType()
	setFloor()
	getHouse() House
}

func getBuilder(builderType string) (IBuilder, error) {
	if builderType == "normal" {
		return newNormalBuilder(), nil
	}

	if builderType == "igloo" {
		return newIglooBuilder(), nil
	}

	return nil, fmt.Errorf("Invalid bulding type")
}
