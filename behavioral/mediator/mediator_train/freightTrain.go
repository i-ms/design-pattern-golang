package mediator_train

import "fmt"

type FreightTrain struct {
	mediator Mediator
}

func (f *FreightTrain) arrive() {
	if !f.mediator.canArrive(f) {
		fmt.Println("Freight Train : Arrival blocked, waiting")
		return
	}
	fmt.Println("Freight Train : Arrived")
}

func (f *FreightTrain) depart() {
	fmt.Println("Freight Train : Leaving")
	f.mediator.notifyAboutDeparture()
}

func (f *FreightTrain) permitArrival() {
	fmt.Println("Freight Train : Arrival permitted, arriving")
	f.arrive()
}
