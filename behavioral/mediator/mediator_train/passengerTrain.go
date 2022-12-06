package mediator_train

import "fmt"

type PassengerTrain struct {
	mediator Mediator
}

func (p *PassengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("Passenger Train : Arrival blocked, waiting")
		return
	}
	fmt.Println("Passenger Train : Arrived")
}

func (p *PassengerTrain) depart() {
	fmt.Println("Passenger Train : Leaving")
	p.mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) permitArrival() {
	fmt.Println("Passenger Train : Arrival permitted, arriving")
	p.arrive()
}
