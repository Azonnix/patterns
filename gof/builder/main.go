package main

import "fmt"

func main() {
	normBuild := getBuilder("normal")
	iglooBuild := getBuilder("igloo")

	dir := newDirector(normBuild)

	normalHouse := dir.build()

	fmt.Println(normalHouse.windowType, normalHouse.doorType, normalHouse.floor)

	dir.setBuilder(iglooBuild)
	iglooHouse := dir.build()

	fmt.Println(iglooHouse.windowType, iglooHouse.doorType, iglooHouse.floor)
}
