package main

import "fmt"

func main() {
	adidasFactory := getSportsFactory("adidas")
	pumaFactory := getSportsFactory("puma")

	aShort := adidasFactory.makeShort()
	aShoe := adidasFactory.makeShoe()

	pShort := pumaFactory.makeShort()
	pShoe := pumaFactory.makeShoe()

	fmt.Println(aShort.getName(), aShort.getColor())
	fmt.Println(aShoe.getName(), aShoe.getSize())

	fmt.Println(pShort.getName(), pShort.getColor())
	fmt.Println(pShoe.getName(), pShoe.getSize())
}
