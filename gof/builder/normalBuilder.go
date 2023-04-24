package main

type normalBuilder struct {
	house
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) setWindowType() {
	n.windowType = "wood window"
}

func (n *normalBuilder) setDoorType() {
	n.doorType = "wood door"
}

func (n *normalBuilder) setFloor() {
	n.floor = 3
}

func (n *normalBuilder) getHouse() house {
	return house{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}
