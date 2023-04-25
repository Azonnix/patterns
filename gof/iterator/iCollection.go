package main

type iCollection interface {
	createIterator() *iterator
}
