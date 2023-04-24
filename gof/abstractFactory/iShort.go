package main

type iShort interface {
	setName(name string)
	setColor(color string)
	getName() string
	getColor() string
}
