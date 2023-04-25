package main

import "fmt"

type passengerTrain struct {
	mediator iMediator
}

func (t *passengerTrain) requsetArrival() {
	if t.mediator.canLand(t) {
		fmt.Println("pass train eeee arived")
		return
	}
	fmt.Println("pass train add to queue")
}

func (t *passengerTrain) departure() {
	fmt.Println("pass train good by")
	t.mediator.notifyFree()
}

func (t *passengerTrain) permitArrival() {
	fmt.Println("pass tring eee arrived form permit")
}
