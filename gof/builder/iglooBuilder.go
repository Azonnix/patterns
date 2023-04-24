package main

type iglooBuilder struct {
	house
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (i *iglooBuilder) setWindowType() {
	i.windowType = "ice window"
}

func (i *iglooBuilder) setDoorType() {
	i.doorType = "ice door"
}

func (i *iglooBuilder) setFloor() {
	i.floor = 1
}

func (i *iglooBuilder) getHouse() house {
	return house{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}
