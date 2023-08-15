package factory_gun

import "fmt"

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type entered.")
}
