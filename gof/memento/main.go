package main

import "fmt"

func main() {
	o := original{state: ""}
	cMemento := collectionMemento{}

	o.setState("Hello")
	cMemento.addMemento(o.saveState())

	o.setState(o.getState() + "World")
	cMemento.addMemento(o.saveState())

	o.setState(o.getState() + "!")
	cMemento.addMemento(o.saveState())

	fmt.Println(o.getState())

	o.loadState(cMemento.getMemento())
	fmt.Println(o.getState())

	o.loadState(cMemento.getMemento())
	fmt.Println(o.getState())

	o.loadState(cMemento.getMemento())
	fmt.Println(o.getState())
}
