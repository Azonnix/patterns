package main

type iObserver interface {
	update(string)
	getEmail() string
}
