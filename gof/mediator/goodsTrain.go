package main

import "fmt"

type goodsTrain struct {
	mediator iMediator
}

func (t *goodsTrain) requsetArrival() {
	if t.mediator.canLand(t) {
		fmt.Println("goots train EEEE Arrived")
		return
	}

	fmt.Println("goods train add to queye")
}

func (t *goodsTrain) departure() {
	t.mediator.notifyFree()
	fmt.Println("goods train Good by")
}

func (t *goodsTrain) permitArrival() {
	fmt.Println("goods train EEE Arrived from permit")
}
