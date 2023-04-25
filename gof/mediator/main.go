package main

func main() {
	statManager := newStationManager()

	gTrain := goodsTrain{mediator: statManager}
	pTrain := passengerTrain{mediator: statManager}

	gTrain.requsetArrival()
	pTrain.requsetArrival()

	gTrain.departure()
	pTrain.departure()
}
