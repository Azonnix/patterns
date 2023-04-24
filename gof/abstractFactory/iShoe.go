package main

type iShoe interface {
	setName(name string)
	setSize(size int)
	getName() string
	getSize() int
}
