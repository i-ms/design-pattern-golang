package mediator_train

import "testing"

func TestMediator(t *testing.T) {
	stationManager := newStationManager()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}

	freightTrain := &FreightTrain{
		mediator: stationManager,
	}
	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
