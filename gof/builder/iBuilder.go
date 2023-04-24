package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}

	return nil
}
