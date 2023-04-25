package main

type iTrain interface {
	requsetArrival()
	departure()
	permitArrival()
}
